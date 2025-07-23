package functions

import (
	"strings"
)

// Generate ASCII art from input string
func AsciiRepresentation(str string, asciiMap map[rune][]string) string {
	result := ""
	words := strings.Split(str, "\r\n")
	slice := [][]string{}

	for _, word := range words {
		if len(word) == 0 {
			result += "\n"
			continue
		}

		for _, char := range word {
			for key, value := range asciiMap {
				if char == key {
					slice = append(slice, value)
				}
			}
		}

		for i := 0; i < 8; i++ {
			for _, line := range slice {
				result += line[i]
			}
			result += "\n"
		}

		slice = nil
	}

	return result
}
