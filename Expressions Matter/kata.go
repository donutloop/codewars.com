package Expressions_Matter


func ExpressionMatter(a int, b int, c int) int {

	var max int
	y := a * (b + c)
	max = y

	y = a * b * c
	if y > max {
		max = y
	}

	y = a + b * c
	if y > max {
		max = y
	}

	y = (a + b) * c
	if y > max {
		max = y
	}

	y = a + b + c
	if y > max {
		max = y
	}

	return max
}