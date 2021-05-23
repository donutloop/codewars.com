package Find_the_unique_number

import "sort"

func FindUniq(arr []float32) float32 {
	if len(arr) < 2 {
		return -1
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	if arr[1] == arr[0] && arr[1] == arr[len(arr)-1] {
		return -1
	} else if arr[1] == arr[len(arr)-1] {
		return arr[0]
	}
	return arr[len(arr)-1]
}
