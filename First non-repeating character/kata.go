package kata

func FirstNonRepeating(str string) string {
	counter := make(map[rune]int)
	for _, codepoint := range str {
		counter[norm(codepoint)]++
	}
	for _, codepoint := range str {
		if counter[norm(codepoint)] == 1 {
			return string(codepoint)
		}
	}
	return ""
}

func norm(codepoint rune) rune {
	if codepoint >= 'A' && codepoint <= 'Z' {
		codepoint += 32
	}
	return codepoint
}
