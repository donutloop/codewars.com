package kata

func Between(a, b int) []int {
	res := make([]int, 0, b-a)
	for ; a <= b; a++ {
		res = append(res, a)
	}
	return res
}
