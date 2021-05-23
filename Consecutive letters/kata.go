package Consecutive_letters

import "sort"

func Solve(s string) bool {
	if len(s) <= 1 {
		return true
	}
	var ss []int
	for i := 0; i < len(s); i++ {
		ss = append(ss, int(s[i]))
	}
	sort.Ints(ss)

	lastChar := ss[0]
	for i := 1; i < len(ss); i++ {
		if lastChar == ss[i] {
			return false
		}
		if lastChar+1 != ss[i] {
			return false
		}
		lastChar = ss[i]
	}
	return true
}