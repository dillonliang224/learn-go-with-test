package integers

func Add(x, y int) int {
	return x + y
}

func Repeat(character string, num int) string {
	var repeated string
	for i := 0; i < num; i++ {
		repeated += character
	}

	return repeated
}
