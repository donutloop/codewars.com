package Consonant_value

func Solve(str string) int {
	maxSum := 0
	sum := 0
	for _, ch := range str {
		if ch=='a' || ch=='e' || ch=='i' || ch=='o' || ch=='u' {
			if sum > maxSum {
				maxSum = sum
			}
			sum = 0
			continue
		}
		sum = sum + (int(ch) - 96)
	}
	if sum > 0 {
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}