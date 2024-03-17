import "math"

func FindNextSquare(sq int64) int64 {
	value := math.Sqrt(float64(sq))
	correctValue := math.Ceil(value)
	if value != correctValue {
		return -1
	}
	value++
	return int64(value * value)
}