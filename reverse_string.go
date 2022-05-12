package reverse_string

import "strings"

func ReverseString(input string) (output string) {
	var sb strings.Builder
	runes := []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {
		sb.WriteRune(runes[i])
	}
	output = sb.String()
	return
}
