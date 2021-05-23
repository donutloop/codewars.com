package IP_Validation

func Is_valid_ip(ip string) bool {
	if ip == "" {
		return false
	}

	digit := ""
	dots := 0
	for n := 0; n < len(ip); n++ {
		if ip[n] == '.' {
			ok := check(digit)
			if !ok {
				return false
			}
			dots++
			digit = ""
		} else {
			digit = digit + string(ip[n])
		}
	}

	if dots != 3 {
		return false
	}
	if digit != "" {
		ok := check(digit)
		if !ok {
			return false
		}
	}

	return true
}

func check(digits string) bool {
	if len(digits) == 0 {
		return false
	}

	if digits[0] == '0' && len(digits) > 1 {
		return false
	}
	if len(digits) > 3 {
		return false
	}

	if len(digits) == 3 {

		if digits[0] == 49 {

			if digits[1] < 48 || digits[1] > 57 {
				return false
			}

			if digits[2] < 48 || digits[2]  > 57  {
				return false
			}

			return true
		} else {

			if digits[0] < 48 || digits[0] > 50 {
				return false
			}

			if digits[1] < 48 || digits[1] > 53 {
				return false
			}

			if digits[1] != 53 {

				if digits[2] < 48 || digits[2]  > 57 {
					return false
				}


			} else {
				if digits[2] < 48 || digits[2]  > 53 {
					return false
				}
			}
			return true
		}
	} else if len(digits) == 2 {
		if digits[0] < 48 || digits[0] > 57 {
			return false
		}

		if digits[1] < 48 || digits[1] > 57 {
			return false
		}
	} else if len(digits) == 1 {

		if digits[0] < 48 || digits[0] > 57 {

			return false
		}
	}

	return true
}