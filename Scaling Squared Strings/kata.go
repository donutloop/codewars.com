package Scaling_Squared_Strings

func Scale(s string, k, n int ) string {
	newS := ""
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] != 10 {
			for j := k; j > 0; j-- {
				word += string(s[i])
			}
		}
		if (s[i] == 10) || len(s)-1 == i {

			nWord := ""
			for j := n; j > 0; j-- {
				nWord += word + string(byte(10))
			}
			if len(s)-1 == i {
				nWord = nWord[:len(nWord)-1]
			}
			newS += nWord
			word = ""
		}
	}
	return newS
}
