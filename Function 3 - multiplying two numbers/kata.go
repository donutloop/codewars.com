package Function_3___multiplying_two_numbers

func Multiply(numbers ...int) (sum int) {
	sum = 1
	for _, n := range numbers {
		sum *= n
	}
	return
}