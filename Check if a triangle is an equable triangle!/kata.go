package kata

import "math"

func EquableTriangle(a, b, c int) bool {
	x := float64(a) + float64(b) + float64(c)
	s := x / 2
	y := math.Sqrt(s * (s - float64(a)) * (s - float64(b)) * (s - float64(c)))
	return x == y
}