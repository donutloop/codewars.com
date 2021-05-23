package Stop_gninnipS_My_sdroW

func SpinWords(str string) string {
	var newStr string
	var word string
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			newStr = newStr + reverse(word)  + " "
			word = ""
		} else {
			word = word + string(str[i])
		}
	}

	if word != "" {
		newStr = newStr + reverse(word)
	}

	return newStr
}

func reverse(word string) string {
	var newWord string
	if len(word) >= 5 {
		for j := len(word)-1; j >= 0; j-- {
			newWord = newWord + string(word[j])
		}
		return newWord
	}
	return word
}