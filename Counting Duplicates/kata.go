package Counting_Duplicates

func Duplicate_count(s1 string) int {

	duplicates := 0
	count := make(map[rune]int)
	for _, char := range s1 {

		if char >= 97 && char <= 122 {
			char = char - 32
		}

		count[char] = count[char] + 1
		if count[char] == 2 {
			duplicates = duplicates + 1
		}
	}

	return duplicates
}
