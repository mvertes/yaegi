// Code generated by 'yaegi extract net/textproto'. DO NOT EDIT.

//go:build go1.20 && !go1.21
// +build go1.20

package stdlib

import (
	"net/textproto"
	"reflect"
)

func init() {
	Symbols["net/textproto/textproto"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"CanonicalMIMEHeaderKey": reflect.ValueOf(textproto.CanonicalMIMEHeaderKey),
		"Dial":                   reflect.ValueOf(textproto.Dial),
		"NewConn":                reflect.ValueOf(textproto.NewConn),
		"NewReader":              reflect.ValueOf(textproto.NewReader),
		"NewWriter":              reflect.ValueOf(textproto.NewWriter),
		"TrimBytes":              reflect.ValueOf(textproto.TrimBytes),
		"TrimString":             reflect.ValueOf(textproto.TrimString),

		// type definitions
		"Conn":          reflect.ValueOf((*textproto.Conn)(nil)),
		"Error":         reflect.ValueOf((*textproto.Error)(nil)),
		"MIMEHeader":    reflect.ValueOf((*textproto.MIMEHeader)(nil)),
		"Pipeline":      reflect.ValueOf((*textproto.Pipeline)(nil)),
		"ProtocolError": reflect.ValueOf((*textproto.ProtocolError)(nil)),
		"Reader":        reflect.ValueOf((*textproto.Reader)(nil)),
		"Writer":        reflect.ValueOf((*textproto.Writer)(nil)),
	}
}
