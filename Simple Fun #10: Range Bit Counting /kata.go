package Simple_Fun__10__Range_Bit_Counting_

import "strconv"

func RangeBitCount(a,b int) int {
	var count int
	for ; a <= b; a++ {
		count += countSetBits(strconv.FormatInt(int64(a), 2))
	}
	return count
}

func countSetBits(num string) int {
	var count int
	for _, digit := range num {
		if digit == '1' {
			count++
		}
	}

	return count
}