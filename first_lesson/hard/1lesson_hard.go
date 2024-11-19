package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// main reads a text from stdin and prints out the count of each word
// with all punctuation removed and in lower case.
func main() {
	wordCount := make(map[string]int)
	re := regexp.MustCompile(`[[:punct:]]`)

	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	for i := 0; i < len(strings.Fields(text)); i++ {
		word := re.ReplaceAllString(text, "")

		if strings.Fields(word)[i] != "" {
			wordCount[strings.ToLower(strings.Fields(word)[i])]++
		}
	}

	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}