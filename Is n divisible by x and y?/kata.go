package Is_n_divisible_by_x_and_y_

func IsDivisible(n, x, y int) bool {
	return n % x == 0 && n % y == 0
}