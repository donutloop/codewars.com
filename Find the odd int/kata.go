package Find_the_odd_int

func FindOdd(seq []int) (odd int) {
	numbers := make(map[int]int)
	for _, s := range seq {
		numbers[s] = numbers[s] + 1
	}
	for n, c := range numbers {
		if c % 2 != 0 {
			odd = n
			break
		}
	}
	return
}