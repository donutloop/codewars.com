package Vowel_Count

func GetCount(str string) (count int) {

	for _, char := range str {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			count +=1
		}
	}

	return count
}
