package Printer_Errors

import "strconv"

func PrinterError(s string) string {
	var count int
	for _, char := range s {
		if char > 'm' {
			count++
		}
	}

	return strconv.Itoa(count) + "/" + strconv.Itoa(len(s))
}
