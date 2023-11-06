// Code generated by 'yaegi extract reflect'. DO NOT EDIT.

//go:build go1.20 && !go1.21
// +build go1.20

package stdlib

import (
	"reflect"
)

func init() {
	Symbols["reflect/reflect"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Append":          reflect.ValueOf(reflect.Append),
		"AppendSlice":     reflect.ValueOf(reflect.AppendSlice),
		"Array":           reflect.ValueOf(reflect.Array),
		"ArrayOf":         reflect.ValueOf(reflect.ArrayOf),
		"Bool":            reflect.ValueOf(reflect.Bool),
		"BothDir":         reflect.ValueOf(reflect.BothDir),
		"Chan":            reflect.ValueOf(reflect.Chan),
		"ChanOf":          reflect.ValueOf(reflect.ChanOf),
		"Complex128":      reflect.ValueOf(reflect.Complex128),
		"Complex64":       reflect.ValueOf(reflect.Complex64),
		"Copy":            reflect.ValueOf(reflect.Copy),
		"DeepEqual":       reflect.ValueOf(reflect.DeepEqual),
		"Float32":         reflect.ValueOf(reflect.Float32),
		"Float64":         reflect.ValueOf(reflect.Float64),
		"Func":            reflect.ValueOf(reflect.Func),
		"FuncOf":          reflect.ValueOf(reflect.FuncOf),
		"Indirect":        reflect.ValueOf(reflect.Indirect),
		"Int":             reflect.ValueOf(reflect.Int),
		"Int16":           reflect.ValueOf(reflect.Int16),
		"Int32":           reflect.ValueOf(reflect.Int32),
		"Int64":           reflect.ValueOf(reflect.Int64),
		"Int8":            reflect.ValueOf(reflect.Int8),
		"Interface":       reflect.ValueOf(reflect.Interface),
		"Invalid":         reflect.ValueOf(reflect.Invalid),
		"MakeChan":        reflect.ValueOf(reflect.MakeChan),
		"MakeFunc":        reflect.ValueOf(reflect.MakeFunc),
		"MakeMap":         reflect.ValueOf(reflect.MakeMap),
		"MakeMapWithSize": reflect.ValueOf(reflect.MakeMapWithSize),
		"MakeSlice":       reflect.ValueOf(reflect.MakeSlice),
		"Map":             reflect.ValueOf(reflect.Map),
		"MapOf":           reflect.ValueOf(reflect.MapOf),
		"New":             reflect.ValueOf(reflect.New),
		"NewAt":           reflect.ValueOf(reflect.NewAt),
		"Pointer":         reflect.ValueOf(reflect.Pointer),
		"PointerTo":       reflect.ValueOf(reflect.PointerTo),
		"Ptr":             reflect.ValueOf(reflect.Ptr),
		"PtrTo":           reflect.ValueOf(reflect.PtrTo),
		"RecvDir":         reflect.ValueOf(reflect.RecvDir),
		"Select":          reflect.ValueOf(reflect.Select),
		"SelectDefault":   reflect.ValueOf(reflect.SelectDefault),
		"SelectRecv":      reflect.ValueOf(reflect.SelectRecv),
		"SelectSend":      reflect.ValueOf(reflect.SelectSend),
		"SendDir":         reflect.ValueOf(reflect.SendDir),
		"Slice":           reflect.ValueOf(reflect.Slice),
		"SliceOf":         reflect.ValueOf(reflect.SliceOf),
		"String":          reflect.ValueOf(reflect.String),
		"Struct":          reflect.ValueOf(reflect.Struct),
		"StructOf":        reflect.ValueOf(reflect.StructOf),
		"Swapper":         reflect.ValueOf(reflect.Swapper),
		"TypeOf":          reflect.ValueOf(reflect.TypeOf),
		"Uint":            reflect.ValueOf(reflect.Uint),
		"Uint16":          reflect.ValueOf(reflect.Uint16),
		"Uint32":          reflect.ValueOf(reflect.Uint32),
		"Uint64":          reflect.ValueOf(reflect.Uint64),
		"Uint8":           reflect.ValueOf(reflect.Uint8),
		"Uintptr":         reflect.ValueOf(reflect.Uintptr),
		"UnsafePointer":   reflect.ValueOf(reflect.UnsafePointer),
		"ValueOf":         reflect.ValueOf(reflect.ValueOf),
		"VisibleFields":   reflect.ValueOf(reflect.VisibleFields),
		"Zero":            reflect.ValueOf(reflect.Zero),

		// type definitions
		"ChanDir":      reflect.ValueOf((*reflect.ChanDir)(nil)),
		"Kind":         reflect.ValueOf((*reflect.Kind)(nil)),
		"MapIter":      reflect.ValueOf((*reflect.MapIter)(nil)),
		"Method":       reflect.ValueOf((*reflect.Method)(nil)),
		"SelectCase":   reflect.ValueOf((*reflect.SelectCase)(nil)),
		"SelectDir":    reflect.ValueOf((*reflect.SelectDir)(nil)),
		"SliceHeader":  reflect.ValueOf((*reflect.SliceHeader)(nil)),
		"StringHeader": reflect.ValueOf((*reflect.StringHeader)(nil)),
		"StructField":  reflect.ValueOf((*reflect.StructField)(nil)),
		"StructTag":    reflect.ValueOf((*reflect.StructTag)(nil)),
		"Type":         reflect.ValueOf((*reflect.Type)(nil)),
		"Value":        reflect.ValueOf((*reflect.Value)(nil)),
		"ValueError":   reflect.ValueOf((*reflect.ValueError)(nil)),

		// interface wrapper definitions
		"_Type": reflect.ValueOf((*_reflect_Type)(nil)),
	}
}

// _reflect_Type is an interface wrapper for Type type
type _reflect_Type struct {
	IValue           interface{}
	WAlign           func() int
	WAssignableTo    func(u reflect.Type) bool
	WBits            func() int
	WChanDir         func() reflect.ChanDir
	WComparable      func() bool
	WConvertibleTo   func(u reflect.Type) bool
	WElem            func() reflect.Type
	WField           func(i int) reflect.StructField
	WFieldAlign      func() int
	WFieldByIndex    func(index []int) reflect.StructField
	WFieldByName     func(name string) (reflect.StructField, bool)
	WFieldByNameFunc func(match func(string) bool) (reflect.StructField, bool)
	WImplements      func(u reflect.Type) bool
	WIn              func(i int) reflect.Type
	WIsVariadic      func() bool
	WKey             func() reflect.Type
	WKind            func() reflect.Kind
	WLen             func() int
	WMethod          func(a0 int) reflect.Method
	WMethodByName    func(a0 string) (reflect.Method, bool)
	WName            func() string
	WNumField        func() int
	WNumIn           func() int
	WNumMethod       func() int
	WNumOut          func() int
	WOut             func(i int) reflect.Type
	WPkgPath         func() string
	WSize            func() uintptr
	WString          func() string
}

func (W _reflect_Type) Align() int {
	return W.WAlign()
}
func (W _reflect_Type) AssignableTo(u reflect.Type) bool {
	return W.WAssignableTo(u)
}
func (W _reflect_Type) Bits() int {
	return W.WBits()
}
func (W _reflect_Type) ChanDir() reflect.ChanDir {
	return W.WChanDir()
}
func (W _reflect_Type) Comparable() bool {
	return W.WComparable()
}
func (W _reflect_Type) ConvertibleTo(u reflect.Type) bool {
	return W.WConvertibleTo(u)
}
func (W _reflect_Type) Elem() reflect.Type {
	return W.WElem()
}
func (W _reflect_Type) Field(i int) reflect.StructField {
	return W.WField(i)
}
func (W _reflect_Type) FieldAlign() int {
	return W.WFieldAlign()
}
func (W _reflect_Type) FieldByIndex(index []int) reflect.StructField {
	return W.WFieldByIndex(index)
}
func (W _reflect_Type) FieldByName(name string) (reflect.StructField, bool) {
	return W.WFieldByName(name)
}
func (W _reflect_Type) FieldByNameFunc(match func(string) bool) (reflect.StructField, bool) {
	return W.WFieldByNameFunc(match)
}
func (W _reflect_Type) Implements(u reflect.Type) bool {
	return W.WImplements(u)
}
func (W _reflect_Type) In(i int) reflect.Type {
	return W.WIn(i)
}
func (W _reflect_Type) IsVariadic() bool {
	return W.WIsVariadic()
}
func (W _reflect_Type) Key() reflect.Type {
	return W.WKey()
}
func (W _reflect_Type) Kind() reflect.Kind {
	return W.WKind()
}
func (W _reflect_Type) Len() int {
	return W.WLen()
}
func (W _reflect_Type) Method(a0 int) reflect.Method {
	return W.WMethod(a0)
}
func (W _reflect_Type) MethodByName(a0 string) (reflect.Method, bool) {
	return W.WMethodByName(a0)
}
func (W _reflect_Type) Name() string {
	return W.WName()
}
func (W _reflect_Type) NumField() int {
	return W.WNumField()
}
func (W _reflect_Type) NumIn() int {
	return W.WNumIn()
}
func (W _reflect_Type) NumMethod() int {
	return W.WNumMethod()
}
func (W _reflect_Type) NumOut() int {
	return W.WNumOut()
}
func (W _reflect_Type) Out(i int) reflect.Type {
	return W.WOut(i)
}
func (W _reflect_Type) PkgPath() string {
	return W.WPkgPath()
}
func (W _reflect_Type) Size() uintptr {
	return W.WSize()
}
func (W _reflect_Type) String() string {
	if W.WString == nil {
		return ""
	}
	return W.WString()
}
