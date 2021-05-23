package Abbreviate_a_Two_Word_Name

func AbbrevName(name string) string{
	if len(name) == 0 {
		return ""
	}

	var abbrevName []byte
	abbrevName = append(abbrevName, convertChar(name[0]))
	for i := 1; i < len(name); i++ {
		if name[i] == ' ' {
			if i+1 < len(name) {
				i++
				abbrevName = append(abbrevName, '.' ,convertChar(name[i]))
			}
		}
	}
	return string(abbrevName)
}

func convertChar(char byte) byte {
	if char >= 97 {
		char = char - 32
	}
	return char
}