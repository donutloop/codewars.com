package kata

func Spacify(s string) string {
	if len(s) == 0 {
		return s
	}

	sb := make([]byte, (len(s)*2)-1)
	sb[0] = s[0]
	j := 1
	for i := 1; i < len(s); i++ {
		sb[j] = ' '
		j = j + 1
		sb[j] = s[i]
		j = j + 1
	}
	return string(sb)
}
