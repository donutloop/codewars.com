package Counting_sheep___

func CountSheeps(numbers []bool) int {
	var c int
	for _, b := range numbers {
		if b {
			c++
		}
	}
	return c
}
