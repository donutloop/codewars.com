package moves_in_squared_strings__I_

const whitespace = '\n'

func VertMirror(s string) string {

	transformedS := make([]byte, 0)
	var lastWhitespace *int
	for i := 0; i < len(s); i++ {
		if whitespace == s[i] || len(s)-1 == i {
			if lastWhitespace == nil {
				lastWhitespace = new(int)
				for j := i-1; j >= 0; j-- {
					transformedS = append(transformedS, s[j])
				}
				*lastWhitespace = i
			} else {
				if len(s)-1 == i {
					transformedS = append(transformedS, whitespace)
				}

				for j := i; j > *lastWhitespace; j-- {
					transformedS = append(transformedS, s[j])
				}
				*lastWhitespace = i
			}
		}
	}
	return string(transformedS)
}

func HorMirror(s string) string {
	transformedS := make([]byte, 0)

	var lastWhitespace *int
	for i := len(s)-1; i >= 0; i-- {
		if whitespace == s[i] || i == 0 {
			if lastWhitespace == nil {
				lastWhitespace = new(int)
				transformedS = append(transformedS, s[i+1:]...)
				*lastWhitespace = i
			} else {
				if i == 0 {
					transformedS = append(transformedS, whitespace)
				}
				transformedS = append(transformedS, s[i:*lastWhitespace]...)
				*lastWhitespace = i
			}
		}
	}

	return string(transformedS)
}

type FParam func(string) string

func Oper(f FParam, x string) string {
	return f(x)
}