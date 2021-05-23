package Is_this_a_triangle_

func IsTriangle(a, b, c int) bool {
	ok := (a + b) > c
	if !ok {
		return false
	}
	ok = (a + c) > b
	if !ok {
		return false
	}
	ok = (b + c) > a
	if !ok {
		return false
	}
	return true
}