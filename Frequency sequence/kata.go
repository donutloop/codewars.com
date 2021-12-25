package Frequency_sequence

import "strconv"

func FreqSeq(str string, sep string) string {

	freqCounter := make(map[rune]int)
	for _, char := range str {
		freqCounter[char]++
	}

	b := make([]byte, 0, len(str))
	for _, char := range str {
		n := strconv.Itoa(freqCounter[char])
		b = append(b, []byte(n)...)
		b = append(b, sep...)
	}

	return string(b[:len(b)-1])
}
