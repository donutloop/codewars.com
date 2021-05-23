package Powers_of_3

import "math"

func LargestPower(n uint64) int {
	if n == 1 {
		return -1
	}

	lastExponential := float64(0)
	for uint64(math.Pow(3, lastExponential)) < n {
		lastExponential++
	}
	return int(lastExponential) - 1
}