package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	data, _ := os.ReadFile("words.txt")
	CountWords(data)
	const spaceChar = 32

}

func CountWords(data []byte) int {
	words := bytes.Fields(data)
	return len(words)
}
