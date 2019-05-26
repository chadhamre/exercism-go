package twofer

// ShareWith will build and return a sentence
func ShareWith(name string) string {

	if len(name) == 0 {
		name = "you"
	}
	return "One for " + name + ", one for me."

}
