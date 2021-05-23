package Numericals_of_a_String

import (
	"strconv"
)

func Numericals(s string) string {

	r := []rune(s)
	counter := make(map[rune]int, len(r)/2)
	ns := ""
	for i := 0; i < len(r); i++ {
		counter[r[i]] = counter[r[i]] + 1
		n := strconv.Itoa(counter[r[i]])
		ns = ns + n
	}
	counter = nil
	r = nil

	return ns
}