package iteration

//Repeat accepts character and n(no of repeats) and returns a string with character repeated n times
func Repeat(character string, n int) string {
	var repeated string

	for i := 0; i < n; i++ {
		repeated += character
	}

	return repeated
}
