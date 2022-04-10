package goroot

import (
	"math"
)

// input: f(x), initialValueX0, initialValueX1 precision1, precision2
func Secante(mb *MethodBase) float64 {
	f := mb.f
	x0 := mb.x0
	x1 := mb.x1
	e1 := mb.e1
	e2 := mb.e2

	if math.Abs(f(x0)) < e1 {
		return x0
	}

	if math.Abs(f(x1)) < e1 || math.Abs(x1-x0) < e2 {
		return x1
	}

	var xn float64
	for i := 0; i <= mb.k; i++ {

		xn = x1 - (f(x1)/(f(x1)-f(x0)))*(x1-x0)

		mb.Print(i, xn, f(xn))

		if math.Abs(f(xn)) < e1 || math.Abs(xn-x1) < e2 {
			return xn
		}

		x0 = x1
		x1 = xn
	}
	return xn
}
