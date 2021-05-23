package Fix_string_case_

func solve(str string) string {

	var upperCase int
	var lowerCase int
	for _, s := range str {

		if s >= 'a' && s <= 'z' {
			lowerCase++
		} else if s >= 'A' && s <= 'Z' {
			upperCase++
		}
	}

	var add int
	if lowerCase >= upperCase {
		add = 32
	} else {
		add = -32
	}

	b := []byte(str)
	for i := range b {
		if add == -32 && (b[i] >= 'a' && b[i] <= 'z') {
			b[i] = b[i] + byte(add)
		} else if add == 32 && (b[i] >= 'A' && b[i] <= 'Z') {
			b[i] = b[i] + byte(add)
		}
	}

	return string(b)
}
