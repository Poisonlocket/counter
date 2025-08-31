package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	data, _ := os.ReadFile("words.txt")
	_ = data
	wordCount := 0
	const spaceChar = 32

	for _, x := range data {
		if x == ' ' {
			wordCount++
		}

	}
	wordCount++
	fmt.Println(wordCount)
}
