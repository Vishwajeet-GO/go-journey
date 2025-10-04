package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var sentence string
	fmt.Println("Enter any sentence:")

	reader := bufio.NewReader(os.Stdin)
	sentence, _ = reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence)

	words := strings.Fields(sentence)

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	fmt.Println("\nWord frequency:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
