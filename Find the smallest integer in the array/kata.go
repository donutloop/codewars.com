package Find_the_smallest_integer_in_the_array

func SmallestIntegerFinder(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}

	n := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if n > numbers[i] {
			n = numbers[i]
		}
	}

	return n
}
