package strand

// ToRNA will convert DNA to RNA
func ToRNA(dna string) string {

	mapping := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	var output string
	for i := 0; i < len(dna); i++ {
		output += string(mapping[dna[i]])
	}

	return output

}
