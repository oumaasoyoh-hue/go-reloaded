package main

import (
	"testing"
)

func TestProcessText(t *testing.T) {
	// Table-driven test containing all 4 official audit cases
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Audit Case 1: Multi-word tags & Punctuation spacing",
			input:    "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			expected: "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
		},
		{
			name:     "Audit Case 2: Number Conversions (hex and bin)",
			input:    "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			expected: "I have to pack 5 outfits. Packed 26 just to be sure",
		},
		{
			name:     "Audit Case 3: Standard Punctuation Spacing",
			input:    "Don not be sad ,because sad backwards is das . And das not good",
			expected: "Don not be sad, because sad backwards is das. And das not good",
		},
		{
			name:     "Audit Case 4: Capitalization, Quotes, and Vowels",
			input:    "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			expected: "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ProcessText(tt.input)
			if actual != tt.expected {
				t.Errorf("\n[FAILED] %s\nInput:    %q\nExpected: %q\nActual:   %q", tt.name, tt.input, tt.expected, actual)
			}
		})
	}
}
