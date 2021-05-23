package Simple_remove_duplicates

func Solve(arr []int) []int {

	nums := make(map[int]int, 0)
	for i := len(arr)-1; i >= 0; i-- {
		if _, ok := nums[arr[i]]; ok {
			continue
		}
		nums[arr[i]] = i
	}

	newArr := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		idx, ok := nums[arr[i]]
		if !ok {
			continue
		}
		if idx == i {
			newArr = append(newArr, arr[i])
		}
	}

	return newArr
}