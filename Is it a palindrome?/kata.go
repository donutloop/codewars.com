package Is_it_a_palindrome_

func IsPalindrome(str string) bool {
	b := []byte(str)
	j := len(str)-1
	for i := 0; i < len(b)/2; i++ {
		if b[j] >= 60 && b[j] <= 90 {
			b[j] = b[j] + 32
		}
		if b[i] >= 60 && b[i] <= 90 {
			b[i] = b[i] + 32
		}
		if b[i] != b[j] {
			return false
		}
		j--
	}
	return true
}