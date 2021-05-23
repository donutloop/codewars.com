package String_prefix_and_suffix

func Solve(s string) int {

	if len(s) == 1 || len(s) == 0 {
		return 0
	}

	prefix := make(map[string]bool, 0)
	for i := 0; i <= len(s)/2; i++ {
		prefix[s[:i]] = true
	}

	var c int
	for i := len(s); i >= len(s)/2; i-- {
		ok := prefix[s[i:]]
		if ok {
			c = len(s[i:])
		}
	}

	return c
}
