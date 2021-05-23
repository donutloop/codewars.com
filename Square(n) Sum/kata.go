package Square_n__Sum

func SquareSum(numbers []int) (sum int) {
	for _, n := range numbers {
		sum += n * n
	}
	return
}