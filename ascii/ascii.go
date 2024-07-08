package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Files(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func PrintWord(input string, content []string) string {
	slice := make([]string, 9)

	for _, char := range input {
		for i := 0; i < 9; i++ {
			slice[i] += content[int(char-32)*9+i]
		}
	}
	return strings.Join(slice, "\n")
}

func Ascii(input string, filename string) string {
	if input == "" {
		return ""
	}

	if input == "\\n" || input == "\n" {
		return "\n"
	}

	content, err := Files(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}

	var res strings.Builder
	words := strings.Split(input, "\n")
	for _, word := range words {
		if word == "" {
			res.WriteString("\n")
		} else {
			if English(word) {
				res.WriteString(PrintWord(word, content) + "\n")
			} else {
				res.WriteString("Invalid input: not accepted\n")
			}
		}
	}
	return res.String()
}

func English(words string) bool {
	for _, word := range words {
		if word < 32 || word > 126 {
			return false
		}
	}
	return true
}