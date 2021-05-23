package Beginner_Series__2_Clock

func Past(h, m, s int) int {
	sum := h * 60 * 60 * 1000
	sum += m * 60 * 1000
	sum += s * 1000
	return sum
}