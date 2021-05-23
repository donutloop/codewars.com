package Multiplication_table

func MultiplicationTable(size int) [][]int {
	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
		for j := 0; j < size; j++ {
			if i > 0 {
				matrix[i][j] = matrix[0][j] * (i+1)
			} else {
				matrix[i][j] = j+1
			}
		}
	}

	return matrix
}