package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	data, _ := os.ReadFile("words.txt")
	fmt.Println("data: ", string(data))
}
