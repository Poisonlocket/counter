package main

import (
	"fmt"
	"testing"
)

func TestCountWords(t *testing.T) {
	input := "one two three four five"
	wants := 5
	res := countWords([]byte(input))
	fmt.Println(res)
	if res != wants {
		t.Fail()
	}
}
