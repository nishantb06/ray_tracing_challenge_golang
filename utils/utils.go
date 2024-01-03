package utils

import "math"

const epsilon = 0.00001

func Equal(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}
