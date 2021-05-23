package Count_the_Digit

func NbDig(n int, d int) int {
	var count int
	for k := 0; k <= n; k++ {
		count += findDigit(k * k, d)
	}
	return count
}

func findDigit(k, d int) int {
	if k == d {
		return 1
	}
	var count int
	for k != 0 {
		n := k % 10
		if d == n {
			count++
		}
		k = k / 10
	}
	return count
}