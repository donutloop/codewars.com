package Strong_Number__Special_Numbers_Series__2_

func Strong(n int) string {

	var sum int
	oldN := n
	for n > 0 {
		digit := n % 10
		sum += factorial(digit)
		n /= 10
	}

	if sum == oldN {
		return "STRONG!!!!"
	}

	return "Not Strong !!"
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}

	if x == 1 {
		return 1
	}

	if x == 2 {
		return 2
	}

	return x * factorial(x-1)
}