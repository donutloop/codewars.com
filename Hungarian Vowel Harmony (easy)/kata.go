package Hungarian_Vowel_Harmony__easy_

// Dative returns the dative case for a given Hungarian word.
func Dative(word string) (string) {
	w := []rune(word)
outerLoop:
	for i := len(w)-1; i >= 0; i-- {
		switch rune(w[i]) {
		case 'e', 'é','i', 'í', 'ö', 'ő', 'ü','ű':
			word += "nek"
			break outerLoop
		case 'a', 'á', 'o', 'ó', 'u', 'ú':
			word += "nak"
			break outerLoop

		}
	}
	return word
}