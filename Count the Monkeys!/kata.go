package Count_the_Monkeys_

func MonkeyCount(n int) []int {
	numbers := make([]int, n)
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i + 1
	}
	return numbers
}