package kata

import (
	"strconv"
)

func WordPattern(word string) string {
	wordCounter := make(map[rune]int)

	var c int
	var encoded string
	for _, codepoint := range word {
		if codepoint >= 65 && codepoint <= 90 {
			codepoint += 32
		}

		preC, ok := wordCounter[codepoint]
		if !ok {
			wordCounter[codepoint] = c
			encoded += strconv.Itoa(c)
			encoded += "."
			c++
		} else {
			encoded += strconv.Itoa(preC)
			encoded += "."
		}
	}

	return encoded[:len(encoded)-1]
}
