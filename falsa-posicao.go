package goroot

import (
	"fmt"
	"math"
)

// input: f(x),[a,b], precision
func FalsaPosicao(mb *MethodBase) float64 {
	var root float64

	a := mb.a
	b := mb.b
	f := mb.f
	e1 := mb.e1
	if f(a)*f(b) >= 0 {
		fmt.Println("o intervalo não é válido escolha outro valor para a e b")
		return 0.0
	}

	for i := 1; i <= mb.k; i++ {

		if (b - a) < e1 {
			mb.Print(i, root, f(root), (b - a))
			return root
		}

		root = ((a * f(b)) - (b * f(a))) / (f(b) - f(a))

		if math.Abs(f(a)) < e1 {
			mb.Print(i, root, f(root), (b - a))
			return a
		}

		if math.Abs(f(b)) < e1 {
			mb.Print(i, root, f(root), (b - a))
			return b
		}

		if math.Abs(f(root)) < e1 {
			mb.Print(i, root, f(root), (b - a))
			return root
		}

		mb.Print(i, root, f(root), (b - a))

		if f(a)*f(root) > 0 {
			a = root
		} else {
			b = root
		}
	}
	return root
}
