package elemtypes

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestSingleUint64Val_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    SingleUint64Val
		expected string
	}{
		{
			name:     "With value",
			input:    SingleUint64Val{Val: 10},
			expected: `<w:kern w:val="10"></w:kern>`,
		},
		{
			name:     "Empty value",
			input:    SingleUint64Val{Val: 18446744073709551615},
			expected: `<w:kern w:val="18446744073709551615"></w:kern>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:kern"}}

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

func TestSingleUint64Val_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected SingleUint64Val
	}{
		{
			name:     "With value",
			inputXML: `<w:kern w:val="00122"></w:kern>`,
			expected: SingleUint64Val{Val: 122},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result SingleUint64Val

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
