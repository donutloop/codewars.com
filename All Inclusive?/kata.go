package All_Inclusive_

func ContainAllRots(strng string, arr []string) bool {
	if len(strng) == 0 {
		return true
	}

	if len(strng) == 0 && len(arr) == 0 {
		return true
	}

	set := make(map[string]bool)
	b := []byte(strng)
	set[string(b)] = true
	for i := len(b); i > 0; i-- {
		tmp := b[len(b)-1]
		for j := len(b)-2; j >= 0; j-- {
			b[j+1] = b[j]
		}
		b[0] = tmp
		set[string(b)] = true
	}

	for _, s := range arr {
		if len(s) == len(b) && set[s] {
			delete(set, s)
		}
	}

	return len(set) == 0
}
