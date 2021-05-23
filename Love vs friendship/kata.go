package Love_vs_friendship

func WordsToMarks(s string) int {
	var sum int
	for _, c := range s {
		sum += int(c) - 96
	}
	return sum
}