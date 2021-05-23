package Sum_of_odd_numbers

func RowSumOddNumbers(n int) int {
	if n == 1 {
		return n
	}

	number := 1
	for  i := 2; i < n+1; i++ {
		number = number + (i * 2)
	}

	sum := number
	for i := n; i > 1; i-- {
		number = number - 2
		sum = sum + number
	}

	return sum
}