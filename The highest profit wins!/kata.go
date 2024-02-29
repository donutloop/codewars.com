package kata

import "math"

func MinMax(arr []int) [2]int {

	var max, min int = math.MinInt, math.MaxInt
	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return [2]int{min, max}
}
