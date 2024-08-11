package util

import "math"

func FloatToCents(amount float64) int {
	return int(math.Round(amount * 100))
}

func CentsToFloat64(cents int) float64 {
	return float64(cents) / 100.0
}
