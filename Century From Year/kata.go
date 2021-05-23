package Century_From_Year

import (
	"math"
)

func century(year int) int {
	base := math.Floor(float64(year)/100) * 100
	diff := float64(year) - base
	if diff >= 1 {
		return int((base / 100) + 1)
	}
	return int(base /  100)
}