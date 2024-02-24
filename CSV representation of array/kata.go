package kata

import (
	"strconv"
)

func ToCsvText(array [][]int) string {
	if len(array) == 0 {
		return ""
	}
	var csv string
	for _, nums := range array {
		var row string
		for _, num := range nums {
			col := strconv.Itoa(num)
			row += col + ","
		}
		csv += row[:len(row)-1] + "\n"
	}
	return csv[:len(csv)-1]
}
