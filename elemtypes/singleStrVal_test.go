package elemtypes

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestSingleStrVal_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    SingleStrVal
		expected string
	}{
		{
			name:     "With value",
			input:    SingleStrVal{Val: "example"},
			expected: `<w:rStyle w:val="example"></w:rStyle>`,
		},
		{
			name:     "Empty value",
			input:    SingleStrVal{Val: ""},
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

func TestSingleStrVal_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected SingleStrVal
	}{
		{
			name:     "With value",
			inputXML: `<w:rStyle w:val="example"></w:rStyle>`,
			expected: SingleStrVal{Val: "example"},
		},
		{
			name:     "Empty value",
			inputXML: `<w:rStyle w:val=""></w:rStyle>`,
			expected: SingleStrVal{Val: ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result SingleStrVal

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
