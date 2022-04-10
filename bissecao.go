package goroot

import (
	"fmt"
)

// input: f, [a,b], precision
func Bissecao(mb *MethodBase) float64 {

	a := mb.a
	b := mb.b
	f := mb.f
	e1 := mb.e1
	if f(a)*f(b) >= 0 {
		fmt.Println("o intervalo não é válido escolha outro valor para a e b")
		return 0.0
	}

	var xn float64
	for i := 1; i <= mb.k; i++ {
		xn = (a + b) / 2

		if (b - a) < e1 {
			mb.Print(i, xn, f(xn), (b - a))
			return xn
		}

		mb.Print(i, xn, f(xn), (b - a))
		if f(a)*f(xn) > 0 {
			a = xn
		} else {
			b = xn
		}

	}
	fmt.Println("limite de  iteracoes atingido")
	return xn
}
