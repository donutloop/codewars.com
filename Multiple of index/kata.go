package Multiple_of_index

func multipleOfIndex (ints []int) []int {
	result := make([]int, 0, len(ints)/2)
	for i := 1; i < len(ints); i++ {
		if ints[i] % i == 0 {
			result = append(result, ints[i])
		}
	}
	return result
}