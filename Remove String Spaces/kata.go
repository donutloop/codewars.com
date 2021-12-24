package Remove_String_Spaces

func NoSpace(word string) string {

	b := make([]rune, 0, len(word)/3*2)
	for _, char := range word {
		if char != 32 {
			b = append(b, char)
		}
	}

	return string(b)
}
