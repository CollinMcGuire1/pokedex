package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  trim  whitespace  ",
			expected: []string{"trim", "whitespace"},
		},
		{
			input:    "this is a longer string",
			expected: []string{"this", "is", "a", "longer", "string"},
		},
		{
			input:    "short",
			expected: []string{"short"},
		},
		{
			input:    "UPPER TO LOWER",
			expected: []string{"upper", "to", "lower"},
		},
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			t.Errorf("length of test (%v) != length of expected (%v)", len(actual), len(c.expected))
			continue // error and continue here
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			if word != expectedWord {
				t.Errorf("word (%v) does not match expected (%v)", word, expectedWord)
			}
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
}
