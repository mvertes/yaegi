// Code generated by 'goexports context'. DO NOT EDIT.

// +build go1.14,!go1.15

package stdlib

import (
	"context"
	"reflect"
	"time"

	"github.com/containous/yaegi/interp"
)

func init() {
	Symbols["context"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Background":       reflect.ValueOf(context.Background),
		"Canceled":         reflect.ValueOf(&context.Canceled).Elem(),
		"DeadlineExceeded": reflect.ValueOf(&context.DeadlineExceeded).Elem(),
		"TODO":             reflect.ValueOf(context.TODO),
		"WithCancel":       reflect.ValueOf(context.WithCancel),
		"WithDeadline":     reflect.ValueOf(context.WithDeadline),
		"WithTimeout":      reflect.ValueOf(context.WithTimeout),
		"WithValue":        reflect.ValueOf(context.WithValue),

		// type definitions
		"CancelFunc": reflect.ValueOf((*context.CancelFunc)(nil)),
		"Context":    reflect.ValueOf((*context.Context)(nil)),
	}
}
func (_w Wrapper) Deadline() (deadline time.Time, ok bool) {
	_f := interp.Method("Deadline", _w.Wrap).(func() (deadline time.Time, ok bool))
	return _f()
}
func (_w Wrapper) Done() <-chan struct{} {
	_f := interp.Method("Done", _w.Wrap).(func() <-chan struct{})
	return _f()
}
func (_w Wrapper) Err() error {
	_f := interp.Method("Err", _w.Wrap).(func() error)
	return _f()
}
func (_w Wrapper) Value(key interface{}) interface{} {
	_f := interp.Method("Value", _w.Wrap).(func(key interface{}) interface{})
	return _f(key)
}
