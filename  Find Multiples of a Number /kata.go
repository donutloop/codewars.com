package kata

func FindMultiples(integer, limit int) []int {
	var nums []int
	for num := integer; num < limit+1; num=num+integer {
		nums = append(nums, num)
	}
	return nums
}