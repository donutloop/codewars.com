package Mumbling

func Accum(s string) string {
	b := make([]rune, 0)
	for i, char := range s {
		if char >= 97 && char <= 122 {
			char = char - rune(32)
		}

		b = append(b, char)
		char = char + rune(32)

		for j := i; j > 0; j-- {
			b = append(b, char)
		}

		if len(s)-1 != i {
			b = append(b, '-')
		}
	}

	return string(b)
}