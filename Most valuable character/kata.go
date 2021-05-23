package Most_valuable_character

import "math"

func Solve(s string) rune {

	type stats struct {
		maxIdx int
		minIdx int
	}

	chars := make(map[rune]stats)
	for i, char := range s {
		if char == ' ' {
			continue
		}

		v, ok := chars[char]
		if !ok {
			chars[char] = stats{maxIdx: i, minIdx: i}
		} else {
			v.maxIdx = i
			chars[char] = v
		}
	}


	var c rune
	maxDiff := int(math.Inf(-1))
	for char, stats := range chars {
		diff := stats.maxIdx - stats.minIdx

		if int(math.Inf(-1)) == maxDiff {
			c = char
			maxDiff = diff
			continue
		}

		if diff == maxDiff && char < c {
			c = char
			maxDiff = diff
		} else if diff > maxDiff {
			c = char
			maxDiff = diff
		}
	}

	return c
}