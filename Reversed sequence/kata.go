package Reversed_sequence

func ReverseSeq(n int) []int {
	nums := make([]int, n)
	var j int
	for i := n; i > 0; i-- {
		nums[j] = i
		j++
	}
	return nums
}
