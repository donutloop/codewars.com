package Beginner___Reduce_but_Grow

func Grow(arr []int) (s int) {
	s = 1
	for _, n := range arr {
		s *= n
	}
	return
}