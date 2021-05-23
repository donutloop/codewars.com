package Factorial

func Factorial(n int) int {
	sum := 1
	for i := 1; i <= n; i++ {
		sum *= i
	}
	return sum
}
