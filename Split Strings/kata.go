package Split_Strings

func Solution(str string) []string {
	if len(str) % 2 == 1 {
		str = str + "_"
	}

	pairs := make([]string, len(str)/2)
	var j int
	for i := 0; i < len(str); i = i + 2 {
		pairs[j] = string(str[i]) + string(str[i+1])
		j++
	}
	return pairs
}