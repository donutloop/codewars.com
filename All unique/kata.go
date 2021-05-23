package All_unique

func HasUniqueChar(str string) (ok bool) {
	set := make(map[rune]int)
	for _, char := range str {
		set[char] += 1
		if set[char] > 1 {
			ok = false
			return
		}
	}
	ok = true
	return
}