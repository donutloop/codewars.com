package String_ends_with_

func solution(str, ending string) bool {
	if len(ending) > len(str) {
		return false
	}

	for i := 0; i < len(ending); i++ {
		a := str[len(str)-len(ending)+i]
		b := ending[i]
		if a == b {
			continue
		}
		return false
	}
	return true
}