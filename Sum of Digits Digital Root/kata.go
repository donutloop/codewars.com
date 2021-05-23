package Sum_of_Digits_Digital_Root

import "strconv"

func DigitalRoot(n int) int {
	s := strconv.Itoa(n)
	if len(s) == 1 {
		return n
	}

	var sum int
	for _, d := range s {
		sum = sum + int(d - 48)
	}

	return DigitalRoot(sum)
}