package goroot

import (
	"math"
)

// input: f(x), initialValue, iterationFunc(x), precision1, precision2
func MPF(mb *MethodBase) float64 {
	f := mb.f
	q := mb.g
	x0 := mb.x0
	e1 := mb.e1
	e2 := mb.e2

	if math.Abs(f(x0)) < e1 {
		return x0
	}

	var xn float64
	for i := 1; i <= mb.k; i++ {
		xn := q(x0)

		if math.Abs(f(xn)) < e1 || math.Abs(xn-x0) < e2 {
			mb.Print(i, xn, f(xn))
			return xn
		}

		mb.Print(i, xn, f(xn))
		x0 = xn
	}

	return xn
}
