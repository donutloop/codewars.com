package Grasshopper___Summation

func Summation(n int) int {
	var sum int
	for i := 0; i <= n; i++ {
		sum = sum + i
	}
	return sum
}