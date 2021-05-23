package Extra_Perfect_Numbers__Special_Numbers_Series__7_

func ExtraPerfect(n int) []int {
	numbers := make([]int, 0)
	for num := 1; num <= n; num++ {
		if IsSetBitsValid(num) {
			numbers = append(numbers, num)
		}
	}
	return numbers
}

func IsSetBitsValid(n int) bool {
	count := 0
	first := true
	for n > 0 {
		k := n & 1
		if first {
			if k == 1 {
				count += k
			} else {
				return false
			}
			first = false
		}
		n >>= 1
		if n == 0 && k == 1 {
			count++
		}
	}
	return count == 2
}
