package stdlib

import (
	"reflect"
)

func init() {
	Symbols["wrapper"] = map[string]reflect.Value{
		"Wrapper": reflect.ValueOf((*Wrapper)(nil)),
	}
}

// Wrapper defines a generic interface wrapper to call methods in yaegi interpreter.
type Wrapper struct {
	Wrap reflect.Value
}

/*
func (w Wrapper) String() (s string) {
	f := interp.Method("String", w.Value).Interface().(func() string)
	return f()
}

func (w Wrapper) Read(p []byte) (n int, err error) {
	f := interp.Method("Read", w.Value).Interface().(func(p []byte) (n int, err error))
	return f(p)
}
*/
