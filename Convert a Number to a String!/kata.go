package kata

import "math"

func NumberToString(n int) string {
	s := []byte{}
	if n < 0 {
		s = append(s, '-')
		n = n * -1
	}

	digitCount := int(math.Log10(float64(n)))

	for i := 0; i <= digitCount; i++ {
		digit := n/int(math.Pow(10, float64(digitCount-i)))%10 + 48
		s = append(s, byte(digit))
	}

	return string(s)
}
