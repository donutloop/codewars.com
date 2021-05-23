package Ordered_Count_of_Characters

import "sort"

type Tuple struct {
  Char  rune
  Count int
}

func OrderedCount(text string) []Tuple {

	type Stats struct {
		Count int
		First int
	}

	charCounts := make(map[rune]Stats)
	for i, char := range text {
		v, ok := charCounts[char]
		if ok {
			v.Count++
			charCounts[char] = v
		} else {
			charCounts[char] = Stats{Count: 1, First: i}
		}
	}

	tubles := make([]Tuple, 0)
	for char, stats := range charCounts {
		tubles = append(tubles, Tuple{Count: stats.Count, Char: char})
	}

	sort.Slice(tubles, func(i, j int) bool {
		return charCounts[tubles[i].Char].First < charCounts[tubles[j].Char].First
	})

	return tubles
}