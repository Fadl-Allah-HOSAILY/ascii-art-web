package functions

import (
	"strings"
)

func AsciiRepresentation(str string, asciiMap map[rune][]string) string {
	var result strings.Builder
	words := strings.Split(str, "\n")
	if words[0] == "" {
		words = words[1:]
	}
	slice := [][]string{}
	for _, word := range words {
		for _, char := range word {
			for key, value := range asciiMap {
				if char == key {
					slice = append(slice, value)
				}
			}
		}
		for i := 0; i < 8; i++ {
			for _, char := range slice {
				result.WriteString(char[i])
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}
