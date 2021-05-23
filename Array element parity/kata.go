package Array_element_parity

import "sort"

func Solve(arr []int) int {

	sort.Slice(arr, func(i, j int) bool {
		a := arr[i]
		b := arr[j]
		if a < 0 {
			a = a * -1
		}
		if b < 0 {
			b = b * -1
		}
		return a < b
	})

	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if i % 2 == 1 {
			if sum > 0 || sum < 0 {
				return arr[i-1]
			}
		}
	}

	return sum
}