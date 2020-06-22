// Code generated by 'goexports encoding/xml'. DO NOT EDIT.

// +build go1.14,!go1.15

package stdlib

import (
	"encoding/xml"
	"reflect"

	"github.com/containous/yaegi/interp"
)

func init() {
	Symbols["encoding/xml"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"CopyToken":       reflect.ValueOf(xml.CopyToken),
		"Escape":          reflect.ValueOf(xml.Escape),
		"EscapeText":      reflect.ValueOf(xml.EscapeText),
		"HTMLAutoClose":   reflect.ValueOf(&xml.HTMLAutoClose).Elem(),
		"HTMLEntity":      reflect.ValueOf(&xml.HTMLEntity).Elem(),
		"Header":          reflect.ValueOf(xml.Header),
		"Marshal":         reflect.ValueOf(xml.Marshal),
		"MarshalIndent":   reflect.ValueOf(xml.MarshalIndent),
		"NewDecoder":      reflect.ValueOf(xml.NewDecoder),
		"NewEncoder":      reflect.ValueOf(xml.NewEncoder),
		"NewTokenDecoder": reflect.ValueOf(xml.NewTokenDecoder),
		"Unmarshal":       reflect.ValueOf(xml.Unmarshal),

		// type definitions
		"Attr":                 reflect.ValueOf((*xml.Attr)(nil)),
		"CharData":             reflect.ValueOf((*xml.CharData)(nil)),
		"Comment":              reflect.ValueOf((*xml.Comment)(nil)),
		"Decoder":              reflect.ValueOf((*xml.Decoder)(nil)),
		"Directive":            reflect.ValueOf((*xml.Directive)(nil)),
		"Encoder":              reflect.ValueOf((*xml.Encoder)(nil)),
		"EndElement":           reflect.ValueOf((*xml.EndElement)(nil)),
		"Marshaler":            reflect.ValueOf((*xml.Marshaler)(nil)),
		"MarshalerAttr":        reflect.ValueOf((*xml.MarshalerAttr)(nil)),
		"Name":                 reflect.ValueOf((*xml.Name)(nil)),
		"ProcInst":             reflect.ValueOf((*xml.ProcInst)(nil)),
		"StartElement":         reflect.ValueOf((*xml.StartElement)(nil)),
		"SyntaxError":          reflect.ValueOf((*xml.SyntaxError)(nil)),
		"TagPathError":         reflect.ValueOf((*xml.TagPathError)(nil)),
		"Token":                reflect.ValueOf((*xml.Token)(nil)),
		"TokenReader":          reflect.ValueOf((*xml.TokenReader)(nil)),
		"UnmarshalError":       reflect.ValueOf((*xml.UnmarshalError)(nil)),
		"Unmarshaler":          reflect.ValueOf((*xml.Unmarshaler)(nil)),
		"UnmarshalerAttr":      reflect.ValueOf((*xml.UnmarshalerAttr)(nil)),
		"UnsupportedTypeError": reflect.ValueOf((*xml.UnsupportedTypeError)(nil)),
	}
}
func (_w Wrapper) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	_f := interp.Method("MarshalXML", _w.Wrap).(func(e *xml.Encoder, start xml.StartElement) error)
	return _f(e, start)
}

func (_w Wrapper) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	_f := interp.Method("MarshalXMLAttr", _w.Wrap).(func(name xml.Name) (xml.Attr, error))
	return _f(name)
}

func (_w Wrapper) Token() (xml.Token, error) {
	_f := interp.Method("Token", _w.Wrap).(func() (xml.Token, error))
	return _f()
}

func (_w Wrapper) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	_f := interp.Method("UnmarshalXML", _w.Wrap).(func(d *xml.Decoder, start xml.StartElement) error)
	return _f(d, start)
}

func (_w Wrapper) UnmarshalXMLAttr(attr xml.Attr) error {
	_f := interp.Method("UnmarshalXMLAttr", _w.Wrap).(func(attr xml.Attr) error)
	return _f(attr)
}
