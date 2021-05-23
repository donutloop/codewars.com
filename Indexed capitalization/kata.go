package Indexed_capitalization

func Capitalize(st string, arr []int) string {
	b := []byte(st)

	for _, i := range arr {

		if i < 0 {
			continue
		}

		if i >= len(b) {
			continue
		}

		b[i] = b[i] - 32
	}

	return string(b)
}
