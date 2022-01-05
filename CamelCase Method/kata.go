package CamelCase_Method

import (
	"unicode"
)

func CamelCase(s string) string {

	if s == "" {
		return ""
	}

	var b []rune
	next := true
	for _, r := range s {
		if r == 32 {
			next = true
			continue
		}

		if next {
			r = unicode.ToUpper(r)
			next = false
		}
		b = append(b, r)
	}

	return string(b)
}
