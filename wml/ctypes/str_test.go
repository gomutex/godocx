package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestCTString_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    CTString
		expected string
	}{
		{
			name:     "With value",
			input:    CTString{Val: "example"},
			expected: `<w:rStyle w:val="example"></w:rStyle>`,
		},
		{
			name:     "Empty value",
			input:    CTString{Val: ""},
			expected: `<w:rStyle w:val=""></w:rStyle>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:rStyle"}}

			err := tt.input.MarshalXML(encoder, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			// Finalize encoding
			encoder.Flush()

			if result.String() != tt.expected {
				t.Errorf("Expected XML:\n%s\nGot:\n%s", tt.expected, result.String())
			}
		})
	}
}

func TestCTString_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected CTString
	}{
		{
			name:     "With value",
			inputXML: `<w:rStyle w:val="example"></w:rStyle>`,
			expected: CTString{Val: "example"},
		},
		{
			name:     "Empty value",
			inputXML: `<w:rStyle w:val=""></w:rStyle>`,
			expected: CTString{Val: ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result CTString

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if result.Val != tt.expected.Val {
				t.Errorf("Expected Val %s but got %s", tt.expected.Val, result.Val)
			}
		})
	}
}
