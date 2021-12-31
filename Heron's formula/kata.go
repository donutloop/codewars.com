package Heron_s_formula

import "math"

func Heron(a, b, c int) (area float32) {
	s := float64((a + b + c)) / 2
	s = s * (s - float64(a)) * (s - float64(b)) * (s - float64(c))
	s = math.Sqrt(s)
	return float32(s)
}
