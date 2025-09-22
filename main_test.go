package main_test

import (
	counter "github.com/poisonlocket/counter"
	"testing"
)

func TestCountWords(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "five words",
			input: "one two three four five",
			wants: 5,
		},
		{
			name:  "empty input",
			input: "",
			wants: 0,
		},
		{
			name:  "single space",
			input: " ",
			wants: 0,
		},
		{
			name:  "newline",
			input: "one two three\nfour five",
			wants: 5,
		},
		{
			name:  "multiple spaces",
			input: "This is a sentence.  This is another",
			wants: 7,
		},
		{
			name:  "prefixed multiple spaces",
			input: "     Hello",
			wants: 1,
		},
		{
			name:  "suffixed multiple spaces",
			input: "Hello     ",
			wants: 1,
		},
		{
			name:  "Tab character in code",
			input: "Hello\tWord\n",
			wants: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := counter.CountWords([]byte(tc.input))
			if res != tc.wants {
				t.Log("expected: ", tc.wants, "got: ", res)
				t.Fail()
			}
		})
	}
}
