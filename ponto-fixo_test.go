package goroot

import (
	"testing"
)

func TestMPFMethod(t *testing.T) {

	testsCases := []struct {
		description string
		function    fn
		qFunc       fn
		x0          float64
		iterations  int
		precision   float64
		expected    float64
	}{
		{
			description: "Fixed-Point->x^3-9x+3",
			function:    func(x float64) float64 { return ((x * x * x) - (9 * x) + 3) },
			qFunc:       func(x float64) float64 { return ((x * x * x) / 9.0) + (1.0 / 3.0) },
			x0:          0.5,
			iterations:  20,
			precision:   5e-4,
			expected:    0.3376232,
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.description, func(t *testing.T) {
			m := &MethodBase{
				f:  tc.function,
				g:  tc.qFunc,
				x0: tc.x0,
				k:  tc.iterations,
				e1: tc.precision,
				e2: tc.precision,
			}

			got := Round(MPF(m), 7)

			if got != tc.expected {
				t.Errorf("got %f but expected %f", got, tc.expected)
			}

		})
	}
}
