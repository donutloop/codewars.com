package Breaking_chocolate_problem

func BreakChocolate(n, m int) int {
	x := n * m
	if x == 0 {
		return x
	}
	return x - 1
}