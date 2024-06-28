package stypes

import (
	"testing"
)

func TestLongHexNumFromStr(t *testing.T) {
	tests := []struct {
		input    string
		expected LongHexNum
		err      bool
	}{
		{"", "", false},         // Empty string
		{"1a2b", "1A2B", false}, // Valid lowercase hex string
		{"1A2B", "1A2B", false}, // Valid uppercase hex string
		{"123", "", true},       // Invalid length (too short)
		{"12345", "", true},     // Invalid length (too long)
		{"1G2H", "", true},      // Invalid characters
		{"ZZZZ", "", true},      // Invalid characters
	}

	for _, tt := range tests {
		result, err := LongHexNumFromStr(tt.input)
		if (err != nil) != tt.err {
			t.Errorf("LongHexNumFromStr(%q) error = %v, expected error = %v", tt.input, err, tt.err)
		}
		if result != tt.expected {
			t.Errorf("LongHexNumFromStr(%q) = %v, expected %v", tt.input, result, tt.expected)
		}
	}
}
