package Exclusive__or___xor__Logical_Operator

func Xor(a, b bool) bool {
	if (a == true && b == false) || (a == false && b == true) {
		return true
	}
	return false
}