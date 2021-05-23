package String_repeat

func RepeatStr(repetitions int, value string) string {
	s := ""
	for repetitions > 0 {
		s += value
		repetitions--
	}
	return s
}