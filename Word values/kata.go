package Word_values

func NameValue(words []string) []int {
	wordValues := make([]int, len(words))
	for i, word := range words {
		var sum int
		for _, char := range word {
			if char >= 96 && char <= 122 {
				sum += int(char) - 96
			}
		}
		w := i + 1
		value := w * sum
		wordValues[i] = value
	}
	return wordValues
}