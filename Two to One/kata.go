package Two_to_One

import "sort"

func TwoToOne(s1 string, s2 string) string {

	s := []rune(s1 + s2)

	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	set := make(map[rune]bool)
	chars := make([]rune, 0)
	for _, char := range s {
		_, ok := set[char]
		if !ok {
			set[char] = true
			chars = append(chars, char)
		}
	}

	return string(chars)
}