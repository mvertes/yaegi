// Code generated by 'goexports net/rpc'. DO NOT EDIT.

// +build go1.14,!go1.15

package stdlib

import (
	"net/rpc"
	"reflect"

	"github.com/containous/yaegi/interp"
)

func init() {
	Symbols["net/rpc"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Accept":             reflect.ValueOf(rpc.Accept),
		"DefaultDebugPath":   reflect.ValueOf(rpc.DefaultDebugPath),
		"DefaultRPCPath":     reflect.ValueOf(rpc.DefaultRPCPath),
		"DefaultServer":      reflect.ValueOf(&rpc.DefaultServer).Elem(),
		"Dial":               reflect.ValueOf(rpc.Dial),
		"DialHTTP":           reflect.ValueOf(rpc.DialHTTP),
		"DialHTTPPath":       reflect.ValueOf(rpc.DialHTTPPath),
		"ErrShutdown":        reflect.ValueOf(&rpc.ErrShutdown).Elem(),
		"HandleHTTP":         reflect.ValueOf(rpc.HandleHTTP),
		"NewClient":          reflect.ValueOf(rpc.NewClient),
		"NewClientWithCodec": reflect.ValueOf(rpc.NewClientWithCodec),
		"NewServer":          reflect.ValueOf(rpc.NewServer),
		"Register":           reflect.ValueOf(rpc.Register),
		"RegisterName":       reflect.ValueOf(rpc.RegisterName),
		"ServeCodec":         reflect.ValueOf(rpc.ServeCodec),
		"ServeConn":          reflect.ValueOf(rpc.ServeConn),
		"ServeRequest":       reflect.ValueOf(rpc.ServeRequest),

		// type definitions
		"Call":        reflect.ValueOf((*rpc.Call)(nil)),
		"Client":      reflect.ValueOf((*rpc.Client)(nil)),
		"ClientCodec": reflect.ValueOf((*rpc.ClientCodec)(nil)),
		"Request":     reflect.ValueOf((*rpc.Request)(nil)),
		"Response":    reflect.ValueOf((*rpc.Response)(nil)),
		"Server":      reflect.ValueOf((*rpc.Server)(nil)),
		"ServerCodec": reflect.ValueOf((*rpc.ServerCodec)(nil)),
		"ServerError": reflect.ValueOf((*rpc.ServerError)(nil)),
	}
}
func (_w Wrapper) ReadResponseBody(a0 interface{}) error {
	_f := interp.Method("ReadResponseBody", _w.Wrap).(func(a0 interface{}) error)
	return _f(a0)
}
func (_w Wrapper) ReadResponseHeader(a0 *rpc.Response) error {
	_f := interp.Method("ReadResponseHeader", _w.Wrap).(func(a0 *rpc.Response) error)
	return _f(a0)
}
func (_w Wrapper) WriteRequest(a0 *rpc.Request, a1 interface{}) error {
	_f := interp.Method("WriteRequest", _w.Wrap).(func(a0 *rpc.Request, a1 interface{}) error)
	return _f(a0, a1)
}

func (_w Wrapper) ReadRequestBody(a0 interface{}) error {
	_f := interp.Method("ReadRequestBody", _w.Wrap).(func(a0 interface{}) error)
	return _f(a0)
}
func (_w Wrapper) ReadRequestHeader(a0 *rpc.Request) error {
	_f := interp.Method("ReadRequestHeader", _w.Wrap).(func(a0 *rpc.Request) error)
	return _f(a0)
}
func (_w Wrapper) WriteResponse(a0 *rpc.Response, a1 interface{}) error {
	_f := interp.Method("WriteResponse", _w.Wrap).(func(a0 *rpc.Response, a1 interface{}) error)
	return _f(a0, a1)
}
