package Band_name_generator

func bandNameGenerator(word string) string {
	if len(word) == 0 {
		return ""
	}
	b := []byte(word)

	var secondCase bool
	if b[0] == b[len(b)-1] {
		secondCase = true
	}

	b[0] = toLower(b[0])

	var lastChar byte
	for i := 0; i < len(b); i++ {
		if lastChar == '-' {
			b[i] = toLower(b[i])
		}
		lastChar = b[i]
	}


	if secondCase {
		return string(b) + string(b[1:])
	}

	return "The " + string(b)
}

func toLower(c byte ) byte {
	if c >= 97 && c <= 122 {
		return (c - 32)
	}
	return c
}