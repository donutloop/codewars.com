package Reversed_Strings

import (
	"unicode/utf8"
)

func Solution(word string) string {
	v := make([]rune, len(word))
	w := []byte(word)
	var j int
	for len(w) > 0 {
		char, size := utf8.DecodeLastRune(w)
		w = w[:len(w)-size]
		v[j] = char
		j++
	}

	return string(v)
}