package To_square_root__or_not_to_square_root_

import "math"
func SquareOrSquareRoot(arr []int) []int{
	out := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		v := int(math.Sqrt(float64(arr[i])))
		if v * v == arr[i] {
			out[i] = v
		} else {
			out[i] = arr[i] * arr[i]
		}
	}
	return out
}