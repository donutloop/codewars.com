package Sum_of_Cubes

func SumCubes(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i * i * i
	}
	return
}
