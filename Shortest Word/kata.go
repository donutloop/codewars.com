package Shortest_Word

import "math"

func FindShort(s string) int {

	min := math.Inf(1)
	wordLen := 0.0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if min == math.Inf(1) || min > wordLen {
				min = wordLen
			}
			wordLen = 0.0
			continue
		}
		wordLen = wordLen + 1
	}
	if wordLen > 0.0 &&  min > wordLen {
		return int(wordLen)
	}
	return int(min)
}