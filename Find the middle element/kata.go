package kata

import "math"

func Gimme(array [3]int) int {

	var max, min = math.MinInt, math.MinInt
	for _, num := range array {
		if max == math.MinInt || max < num {
			max = num
		}
		if min == math.MinInt || min > num {
			min = num
		}
	}

	var middleIndex int
	for i, num := range array {
		if num == max || num == min {
			continue
		}
		middleIndex = i
	}

	return middleIndex
}
