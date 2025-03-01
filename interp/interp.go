package interp

import (
	"bufio"
	"context"
	"fmt"
	"go/build"
	"go/scanner"
	"go/token"
	"io"
	"os"
	"os/signal"
	"reflect"
	"strconv"
	"sync"
	"sync/atomic"
)

// Interpreter node structure for AST and CFG
type node struct {
	child  []*node        // child subtrees (AST)
	anc    *node          // ancestor (AST)
	start  *node          // entry point in subtree (CFG)
	tnext  *node          // true branch successor (CFG)
	fnext  *node          // false branch successor (CFG)
	interp *Interpreter   // interpreter context
	frame  *frame         // frame pointer used for closures only (TODO: suppress this)
	index  int64          // node index (dot display)
	findex int            // index of value in frame or frame size (func def, type def)
	level  int            // number of frame indirections to access value
	nleft  int            // number of children in left part (assign)
	nright int            // number of children in right part (assign)
	kind   nkind          // kind of node
	pos    token.Pos      // position in source code, relative to fset
	sym    *symbol        // associated symbol
	typ    *itype         // type of value in frame, or nil
	recv   *receiver      // method receiver node for call, or nil
	types  []reflect.Type // frame types, used by function literals only
	action action         // action
	exec   bltn           // generated function to execute
	gen    bltnGenerator  // generator function to produce above bltn
	val    interface{}    // static generic value (CFG execution)
	rval   reflect.Value  // reflection value to let runtime access interpreter (CFG)
	ident  string         // set if node is a var or func
}

// receiver stores method receiver object access path
type receiver struct {
	node  *node         // receiver value for alias and struct types
	val   reflect.Value // receiver value for interface type and value type
	index []int         // path in receiver value for interface or value type
}

// frame contains values for the current execution level (a function context)
type frame struct {
	anc  *frame          // ancestor frame (global space)
	data []reflect.Value // values

	id uint64 // for cancellation, only access via newFrame/runid/setrunid/clone.

	mutex     sync.RWMutex
	deferred  [][]reflect.Value  // defer stack
	recovered interface{}        // to handle panic recover
	done      reflect.SelectCase // for cancellation of channel operations
}

func newFrame(anc *frame, len int, id uint64) *frame {
	f := &frame{
		anc:  anc,
		data: make([]reflect.Value, len),
		id:   id,
	}
	if anc != nil {
		f.done = anc.done
	}
	return f
}

func (f *frame) runid() uint64      { return atomic.LoadUint64(&f.id) }
func (f *frame) setrunid(id uint64) { atomic.StoreUint64(&f.id, id) }
func (f *frame) clone() *frame {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return &frame{
		anc:       f.anc,
		data:      f.data,
		deferred:  f.deferred,
		recovered: f.recovered,
		id:        f.runid(),
		done:      f.done,
	}
}

// Exports stores the map of binary packages per package path
type Exports map[string]map[string]reflect.Value

// imports stores the map of source packages per package path
type imports map[string]map[string]*symbol

// opt stores interpreter options
type opt struct {
	astDot   bool          // display AST graph (debug)
	cfgDot   bool          // display CFG graph (debug)
	noRun    bool          // compile, but do not run
	fastChan bool          // disable cancellable chan operations
	context  build.Context // build context: GOPATH, build constraints
}

// Interpreter contains global resources and state
type Interpreter struct {
	Name string // program name

	opt                        // user settable options
	cancelChan bool            // enables cancellable chan operations
	nindex     int64           // next node index
	fset       *token.FileSet  // fileset to locate node in source code
	binPkg     Exports         // binary packages used in interpreter, indexed by path
	rdir       map[string]bool // for src import cycle detection

	id uint64 // for cancellation, only accessed via runid/stop

	mutex    sync.RWMutex
	frame    *frame            // program data storage during execution
	universe *scope            // interpreter global level scope
	scopes   map[string]*scope // package level scopes, indexed by package name
	srcPkg   imports           // source packages used in interpreter, indexed by path
	done     chan struct{}     // for cancellation of channel operations
}

const (
	mainID   = "main"
	selfPath = "github.com/containous/yaegi/interp"
)

// Symbols exposes interpreter values
var Symbols = Exports{
	selfPath: map[string]reflect.Value{
		"New": reflect.ValueOf(New),

		"Interpreter": reflect.ValueOf((*Interpreter)(nil)),
		"Options":     reflect.ValueOf((*Options)(nil)),
	},
}

func init() { Symbols[selfPath]["Symbols"] = reflect.ValueOf(Symbols) }

// _error is a wrapper of error interface type
type _error struct {
	WError func() string
}

func (w _error) Error() string { return w.WError() }

// Walk traverses AST n in depth first order, call cbin function
// at node entry and cbout function at node exit.
func (n *node) Walk(in func(n *node) bool, out func(n *node)) {
	if in != nil && !in(n) {
		return
	}
	for _, child := range n.child {
		child.Walk(in, out)
	}
	if out != nil {
		out(n)
	}
}

// Options are the interpreter options.
type Options struct {
	// GoPath sets GOPATH for the interpreter
	GoPath string
	// BuildTags sets build constraints for the interpreter
	BuildTags []string
}

// New returns a new interpreter
func New(options Options) *Interpreter {
	i := Interpreter{
		opt:      opt{context: build.Default},
		frame:    &frame{data: []reflect.Value{}},
		fset:     token.NewFileSet(),
		universe: initUniverse(),
		scopes:   map[string]*scope{},
		binPkg:   Exports{"": map[string]reflect.Value{"_error": reflect.ValueOf((*_error)(nil))}},
		srcPkg:   imports{},
		rdir:     map[string]bool{},
	}

	i.opt.context.GOPATH = options.GoPath
	if len(options.BuildTags) > 0 {
		i.opt.context.BuildTags = options.BuildTags
	}

	// astDot activates AST graph display for the interpreter
	i.opt.astDot, _ = strconv.ParseBool(os.Getenv("YAEGI_AST_DOT"))

	// cfgDot activates AST graph display for the interpreter
	i.opt.cfgDot, _ = strconv.ParseBool(os.Getenv("YAEGI_CFG_DOT"))

	// noRun disables the execution (but not the compilation) in the interpreter
	i.opt.noRun, _ = strconv.ParseBool(os.Getenv("YAEGI_NO_RUN"))

	// fastChan disables the cancellable version of channel operations in evalWithContext
	i.opt.fastChan, _ = strconv.ParseBool(os.Getenv("YAEGI_FAST_CHAN"))
	return &i
}

func initUniverse() *scope {
	sc := &scope{global: true, sym: map[string]*symbol{
		// predefined Go types
		"bool":        {kind: typeSym, typ: &itype{cat: boolT, name: "bool"}},
		"byte":        {kind: typeSym, typ: &itype{cat: byteT, name: "byte"}},
		"complex64":   {kind: typeSym, typ: &itype{cat: complex64T, name: "complex64"}},
		"complex128":  {kind: typeSym, typ: &itype{cat: complex128T, name: "complex128"}},
		"error":       {kind: typeSym, typ: &itype{cat: errorT, name: "error"}},
		"float32":     {kind: typeSym, typ: &itype{cat: float32T, name: "float32"}},
		"float64":     {kind: typeSym, typ: &itype{cat: float64T, name: "float64"}},
		"int":         {kind: typeSym, typ: &itype{cat: intT, name: "int"}},
		"int8":        {kind: typeSym, typ: &itype{cat: int8T, name: "int8"}},
		"int16":       {kind: typeSym, typ: &itype{cat: int16T, name: "int16"}},
		"int32":       {kind: typeSym, typ: &itype{cat: int32T, name: "int32"}},
		"int64":       {kind: typeSym, typ: &itype{cat: int64T, name: "int64"}},
		"interface{}": {kind: typeSym, typ: &itype{cat: interfaceT}},
		"rune":        {kind: typeSym, typ: &itype{cat: runeT, name: "rune"}},
		"string":      {kind: typeSym, typ: &itype{cat: stringT, name: "string"}},
		"uint":        {kind: typeSym, typ: &itype{cat: uintT, name: "uint"}},
		"uint8":       {kind: typeSym, typ: &itype{cat: uint8T, name: "uint8"}},
		"uint16":      {kind: typeSym, typ: &itype{cat: uint16T, name: "uint16"}},
		"uint32":      {kind: typeSym, typ: &itype{cat: uint32T, name: "uint32"}},
		"uint64":      {kind: typeSym, typ: &itype{cat: uint64T, name: "uint64"}},
		"uintptr":     {kind: typeSym, typ: &itype{cat: uintptrT, name: "uintptr"}},

		// predefined Go constants
		"false": {kind: constSym, typ: &itype{cat: boolT, name: "bool"}, rval: reflect.ValueOf(false)},
		"true":  {kind: constSym, typ: &itype{cat: boolT, name: "bool"}, rval: reflect.ValueOf(true)},
		"iota":  {kind: constSym, typ: &itype{cat: intT}},

		// predefined Go zero value
		"nil": {typ: &itype{cat: nilT, untyped: true}},

		// predefined Go builtins
		"append":  {kind: bltnSym, builtin: _append},
		"cap":     {kind: bltnSym, builtin: _cap},
		"close":   {kind: bltnSym, builtin: _close},
		"complex": {kind: bltnSym, builtin: _complex},
		"imag":    {kind: bltnSym, builtin: _imag},
		"copy":    {kind: bltnSym, builtin: _copy},
		"delete":  {kind: bltnSym, builtin: _delete},
		"len":     {kind: bltnSym, builtin: _len},
		"make":    {kind: bltnSym, builtin: _make},
		"new":     {kind: bltnSym, builtin: _new},
		"panic":   {kind: bltnSym, builtin: _panic},
		"print":   {kind: bltnSym, builtin: _print},
		"println": {kind: bltnSym, builtin: _println},
		"real":    {kind: bltnSym, builtin: _real},
		"recover": {kind: bltnSym, builtin: _recover},
	}}
	return sc
}

// resizeFrame resizes the global frame of interpreter
func (interp *Interpreter) resizeFrame() {
	l := len(interp.universe.types)
	b := len(interp.frame.data)
	if l-b <= 0 {
		return
	}
	data := make([]reflect.Value, l)
	copy(data, interp.frame.data)
	for j, t := range interp.universe.types[b:] {
		data[b+j] = reflect.New(t).Elem()
	}
	interp.frame.data = data
}

func (interp *Interpreter) main() *node {
	interp.mutex.RLock()
	defer interp.mutex.RUnlock()
	if m, ok := interp.scopes[mainID]; ok && m.sym[mainID] != nil {
		return m.sym[mainID].node
	}
	return nil
}

// Eval evaluates Go code represented as a string. It returns a map on
// current interpreted package exported symbols
func (interp *Interpreter) Eval(src string) (reflect.Value, error) {
	var res reflect.Value

	// Parse source to AST
	pkgName, root, err := interp.ast(src, interp.Name)
	if err != nil || root == nil {
		return res, err
	}

	if interp.astDot {
		root.astDot(dotX(), interp.Name)
		if interp.noRun {
			return res, err
		}
	}

	// Global type analysis
	revisit, err := interp.gta(root, pkgName)
	if err != nil {
		return res, err
	}
	for _, n := range revisit {
		if _, err = interp.gta(n, pkgName); err != nil {
			return res, err
		}
	}

	// Annotate AST with CFG infos
	initNodes, err := interp.cfg(root)
	if err != nil {
		return res, err
	}

	// Add main to list of functions to run, after all inits
	if m := interp.main(); m != nil {
		initNodes = append(initNodes, m)
	}

	if root.kind != fileStmt {
		// REPL may skip package statement
		setExec(root.start)
	}
	interp.mutex.Lock()
	if interp.universe.sym[pkgName] == nil {
		// Make the package visible under a path identical to its name
		interp.srcPkg[pkgName] = interp.scopes[pkgName].sym
		interp.universe.sym[pkgName] = &symbol{kind: pkgSym, typ: &itype{cat: srcPkgT, path: pkgName}}
	}
	interp.mutex.Unlock()

	if interp.cfgDot {
		root.cfgDot(dotX())
	}

	if interp.noRun {
		return res, err
	}

	// Generate node exec closures
	if err = genRun(root); err != nil {
		return res, err
	}

	// Init interpreter execution memory frame
	interp.frame.setrunid(interp.runid())
	interp.frame.mutex.Lock()
	interp.resizeFrame()
	interp.frame.mutex.Unlock()

	// Execute node closures
	interp.run(root, nil)

	for _, n := range initNodes {
		interp.run(n, interp.frame)
	}
	v := genValue(root)
	res = v(interp.frame)

	// If result is an interpreter node, wrap it in a runtime callable function
	if res.IsValid() {
		if n, ok := res.Interface().(*node); ok {
			res = genFunctionWrapper(n)(interp.frame)
		}
	}

	return res, err
}

// EvalWithContext evaluates Go code represented as a string. It returns
// a map on current interpreted package exported symbols.
func (interp *Interpreter) EvalWithContext(ctx context.Context, src string) (reflect.Value, error) {
	var v reflect.Value
	var err error

	interp.mutex.Lock()
	interp.done = make(chan struct{})
	interp.cancelChan = !interp.opt.fastChan
	interp.mutex.Unlock()

	done := make(chan struct{})
	go func() {
		defer close(done)
		v, err = interp.Eval(src)
	}()

	select {
	case <-ctx.Done():
		interp.stop()
		return reflect.Value{}, ctx.Err()
	case <-done:
		return v, err
	}
}

// stop sends a semaphore to all running frames and closes the chan
// operation short circuit channel. stop may only be called once per
// invocation of EvalWithContext.
func (interp *Interpreter) stop() {
	atomic.AddUint64(&interp.id, 1)
	close(interp.done)
}

func (interp *Interpreter) runid() uint64 { return atomic.LoadUint64(&interp.id) }

// getWrapper returns the wrapper type of the corresponding interface, or nil if not found
func (interp *Interpreter) getWrapper(t reflect.Type) reflect.Type {
	if p, ok := interp.binPkg[t.PkgPath()]; ok {
		return p["_"+t.Name()].Type().Elem()
	}
	return nil
}

// Use loads binary runtime symbols in the interpreter context so
// they can be used in interpreted code
func (interp *Interpreter) Use(values Exports) {
	for k, v := range values {
		interp.binPkg[k] = v
	}
}

// REPL performs a Read-Eval-Print-Loop on input reader.
// Results are printed on output writer.
func (interp *Interpreter) REPL(in io.Reader, out io.Writer) {
	s := bufio.NewScanner(in)
	prompt := getPrompt(in, out)
	prompt()
	src := ""
	for s.Scan() {
		src += s.Text() + "\n"
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		handleSignal(ctx, cancel)
		v, err := interp.EvalWithContext(ctx, src)
		signal.Reset()
		if err != nil {
			switch err.(type) {
			case scanner.ErrorList:
				// Early failure in the scanner: the source is incomplete
				// and no AST could be produced, neither compiled / run.
				// Get one more line, and retry
				continue
			default:
				fmt.Fprintln(out, err)
			}
		} else if v.IsValid() {
			fmt.Fprintln(out, v)
		}
		src = ""
		prompt()
	}
}

// Repl performs a Read-Eval-Print-Loop on input file descriptor.
// Results are printed on output.
// Deprecated: use REPL instead
func (interp *Interpreter) Repl(in, out *os.File) {
	interp.REPL(in, out)
}

// getPrompt returns a function which prints a prompt only if input is a terminal.
func getPrompt(in io.Reader, out io.Writer) func() {
	s, ok := in.(interface{ Stat() (os.FileInfo, error) })
	if !ok {
		return func() {}
	}
	stat, err := s.Stat()
	if err == nil && stat.Mode()&os.ModeCharDevice != 0 {
		return func() { fmt.Fprint(out, "> ") }
	}
	return func() {}
}

// handleSignal wraps signal handling for eval cancellation.
func handleSignal(ctx context.Context, cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()
}
