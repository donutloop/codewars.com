package Alphabetical_Addition

func AddLetters(letters []rune) rune {
	if len(letters) == 0 {
		return 122
	}
	var r rune
	for _, l := range letters {
		r += (l - 96)
	}
	r += 96
	for r > 122 {
		r = 96 + (r - 122)
	}
	return r
}
