package unpackstr

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	validTestCases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name: "with valid digits",
			input: "a4bc2d5e",
			output: "aaaabccddddde",
		},
		{
			name: "without digits",
			input: "abcd",
			output: "abcd",
		},
	}

	invalidTestCases := []struct {
		name   string
		input  string
	}{
		{
			name: "with one rune ",
			input: "a",
		},
		{
			name: "with first digit",
			input: "1abb",
		},
		{
			name: "with a sequence of digits",
			input: "abc44d",
		},
	}

	for _, tc := range validTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got, _ := Unpack(tc.input)
			if got != tc.output {
				t.Errorf("Expected %v, got %v", tc.output, got)
			}
		})
	}

	for _, tc := range invalidTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got, error := Unpack(tc.input)
			if error == nil {
				t.Errorf("Expected to get error, got %v", got)
			}
		})
	}
}
