package Are_the_numbers_in_order_

func InAscOrder(numbers []int) bool {
	if len(numbers) == 0 {
		return false
	}
	pre := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < pre {
			return false
		}
		pre = numbers[i]
	}

	return true
}