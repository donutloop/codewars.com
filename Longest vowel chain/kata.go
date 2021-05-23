package Longest_vowel_chain

func Solve(s string) int {
	count := 0
	max := 0
	for _, char := range s {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			count += 1
			if count > max {
				max = count
			}
		} else {
			count = 0
		}
	}
	return max
}