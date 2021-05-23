package How_Green_Is_My_Valley_

import "sort"

func MakeValley(arr []int) []int {
	a := make([]int, len(arr))
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	j := len(arr) - 1
	var i int
	for k, v := range arr {
		if (k % 2) == 0 {
			a[i] = v
			i++
		} else {
			a[j] = v
			j--
		}
	}
	return a
}
