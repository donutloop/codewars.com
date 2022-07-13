package kata

import "math"

func StringToNumber(str string) int {

	sign := 1
	var u int
	var i int
	if str[0] == '-' {
		sign = -1
		u = int(math.Pow(10, float64(len(str)-2)))
		i++
	} else {
		u = int(math.Pow(10, float64(len(str)-1)))
	}

	var sum int
	for ; i < len(str); i++ {
		d := int(str[i] - 48)
		sum += d * u
		u = u / 10
	}
	return sum * sign
}
