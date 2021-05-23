package Complementary_DNA

func DNAStrand(dna string) string {
	b := []byte(dna)
	for i := 0; i < len(b); i++ {
		switch b[i] {
		case 'A':
			b[i] = 'T'
		case 'T':
			b[i] = 'A'
		case 'C':
			b[i] = 'G'
		case 'G':
			b[i] = 'C'
		}
	}
	return string(b)
}