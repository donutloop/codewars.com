package Alphabet_symmetry

func solve(slice []string) []int {
	counters := make([]int, len(slice))
	for j, charList := range slice {
		for i, char := range charList {
			if char >= 97 && 122 >= char {
				char = char - 97
			} else if char >= 65 && 90 >= char {
				char = char - 65
			}
			if i == int(char) {
				counters[j] += 1
			}
		}
	}
	return counters
}