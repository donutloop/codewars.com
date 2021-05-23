package Maximum_Length_Difference

import "math"

func MxDifLg(a1 []string, a2 []string) int {
	max1, min1, ok1 := findMaxMin(a1)
	max2, min2, ok2 := findMaxMin(a2)

	if !ok1 || !ok2 {
		return -1
	}

	diff1 := int(math.Abs(float64(max2))) - int(math.Abs(float64(min1)))
	diff2 := int(math.Abs(float64(max1))) - int(math.Abs(float64(min2)))

	var diff int
	if diff1 > diff2 {
		diff = diff1
	} else {
		diff = diff2
	}

	if diff > math.MaxInt32 {
		return -1
	}

	return diff
}

func findMaxMin(strings []string) (int, int, bool) {
	if len(strings) == 0 {
		return -1, -1, false
	}

	max := int(math.Inf(-1))
	min := int(math.Inf(1))
	for _, s := range strings {
		if max == int(math.Inf(-1)) || len(s) > max {
			max = len(s)
		}
		if min == int(math.Inf(1)) || len(s) < min {
			min = len(s)
		}
	}
	return max, min, true
}