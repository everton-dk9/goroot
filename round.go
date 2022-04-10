package goroot

import "math"

func Round(x float64, places int) float64 {
	base := math.Pow10(places)
	return (math.Round(x*base) / base)
}
