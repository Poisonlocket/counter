package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	data, _ := os.ReadFile("words.txt")
	countWords(data)
	const spaceChar = 32

}

func countWords(data []byte) int {
	if len(data) == 0 {
		return 0
	}
	wordCount := 0
	for _, x := range data {
		if x == ' ' {
			wordCount++
		}

	}
	if (len(data)) > 0 {
		wordCount++
	}

	return wordCount
}
