package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func removePunctuation(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1
		}
		return r
	}, s)
}

func main() {

	var sentence string
	fmt.Println("Enter any sentence:")

	reader := bufio.NewReader(os.Stdin)
	sentence, _ = reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence)

	words := strings.Fields(sentence)

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[strings.ToLower(word)]++
	}

	fmt.Println("\nWord frequency:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}

	fmt.Println("\nTotal words:", len(words))
}
