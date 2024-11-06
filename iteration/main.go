package iteration

func Repeat(character string, expected int) string {
	var repeated string
	for i := 0; i < expected; i++ {
		repeated += character
	}
	return repeated
}