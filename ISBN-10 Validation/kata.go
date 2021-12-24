package ISBN_10_Validation

func ValidISBN10(isbn string) bool {

	if len(isbn) != 10 {
		return false
	}

	var sum int
	for i := 0; i < len(isbn); i++ {
		if isbn[i] >= 48 && isbn[i] <= 57 {
			n := (int(isbn[i]) - 48) * (i + 1)
			sum += n
		} else if (isbn[i] == 88 || isbn[i] == 120) && i > 8 {
			sum += 10 * (i + 1)
		} else {
			return false
		}
	}

	return sum%11 == 0
}
