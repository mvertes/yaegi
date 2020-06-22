// Code generated by 'goexports sync'. DO NOT EDIT.

// +build go1.14,!go1.15

package stdlib

import (
	"reflect"
	"sync"

	"github.com/containous/yaegi/interp"
)

func init() {
	Symbols["sync"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"NewCond": reflect.ValueOf(sync.NewCond),

		// type definitions
		"Cond":      reflect.ValueOf((*sync.Cond)(nil)),
		"Locker":    reflect.ValueOf((*sync.Locker)(nil)),
		"Map":       reflect.ValueOf((*sync.Map)(nil)),
		"Mutex":     reflect.ValueOf((*sync.Mutex)(nil)),
		"Once":      reflect.ValueOf((*sync.Once)(nil)),
		"Pool":      reflect.ValueOf((*sync.Pool)(nil)),
		"RWMutex":   reflect.ValueOf((*sync.RWMutex)(nil)),
		"WaitGroup": reflect.ValueOf((*sync.WaitGroup)(nil)),
	}
}
func (_w Wrapper) Lock() {
	_f := interp.Method("Lock", _w.Wrap).(func())
	_f()
}
func (_w Wrapper) Unlock() {
	_f := interp.Method("Unlock", _w.Wrap).(func())
	_f()
}
