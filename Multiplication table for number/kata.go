package Multiplication_table_for_number

import "strconv"

func MultiTable(number int) string {

	var s string
	for i := 1; i <= 10; i++ {
		s += strconv.Itoa(i) + " * " + strconv.Itoa(number) + " = " + strconv.Itoa(i * number) + "\n"
	}
	return s[:len(s)-1]
}
