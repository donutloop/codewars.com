package MakeUpperCase

func MakeUpperCase(str string) string {

	b := []byte(str)
	for i := 0; i < len(b); i++ {
		if b[i] >= 97 && b[i] <= 122 {
			b[i] =  b[i] - 32
		}
	}

	return string(b)
}