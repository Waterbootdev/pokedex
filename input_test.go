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
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello world!",
			expected: []string{"hello", "world!"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "                   ",
			expected: []string{},
		},
		{
			input:    "hello                          world!",
			expected: []string{"hello", "world!"},
		},
	}
	for _, cs := range cases {
		t.Run(cs.input, func(t *testing.T) {
			actual := CleanInput(cs.input)
			if len(actual) != len(cs.expected) {
				t.Errorf("The length are not equal: %v vs %v", len(actual), len(cs.expected))
				return
			}
			for i := range actual {
				actualWord := actual[i]
				expectedWord := cs.expected[i]
				if actualWord != expectedWord {
					t.Errorf("%v does not equal %v", actualWord, expectedWord)
				}
			}
		})
	}
}
