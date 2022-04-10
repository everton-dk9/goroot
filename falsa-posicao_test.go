package goroot

import (
	"testing"
)

func TestMetodoFalsaPosicao(t *testing.T) {
	testsCases := []struct {
		description string
		function    fn
		a, b        float64
		precision   float64
		iterations  int
		want        float64
	}{
		{
			description: "False-Position->x^3-9x+3",
			function:    func(x float64) float64 { return ((x * x * x) - (9 * x) + 3) },
			a:           0,
			b:           1,
			precision:   1e-3,
			iterations:  16,
			want:        0.337635046,
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

			got := Round(FalsaPosicao(m), 9)

			if got != tc.want {
				t.Errorf("got %f but expected %f", got, tc.want)
			}
		})
	}

}
