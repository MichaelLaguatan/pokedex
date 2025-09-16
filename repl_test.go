package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Cool runnings   ",
			expected: []string{"cool", "runnings"},
		},
		{
			input:    "abc def ghi jkl mno",
			expected: []string{"abc", "def", "ghi", "jkl", "mno"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("length of actual does not match expected, actual(%v): %v, expected(%v): %v", len(actual), actual, len(c.expected), c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("words in actual slice do not match expected, actual: %v, expected: %v", actual, c.expected)
			}
		}
	}
}
