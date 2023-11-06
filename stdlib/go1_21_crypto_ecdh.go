// Code generated by 'yaegi extract crypto/ecdh'. DO NOT EDIT.

//go:build go1.21
// +build go1.21

package stdlib

import (
	"crypto/ecdh"
	"io"
	"reflect"
)

func init() {
	Symbols["crypto/ecdh/ecdh"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"P256":   reflect.ValueOf(ecdh.P256),
		"P384":   reflect.ValueOf(ecdh.P384),
		"P521":   reflect.ValueOf(ecdh.P521),
		"X25519": reflect.ValueOf(ecdh.X25519),

		// type definitions
		"Curve":      reflect.ValueOf((*ecdh.Curve)(nil)),
		"PrivateKey": reflect.ValueOf((*ecdh.PrivateKey)(nil)),
		"PublicKey":  reflect.ValueOf((*ecdh.PublicKey)(nil)),

		// interface wrapper definitions
		"_Curve": reflect.ValueOf((*_crypto_ecdh_Curve)(nil)),
	}
}

// _crypto_ecdh_Curve is an interface wrapper for Curve type
type _crypto_ecdh_Curve struct {
	IValue         interface{}
	WGenerateKey   func(rand io.Reader) (*ecdh.PrivateKey, error)
	WNewPrivateKey func(key []byte) (*ecdh.PrivateKey, error)
	WNewPublicKey  func(key []byte) (*ecdh.PublicKey, error)
}

func (W _crypto_ecdh_Curve) GenerateKey(rand io.Reader) (*ecdh.PrivateKey, error) {
	return W.WGenerateKey(rand)
}
func (W _crypto_ecdh_Curve) NewPrivateKey(key []byte) (*ecdh.PrivateKey, error) {
	return W.WNewPrivateKey(key)
}
func (W _crypto_ecdh_Curve) NewPublicKey(key []byte) (*ecdh.PublicKey, error) {
	return W.WNewPublicKey(key)
}
