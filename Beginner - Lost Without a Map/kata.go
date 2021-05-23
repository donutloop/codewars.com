package Beginner___Lost_Without_a_Map

func Maps(x []int) []int {
	y := make([]int, len(x))
	for i := 0; i < len(x); i++ {
		y[i] = x[i] + x[i]
	}
	return y
}