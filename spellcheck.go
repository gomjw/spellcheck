package spellcheck

import (
	"bufio"
	"os"
	"strings"

)

var (
	wordlist []string
)

// CheckWord checks if the input is in a dataset of every english word.
func CheckWord(input string) bool {
	initModel()

	input = strings.ToLower(input)

	if contains(wordlist, input) {
		return true
	}

	return false
}

func initModel() {
	if len(wordlist) == 0 {
		lines, err := readFile("wordlist.txt")
		if err != nil {
			panic(err)
		}
		wordlist = lines
	}
}

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func contains(list []string, search string) bool {
	for _, word := range list {
		if word == search {
			return true
		}
	}
	return false
}
