package kata

import "math"

func RacePodium(blocks int) [3]int {
	n := float64(int(blocks / 3))
	m := float64(blocks) / 3
	diff := math.Round((m - n) * 3)

	var a, b, c int
	if diff > 1 {
		a, b, c = int(n)+1, int(n)+int(diff), int(n)-1
	} else if diff == 1 {
		a, b, c = int(n)+1, int(n)+int(diff)+1, int(n)-2
	} else {
		a, b, c = int(n), int(n)+1, int(n)-1
	}

	if c == 0 {
		a = a - 1
		c = 1
	}

	return [3]int{a, b, c}
}
