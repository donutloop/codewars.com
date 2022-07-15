package kata

func Transpose(matrix [][]int) [][]int {
	transposeMatrix := make([][]int, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		k := 0

		for j := 0; j < len(matrix[i]); j++ {
			if transposeMatrix[k] == nil {
				transposeMatrix[k] = make([]int, len(matrix))
			}

			transposeMatrix[k][i] = matrix[i][j]
			k++
		}
	}

	return transposeMatrix
}
