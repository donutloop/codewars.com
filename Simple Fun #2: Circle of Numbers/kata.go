package Simple_Fun__2__Circle_of_Numbers

import "math"

func CircleOfNumbers(n int, firstNumber int) int {
	sum := (n / 2) + firstNumber
	if sum < n {
		return sum
	}
	sum = (n / 2) - firstNumber
	return int(math.Abs(float64(sum)))
}