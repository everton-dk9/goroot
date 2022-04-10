package goroot

import (
	"testing"
)

func TestSecantMethod(t *testing.T) {

	testsCases := []struct {
		description string
		function    fn
		x0          float64
		x1          float64
		iterations  int
		precision   float64
		expected    float64
	}{
		{
			description: "Fixed-Point->x^3-9x+3",
			function:    func(x float64) float64 { return ((x * x * x) - (9 * x) + 3) },
			x0:          0.3,
			x1:          0.35,
			iterations:  20,
			precision:   1e-4,
			expected:    0.3376089,
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.description, func(t *testing.T) {
			m := &MethodBase{
				f:  tc.function,
				x0: tc.x0,
				x1: tc.x1,
				k:  tc.iterations,
				e1: tc.precision,
				e2: tc.precision,
			}

			got := Round(Secante(m), 7)

			if got != tc.expected {
				t.Errorf("got %f but expected %f", got, tc.expected)
			}
		})
	}
}
