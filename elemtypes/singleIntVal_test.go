package elemtypes

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestSingleIntVal_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    SingleIntVal
		expected string
	}{
		{
			name:     "With value",
			input:    SingleIntVal{Val: 10},
			expected: `<w:spacing w:val="10"></w:spacing>`,
		},
		{
			name:     "Empty value",
			input:    SingleIntVal{Val: -1},
			expected: `<w:spacing w:val="-1"></w:spacing>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:spacing"}}

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

func TestSingleIntVal_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected SingleIntVal
	}{
		{
			name:     "With value",
			inputXML: `<w:spacing w:val="00122"></w:spacing>`,
			expected: SingleIntVal{Val: 122},
		},
		{
			name:     "Empty value",
			inputXML: `<w:spacing w:val="+3"></w:spacing>`,
			expected: SingleIntVal{Val: 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result SingleIntVal

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if result.Val != tt.expected.Val {
				t.Errorf("Expected Val %d but got %d", tt.expected.Val, result.Val)
			}
		})
	}
}
