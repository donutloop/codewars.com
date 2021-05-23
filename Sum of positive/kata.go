package Sum_of_positive

func PositiveSum(numbers []int) (sum int) {
	for _, n := range numbers {
		if n > 0 {
			sum += n
		}
	}
	return
}