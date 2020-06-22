// Code generated by 'goexports image/draw'. DO NOT EDIT.

// +build go1.14,!go1.15

package stdlib

import (
	"image"
	"image/color"
	"image/draw"
	"reflect"

	"github.com/containous/yaegi/interp"
)

func init() {
	Symbols["image/draw"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Draw":           reflect.ValueOf(draw.Draw),
		"DrawMask":       reflect.ValueOf(draw.DrawMask),
		"FloydSteinberg": reflect.ValueOf(&draw.FloydSteinberg).Elem(),
		"Over":           reflect.ValueOf(draw.Over),
		"Src":            reflect.ValueOf(draw.Src),

		// type definitions
		"Drawer":    reflect.ValueOf((*draw.Drawer)(nil)),
		"Image":     reflect.ValueOf((*draw.Image)(nil)),
		"Op":        reflect.ValueOf((*draw.Op)(nil)),
		"Quantizer": reflect.ValueOf((*draw.Quantizer)(nil)),
	}
}
func (_w Wrapper) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	_f := interp.Method("Draw", _w.Wrap).(func(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point))
	_f(dst, r, src, sp)
}

func (_w Wrapper) Quantize(p color.Palette, m image.Image) color.Palette {
	_f := interp.Method("Quantize", _w.Wrap).(func(p color.Palette, m image.Image) color.Palette)
	return _f(p, m)
}
