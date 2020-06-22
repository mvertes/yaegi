// Code generated by 'goexports hash'. DO NOT EDIT.

// +build go1.14,!go1.15

package stdlib

import (
	"hash"
	"reflect"

	"github.com/containous/yaegi/interp"
)

func init() {
	Symbols["hash"] = map[string]reflect.Value{
		// type definitions
		"Hash":   reflect.ValueOf((*hash.Hash)(nil)),
		"Hash32": reflect.ValueOf((*hash.Hash32)(nil)),
		"Hash64": reflect.ValueOf((*hash.Hash64)(nil)),
	}
}
func (_w Wrapper) Sum(b []byte) []byte {
	_f := interp.Method("Sum", _w.Wrap).(func(b []byte) []byte)
	return _f(b)
}

func (_w Wrapper) Sum32() uint32 {
	_f := interp.Method("Sum32", _w.Wrap).(func() uint32)
	return _f()
}

func (_w Wrapper) Sum64() uint64 {
	_f := interp.Method("Sum64", _w.Wrap).(func() uint64)
	return _f()
}
