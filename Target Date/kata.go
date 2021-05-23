package Target_Date

import "time"
import "fmt"

const (
	layoutISO = "2006-01-02"
)

func DateNbDays(a0 int, a int, p int) string {

	var days int
	k := float64(p) / float64(36000)
	a2 := float64(a0)
	for {
		if a2 > float64(a) {
			break
		}
		a2 = k*a2 + a2
		days++
	}

	now := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	then := now.AddDate(0, 0, days)
	return fmt.Sprintf("%d-%02d-%02d", then.Year(), then.Month(), then.Day())
}
