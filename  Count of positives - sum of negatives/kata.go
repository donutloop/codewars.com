package _Count_of_positives___sum_of_negatives

func CountPositivesSumNegatives(numbers []int) []int {
	res := make([]int, 2)
	for _, n := range numbers {
		if n > 0 {
			res[0]++
		} else {
			res[1] += n
		}
	}
	return res
}
