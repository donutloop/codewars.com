package Character_with_longest_consecutive_repetition

type Result struct {
	C rune
	L int
}

func LongestRepetition(text string) Result {
	if text == "" {
		return Result{}
	}
	b := []rune(text)
	c := 1
	char := b[0]
	maxC := 1
	maxChar := char
	for i := 1; i < len(b); i++ {
		if char == b[i] {
			c++
			if c > maxC {
				maxC = c
				maxChar = b[i]
			}
		} else {
			char = b[i]
			c = 1
		}
	}

	return Result{
		C: maxChar,
		L: maxC,
	}
}