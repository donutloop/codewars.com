package Is_the_string_uppercase_

type MyString string

func (s MyString) IsUpperCase() bool {
	for _, char := range s {
		if char >= 97 && char <= 122 {
			return false
		}
	}
	return true
}