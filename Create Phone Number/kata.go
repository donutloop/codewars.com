package Create_Phone_Number

import "fmt"

func CreatePhoneNumber(n [10]uint) string {
	if len(n) != 10 {
		return "" // should return a error instead of a emtpy string
	}

	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", n[0],n[1],n[2],n[3],n[4],n[5],n[6],n[7],n[8],n[9])
}