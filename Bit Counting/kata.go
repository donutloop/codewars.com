package Bit_Counting

func CountBits(n uint) int {
	var count int
	v := n
	for v != 0 {
		r := v % 2
		v = v / 2
		if r == 1 {
			count++
		}
	}
	return count
}
