package poll

import (
	"math"
)

func getVariance() float64 {
	p := a.p / 100
	sSize := float64(a.sSize)
	pSize := float64(a.pSize)

	lhs := p * (1 - p) / sSize
	rhs := (pSize - sSize) / (pSize - 1)
	res := lhs * rhs
	return res
}

func clamp(value float64) float64 {
	min := 0.0
	max := 100.0

	if value > max {
		return max
	}
	if value < min {
		return min
	}
	return value
}

func getConfidentInterval(i float64, variance float64) ConfidenceInterval {
	interval := (2 * i * math.Sqrt(variance)) / 2 * 100

	return ConfidenceInterval{
		a: clamp(a.p - interval),
		b: clamp(a.p + interval),
	}
}