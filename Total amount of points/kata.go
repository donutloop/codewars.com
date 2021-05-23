package Total_amount_of_points

func Points(games []string) int {
	var sum int
	for _, game := range games {
		if len(game) != 3 {
			panic("string format is bad")
		}
		x := game[0] - 48
		y := game[2] - 48

		if x > y {
			sum += 3
		} else if x < y {
			sum += 0
		} else if x == y {
			sum += 1
		}
	}
	return sum
}