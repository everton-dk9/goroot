package goroot

import (
	"testing"
)

func TestMetodoBissecao(t *testing.T) {

	testsCases := []struct {
		description string
		function    func(x float64) float64
		a, b        float64
		precision   float64
		iterations  int
		expected    float64
	}{
		{
			description: "Bisection->x^3-9x+3",
			function:    func(x float64) float64 { return ((x * x * x) - (9 * x) + 3) },
			a:           0,
			b:           1,
			precision:   0.001,
			iterations:  16,
			expected:    0.337402344,
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.description, func(t *testing.T) {
			m := &MethodBase{
				k:  tc.iterations,
				f:  tc.function,
				a:  tc.a,
				b:  tc.b,
				e1: tc.precision,
			}

			got := Round(Bissecao(m), 9)

			if got != tc.expected {
				t.Errorf("got %f but expected %f", got, tc.expected)
			}

		})
	}

}
