package iteration

// this function takes a string and repeats it 5 times
func Repeat(s string) string {
	var result string
	for i := 0; i < 5; i++ {
		result += s
	}
	return result
}
