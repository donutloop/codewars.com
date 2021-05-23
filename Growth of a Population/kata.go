package Growth_of_a_Population

func NbYear(p0 int, percent float64, aug int, p int) int {
	var c int
	sum := float64(p0)
	for sum < float64(p) {
		sum = sum + (sum * (percent/100)) + float64(aug)
		c++
	}
	return c
}