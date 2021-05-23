package Simple_string_characters

func Solve(s string) []int {

	stats := make([]int, 4)
	for _, char := range s {
		if char >= 97 && char <= 122 {
			stats[1]++
		} else if char >= 65 && char <= 90 {
			stats[0]++
		} else if char >= 48 && char <= 57 {
			stats[2]++
		} else {
			stats[3]++
		}
	}

	return stats
}