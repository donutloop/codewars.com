package DNA_to_RNA_Conversion

func DNAtoRNA(dna string) string {
	// your code here
	dnaCopy := []byte(dna)
	for i, char := range dnaCopy {
		if char == 84 {
			dnaCopy[i] = 85
		}
	}
	return string(dnaCopy)
}