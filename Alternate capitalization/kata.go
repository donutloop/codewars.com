package Alternate_capitalization

func Capitalize(st string) []string {
	a := []byte(st)
	b := []byte(st)


	for i := 0; i < len(st); i++ {
		if i % 2 == 1 {
			// odd
			b[i] -= 32
		} else {
			// even
			a[i] -= 32
		}
	}

	return []string{string(a), string(b)}
}
