package iteration

func Repeat(c string, n int) string {
	output := ""
	for i := 0; i < n; i++ {
		output += c
	}

	return output
}
