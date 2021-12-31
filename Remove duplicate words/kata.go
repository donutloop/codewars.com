package Remove_duplicate_words

import "strings"

func RemoveDuplicateWords(str string) string {

	words := make(map[string]bool)

	var word []byte
	var o []byte
	for i := 0; i < len(str); i++ {

		if str[i] == 32 || i == len(str)-1 {
			if i == len(str)-1 && str[i] != 32 {
				word = append(word, str[i])
			}
			ok := words[string(word)]
			if !ok {
				o = append(o, ' ')
				o = append(o, word...)
				words[string(word)] = true
			}
			word = nil
			continue
		}
		if str[i] >= 97 && str[i] <= 122 {
			word = append(word, str[i])
		} else if str[i] >= 65 && str[i] <= 90 {
			word = append(word, str[i])
		}
	}

	return strings.TrimSpace(string(o))
}
