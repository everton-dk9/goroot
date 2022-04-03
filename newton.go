package goroot

import (
	"math"
)

// metodo numerico de Newton-Rhapson:
//
// entradas:
// 	- função f
// 	- derivada de f
// 	- valor inicial x0
// 	- precisões e1, e2
// 	- qtd de iterações k
//
// saida:
// 	- raiz aproximada de f
//
func Newton(mb *MethodBase) float64 {
	f := mb.f
	x0 := mb.x0
	df := mb.g
	e1 := mb.e1
	e2 := mb.e2

	if math.Abs(f(x0)) < e1 {
		return x0
	}

	var xn float64
	for i := 1; i <= mb.k; i++ {
		xn = x0 - (f(x0) / df(x0))

		if math.Abs(f(xn)) < e1 || math.Abs(xn-x0) < e2 {
			mb.Print(i, xn, f(xn))
			return xn
		}

		mb.Print(i, xn, f(xn))
		x0 = xn
	}
	return xn
}
