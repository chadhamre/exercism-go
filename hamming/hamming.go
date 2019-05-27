package hamming

import "errors"

// Distance will return the hamming difference b/w two strings
func Distance(a, b string) (int, error) {

	// check if strings match
	if len(a) != len(b) {
		return -1, errors.New("string length does not match")
	}

	// check if each char matches
	var distance int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil

}
