package Speed_Control

func Gps(s int, x []float64) int {
	var max float64
	for i := 1; i < len(x); i = i + 1 {
		delta := (x[i] - x[i-1])
		r := (3600 * float64(delta)) / float64(s)
		if r > max {
			max = r
		}
	}
	if max <= 1 {
		return 0
	}
	return int(max)
}
