package kata

func SetAlarm(employed, vacation bool) bool {
	return employed && !vacation
}
