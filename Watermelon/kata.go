package Watermelon

func Divide(weight int) bool {
	if weight == 2 {
		return false
	}

	return weight % 2 == 0
}