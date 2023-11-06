// Code generated by 'yaegi extract io'. DO NOT EDIT.

//go:build go1.20 && !go1.21
// +build go1.20

package stdlib

import (
	"go/constant"
	"go/token"
	"io"
	"reflect"
)

func init() {
	Symbols["io/io"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Copy":             reflect.ValueOf(io.Copy),
		"CopyBuffer":       reflect.ValueOf(io.CopyBuffer),
		"CopyN":            reflect.ValueOf(io.CopyN),
		"Discard":          reflect.ValueOf(&io.Discard).Elem(),
		"EOF":              reflect.ValueOf(&io.EOF).Elem(),
		"ErrClosedPipe":    reflect.ValueOf(&io.ErrClosedPipe).Elem(),
		"ErrNoProgress":    reflect.ValueOf(&io.ErrNoProgress).Elem(),
		"ErrShortBuffer":   reflect.ValueOf(&io.ErrShortBuffer).Elem(),
		"ErrShortWrite":    reflect.ValueOf(&io.ErrShortWrite).Elem(),
		"ErrUnexpectedEOF": reflect.ValueOf(&io.ErrUnexpectedEOF).Elem(),
		"LimitReader":      reflect.ValueOf(io.LimitReader),
		"MultiReader":      reflect.ValueOf(io.MultiReader),
		"MultiWriter":      reflect.ValueOf(io.MultiWriter),
		"NewOffsetWriter":  reflect.ValueOf(io.NewOffsetWriter),
		"NewSectionReader": reflect.ValueOf(io.NewSectionReader),
		"NopCloser":        reflect.ValueOf(io.NopCloser),
		"Pipe":             reflect.ValueOf(io.Pipe),
		"ReadAll":          reflect.ValueOf(io.ReadAll),
		"ReadAtLeast":      reflect.ValueOf(io.ReadAtLeast),
		"ReadFull":         reflect.ValueOf(io.ReadFull),
		"SeekCurrent":      reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"SeekEnd":          reflect.ValueOf(constant.MakeFromLiteral("2", token.INT, 0)),
		"SeekStart":        reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),
		"TeeReader":        reflect.ValueOf(io.TeeReader),
		"WriteString":      reflect.ValueOf(io.WriteString),

		// type definitions
		"ByteReader":      reflect.ValueOf((*io.ByteReader)(nil)),
		"ByteScanner":     reflect.ValueOf((*io.ByteScanner)(nil)),
		"ByteWriter":      reflect.ValueOf((*io.ByteWriter)(nil)),
		"Closer":          reflect.ValueOf((*io.Closer)(nil)),
		"LimitedReader":   reflect.ValueOf((*io.LimitedReader)(nil)),
		"OffsetWriter":    reflect.ValueOf((*io.OffsetWriter)(nil)),
		"PipeReader":      reflect.ValueOf((*io.PipeReader)(nil)),
		"PipeWriter":      reflect.ValueOf((*io.PipeWriter)(nil)),
		"ReadCloser":      reflect.ValueOf((*io.ReadCloser)(nil)),
		"ReadSeekCloser":  reflect.ValueOf((*io.ReadSeekCloser)(nil)),
		"ReadSeeker":      reflect.ValueOf((*io.ReadSeeker)(nil)),
		"ReadWriteCloser": reflect.ValueOf((*io.ReadWriteCloser)(nil)),
		"ReadWriteSeeker": reflect.ValueOf((*io.ReadWriteSeeker)(nil)),
		"ReadWriter":      reflect.ValueOf((*io.ReadWriter)(nil)),
		"Reader":          reflect.ValueOf((*io.Reader)(nil)),
		"ReaderAt":        reflect.ValueOf((*io.ReaderAt)(nil)),
		"ReaderFrom":      reflect.ValueOf((*io.ReaderFrom)(nil)),
		"RuneReader":      reflect.ValueOf((*io.RuneReader)(nil)),
		"RuneScanner":     reflect.ValueOf((*io.RuneScanner)(nil)),
		"SectionReader":   reflect.ValueOf((*io.SectionReader)(nil)),
		"Seeker":          reflect.ValueOf((*io.Seeker)(nil)),
		"StringWriter":    reflect.ValueOf((*io.StringWriter)(nil)),
		"WriteCloser":     reflect.ValueOf((*io.WriteCloser)(nil)),
		"WriteSeeker":     reflect.ValueOf((*io.WriteSeeker)(nil)),
		"Writer":          reflect.ValueOf((*io.Writer)(nil)),
		"WriterAt":        reflect.ValueOf((*io.WriterAt)(nil)),
		"WriterTo":        reflect.ValueOf((*io.WriterTo)(nil)),

		// interface wrapper definitions
		"_ByteReader":      reflect.ValueOf((*_io_ByteReader)(nil)),
		"_ByteScanner":     reflect.ValueOf((*_io_ByteScanner)(nil)),
		"_ByteWriter":      reflect.ValueOf((*_io_ByteWriter)(nil)),
		"_Closer":          reflect.ValueOf((*_io_Closer)(nil)),
		"_ReadCloser":      reflect.ValueOf((*_io_ReadCloser)(nil)),
		"_ReadSeekCloser":  reflect.ValueOf((*_io_ReadSeekCloser)(nil)),
		"_ReadSeeker":      reflect.ValueOf((*_io_ReadSeeker)(nil)),
		"_ReadWriteCloser": reflect.ValueOf((*_io_ReadWriteCloser)(nil)),
		"_ReadWriteSeeker": reflect.ValueOf((*_io_ReadWriteSeeker)(nil)),
		"_ReadWriter":      reflect.ValueOf((*_io_ReadWriter)(nil)),
		"_Reader":          reflect.ValueOf((*_io_Reader)(nil)),
		"_ReaderAt":        reflect.ValueOf((*_io_ReaderAt)(nil)),
		"_ReaderFrom":      reflect.ValueOf((*_io_ReaderFrom)(nil)),
		"_RuneReader":      reflect.ValueOf((*_io_RuneReader)(nil)),
		"_RuneScanner":     reflect.ValueOf((*_io_RuneScanner)(nil)),
		"_Seeker":          reflect.ValueOf((*_io_Seeker)(nil)),
		"_StringWriter":    reflect.ValueOf((*_io_StringWriter)(nil)),
		"_WriteCloser":     reflect.ValueOf((*_io_WriteCloser)(nil)),
		"_WriteSeeker":     reflect.ValueOf((*_io_WriteSeeker)(nil)),
		"_Writer":          reflect.ValueOf((*_io_Writer)(nil)),
		"_WriterAt":        reflect.ValueOf((*_io_WriterAt)(nil)),
		"_WriterTo":        reflect.ValueOf((*_io_WriterTo)(nil)),
	}
}

// _io_ByteReader is an interface wrapper for ByteReader type
type _io_ByteReader struct {
	IValue    interface{}
	WReadByte func() (byte, error)
}

func (W _io_ByteReader) ReadByte() (byte, error) {
	return W.WReadByte()
}

// _io_ByteScanner is an interface wrapper for ByteScanner type
type _io_ByteScanner struct {
	IValue      interface{}
	WReadByte   func() (byte, error)
	WUnreadByte func() error
}

func (W _io_ByteScanner) ReadByte() (byte, error) {
	return W.WReadByte()
}
func (W _io_ByteScanner) UnreadByte() error {
	return W.WUnreadByte()
}

// _io_ByteWriter is an interface wrapper for ByteWriter type
type _io_ByteWriter struct {
	IValue     interface{}
	WWriteByte func(c byte) error
}

func (W _io_ByteWriter) WriteByte(c byte) error {
	return W.WWriteByte(c)
}

// _io_Closer is an interface wrapper for Closer type
type _io_Closer struct {
	IValue interface{}
	WClose func() error
}

func (W _io_Closer) Close() error {
	return W.WClose()
}

// _io_ReadCloser is an interface wrapper for ReadCloser type
type _io_ReadCloser struct {
	IValue interface{}
	WClose func() error
	WRead  func(p []byte) (n int, err error)
}

func (W _io_ReadCloser) Close() error {
	return W.WClose()
}
func (W _io_ReadCloser) Read(p []byte) (n int, err error) {
	return W.WRead(p)
}

// _io_ReadSeekCloser is an interface wrapper for ReadSeekCloser type
type _io_ReadSeekCloser struct {
	IValue interface{}
	WClose func() error
	WRead  func(p []byte) (n int, err error)
	WSeek  func(offset int64, whence int) (int64, error)
}

func (W _io_ReadSeekCloser) Close() error {
	return W.WClose()
}
func (W _io_ReadSeekCloser) Read(p []byte) (n int, err error) {
	return W.WRead(p)
}
func (W _io_ReadSeekCloser) Seek(offset int64, whence int) (int64, error) {
	return W.WSeek(offset, whence)
}

// _io_ReadSeeker is an interface wrapper for ReadSeeker type
type _io_ReadSeeker struct {
	IValue interface{}
	WRead  func(p []byte) (n int, err error)
	WSeek  func(offset int64, whence int) (int64, error)
}

func (W _io_ReadSeeker) Read(p []byte) (n int, err error) {
	return W.WRead(p)
}
func (W _io_ReadSeeker) Seek(offset int64, whence int) (int64, error) {
	return W.WSeek(offset, whence)
}

// _io_ReadWriteCloser is an interface wrapper for ReadWriteCloser type
type _io_ReadWriteCloser struct {
	IValue interface{}
	WClose func() error
	WRead  func(p []byte) (n int, err error)
	WWrite func(p []byte) (n int, err error)
}

func (W _io_ReadWriteCloser) Close() error {
	return W.WClose()
}
func (W _io_ReadWriteCloser) Read(p []byte) (n int, err error) {
	return W.WRead(p)
}
func (W _io_ReadWriteCloser) Write(p []byte) (n int, err error) {
	return W.WWrite(p)
}

// _io_ReadWriteSeeker is an interface wrapper for ReadWriteSeeker type
type _io_ReadWriteSeeker struct {
	IValue interface{}
	WRead  func(p []byte) (n int, err error)
	WSeek  func(offset int64, whence int) (int64, error)
	WWrite func(p []byte) (n int, err error)
}

func (W _io_ReadWriteSeeker) Read(p []byte) (n int, err error) {
	return W.WRead(p)
}
func (W _io_ReadWriteSeeker) Seek(offset int64, whence int) (int64, error) {
	return W.WSeek(offset, whence)
}
func (W _io_ReadWriteSeeker) Write(p []byte) (n int, err error) {
	return W.WWrite(p)
}

// _io_ReadWriter is an interface wrapper for ReadWriter type
type _io_ReadWriter struct {
	IValue interface{}
	WRead  func(p []byte) (n int, err error)
	WWrite func(p []byte) (n int, err error)
}

func (W _io_ReadWriter) Read(p []byte) (n int, err error) {
	return W.WRead(p)
}
func (W _io_ReadWriter) Write(p []byte) (n int, err error) {
	return W.WWrite(p)
}

// _io_Reader is an interface wrapper for Reader type
type _io_Reader struct {
	IValue interface{}
	WRead  func(p []byte) (n int, err error)
}

func (W _io_Reader) Read(p []byte) (n int, err error) {
	return W.WRead(p)
}

// _io_ReaderAt is an interface wrapper for ReaderAt type
type _io_ReaderAt struct {
	IValue  interface{}
	WReadAt func(p []byte, off int64) (n int, err error)
}

func (W _io_ReaderAt) ReadAt(p []byte, off int64) (n int, err error) {
	return W.WReadAt(p, off)
}

// _io_ReaderFrom is an interface wrapper for ReaderFrom type
type _io_ReaderFrom struct {
	IValue    interface{}
	WReadFrom func(r io.Reader) (n int64, err error)
}

func (W _io_ReaderFrom) ReadFrom(r io.Reader) (n int64, err error) {
	return W.WReadFrom(r)
}

// _io_RuneReader is an interface wrapper for RuneReader type
type _io_RuneReader struct {
	IValue    interface{}
	WReadRune func() (r rune, size int, err error)
}

func (W _io_RuneReader) ReadRune() (r rune, size int, err error) {
	return W.WReadRune()
}

// _io_RuneScanner is an interface wrapper for RuneScanner type
type _io_RuneScanner struct {
	IValue      interface{}
	WReadRune   func() (r rune, size int, err error)
	WUnreadRune func() error
}

func (W _io_RuneScanner) ReadRune() (r rune, size int, err error) {
	return W.WReadRune()
}
func (W _io_RuneScanner) UnreadRune() error {
	return W.WUnreadRune()
}

// _io_Seeker is an interface wrapper for Seeker type
type _io_Seeker struct {
	IValue interface{}
	WSeek  func(offset int64, whence int) (int64, error)
}

func (W _io_Seeker) Seek(offset int64, whence int) (int64, error) {
	return W.WSeek(offset, whence)
}

// _io_StringWriter is an interface wrapper for StringWriter type
type _io_StringWriter struct {
	IValue       interface{}
	WWriteString func(s string) (n int, err error)
}

func (W _io_StringWriter) WriteString(s string) (n int, err error) {
	return W.WWriteString(s)
}

// _io_WriteCloser is an interface wrapper for WriteCloser type
type _io_WriteCloser struct {
	IValue interface{}
	WClose func() error
	WWrite func(p []byte) (n int, err error)
}

func (W _io_WriteCloser) Close() error {
	return W.WClose()
}
func (W _io_WriteCloser) Write(p []byte) (n int, err error) {
	return W.WWrite(p)
}

// _io_WriteSeeker is an interface wrapper for WriteSeeker type
type _io_WriteSeeker struct {
	IValue interface{}
	WSeek  func(offset int64, whence int) (int64, error)
	WWrite func(p []byte) (n int, err error)
}

func (W _io_WriteSeeker) Seek(offset int64, whence int) (int64, error) {
	return W.WSeek(offset, whence)
}
func (W _io_WriteSeeker) Write(p []byte) (n int, err error) {
	return W.WWrite(p)
}

// _io_Writer is an interface wrapper for Writer type
type _io_Writer struct {
	IValue interface{}
	WWrite func(p []byte) (n int, err error)
}

func (W _io_Writer) Write(p []byte) (n int, err error) {
	return W.WWrite(p)
}

// _io_WriterAt is an interface wrapper for WriterAt type
type _io_WriterAt struct {
	IValue   interface{}
	WWriteAt func(p []byte, off int64) (n int, err error)
}

func (W _io_WriterAt) WriteAt(p []byte, off int64) (n int, err error) {
	return W.WWriteAt(p, off)
}

// _io_WriterTo is an interface wrapper for WriterTo type
type _io_WriterTo struct {
	IValue   interface{}
	WWriteTo func(w io.Writer) (n int64, err error)
}

func (W _io_WriterTo) WriteTo(w io.Writer) (n int64, err error) {
	return W.WWriteTo(w)
}
