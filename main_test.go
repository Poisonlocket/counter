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
	input = ""
	wants = 0

	res = countWords([]byte(input))
	if res != wants {
		t.Fail()
	}
	input = " "
	wants = 0
	res = countWords([]byte(input))
	if res != wants {
		t.Log("expected: ", wants, "got: ", res)
		t.Fail()
	}
}
