package Convert_string_to_camel_case

func ToCamelCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	camelCaseS := make([]byte, 0)
	var lastChar byte
	for i := 0; i < len(s); i++ {
		if lastChar  == '-' || lastChar  == '_' {
			lastChar = s[i]
			if s[i] >= 97 && s[i] <= 122 {
				camelCaseS = append(camelCaseS, s[i] - 32)
			} else {
				camelCaseS = append(camelCaseS, s[i])
			}
			continue
		}
		if s[i] == '-' || s[i]  == '_' {
			lastChar = s[i]
			continue
		}
		camelCaseS = append(camelCaseS, s[i])
		lastChar = s[i]
	}

	return string(camelCaseS)
}
