package _Sum_of_integers_in_string

func SumOfIntegersInString(strng string) int {

	var globalSum int
	for i := len(strng) - 1; i >= 0; i-- {
		if strng[i] >= 48 && strng[i] <= 58 {
			var sum int
			n := 1
			sum += n * int(strng[i]-48)
			i--
			for ; i >= 0; i-- {
				if strng[i] >= 48 && strng[i] <= 58 {
					n = n * 10
					sum += n * int(strng[i]-48)
				} else {
					break
				}
			}
			globalSum += sum
		}
	}

	return globalSum
}
