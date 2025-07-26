package functions

import (
	"net/http"
	"os"
	"strings"
)

// Creating the Map
func ReadAsciiBanner(w http.ResponseWriter, r http.Request, filename string) map[rune][]string {
	file, err := os.ReadFile(filename)
	if err != nil {
		ErrorHandler(w, "Banner error:", http.StatusBadRequest)
		return nil
	}

	str := string(file)

	input := strings.Split(str, "\n\n")

	if len(input[0]) != 8 {
		input[0] = input[0][1:]
	}

	asciiMap := make(map[rune][]string)

	char := 32
	for _, value := range input {
		asciiMap[rune(char)] = strings.Split(value, "\n")
		char++
	}

	return asciiMap
}
