package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {

    cases := []struct {
	input    string
	expected []string
}{
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			err := "mismatched length"
			t.Errorf("mismatched length")
			t.Fatalf("fatal error: %v", err)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				err := "mismatched words"
				t.Errorf("mismatched words")
				t.Fatalf("fatal error: %v", err)
			}
		}
	}
}

