package Highest_Scoring_Word


func High(s string) string {
	// 97 122
	var maxScore int
	var score int
	var word string
	var bestWord string
	for i := 0; i < len(s); i++ {

		if s[i] == ' ' {
			score = 0
			word = ""
		}

		if s[i] >= 97 && s[i] <= 122 {
			score = score + int(s[i]) - 96
			word = word + string(s[i])
		}

		if score > maxScore {
			maxScore = score
			bestWord = word
		}
	}

	return bestWord
}