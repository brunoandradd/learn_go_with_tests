package interation

// import "strings"

const RepeatCount = 4

func Repeat(letter string) string {
	// return strings.Repeat(letter, 4)
	var repeat string
	for i := 0; i < RepeatCount; i++ {
		repeat += letter
	}

	return repeat
}
