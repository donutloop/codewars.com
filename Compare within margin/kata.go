package kata

import "math"

func CloseCompare(a, b, margin float64) int {

	diff := math.Abs(a - b)
	if diff <= margin {
		return 0
	}

	if a > b {
		return 1
	}
	return -1
}
