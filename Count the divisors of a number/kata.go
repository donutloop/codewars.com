package Count_the_divisors_of_a_number

func Divisors(n int)int{
	if n == 1 {
		return 1
	}
	var count int
	k := 1
	for n >= k {
		reminder := n % k
		if reminder == 0 {
			count++
		}
		k++
	}
	return count
}