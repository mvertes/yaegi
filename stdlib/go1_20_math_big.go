// Code generated by 'yaegi extract math/big'. DO NOT EDIT.

//go:build go1.20 && !go1.21
// +build go1.20

package stdlib

import (
	"go/constant"
	"go/token"
	"math/big"
	"reflect"
)

func init() {
	Symbols["math/big/big"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Above":         reflect.ValueOf(big.Above),
		"AwayFromZero":  reflect.ValueOf(big.AwayFromZero),
		"Below":         reflect.ValueOf(big.Below),
		"Exact":         reflect.ValueOf(big.Exact),
		"Jacobi":        reflect.ValueOf(big.Jacobi),
		"MaxBase":       reflect.ValueOf(constant.MakeFromLiteral("62", token.INT, 0)),
		"MaxExp":        reflect.ValueOf(constant.MakeFromLiteral("2147483647", token.INT, 0)),
		"MaxPrec":       reflect.ValueOf(constant.MakeFromLiteral("4294967295", token.INT, 0)),
		"MinExp":        reflect.ValueOf(constant.MakeFromLiteral("-2147483648", token.INT, 0)),
		"NewFloat":      reflect.ValueOf(big.NewFloat),
		"NewInt":        reflect.ValueOf(big.NewInt),
		"NewRat":        reflect.ValueOf(big.NewRat),
		"ParseFloat":    reflect.ValueOf(big.ParseFloat),
		"ToNearestAway": reflect.ValueOf(big.ToNearestAway),
		"ToNearestEven": reflect.ValueOf(big.ToNearestEven),
		"ToNegativeInf": reflect.ValueOf(big.ToNegativeInf),
		"ToPositiveInf": reflect.ValueOf(big.ToPositiveInf),
		"ToZero":        reflect.ValueOf(big.ToZero),

		// type definitions
		"Accuracy":     reflect.ValueOf((*big.Accuracy)(nil)),
		"ErrNaN":       reflect.ValueOf((*big.ErrNaN)(nil)),
		"Float":        reflect.ValueOf((*big.Float)(nil)),
		"Int":          reflect.ValueOf((*big.Int)(nil)),
		"Rat":          reflect.ValueOf((*big.Rat)(nil)),
		"RoundingMode": reflect.ValueOf((*big.RoundingMode)(nil)),
		"Word":         reflect.ValueOf((*big.Word)(nil)),
	}
}
