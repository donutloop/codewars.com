package Previous_multiple_of_three

func PrevMultOfThree(n int) interface{} {
	for n > 0 {
		if n%3 == 0 {
			return n
		}
		n = n / 10
	}
	return nil
}
