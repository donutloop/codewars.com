package Find_The_Parity_Outlier

func FindOutlier(integers []int) int {
	lastIndexOdd := -1
	lastIndexEven := -1
	count := 0
	for i := 0; i < len(integers); i++ {
		if integers[i] % 2 == 0 {
			count++
			lastIndexEven = i
		} else {
			lastIndexOdd = i
		}
	}

	if len(integers) - count == 1 {
		return integers[lastIndexOdd]
	}
	return integers[lastIndexEven]
}