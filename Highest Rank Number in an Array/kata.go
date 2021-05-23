package Highest_Rank_Number_in_an_Array

func HighestRank(nums []int) int {
	counter := make(map[int]int)
	var max int
	var val int
	for i := range nums {
		counter[nums[i]] = counter[nums[i]] + 1
		if (counter[nums[i]] >= max && nums[i] > val) || (counter[nums[i]] > max) {
			max = counter[nums[i]]
			val = nums[i]
		}
	}

	return val
}