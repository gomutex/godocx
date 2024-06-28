package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestPageNumbering_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    PageNumbering
		expected string
	}{
		{
			name:     "With format",
			input:    PageNumbering{Format: stypes.NumFmtDecimal},
			expected: `<w:pgNumType w:fmt="decimal"></w:pgNumType>`,
		},
		{
			name:     "Without format",
			input:    PageNumbering{},
			expected: `<w:pgNumType></w:pgNumType>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:pgNumType"}}

			err := tt.input.MarshalXML(encoder, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			encoder.Flush()

			if result.String() != tt.expected {
				t.Errorf("Expected XML:\n%s\n\nGot:\n%s", tt.expected, result.String())
			}
		})
	}
}

func TestPageNumbering_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected PageNumbering
	}{
		{
			name:     "With format",
			inputXML: `<w:pgNumType w:fmt="decimal"></w:pgNumType>`,
			expected: PageNumbering{Format: stypes.NumFmtDecimal},
		},
		{
			name:     "Without format",
			inputXML: `<w:pgNumType></w:pgNumType>`,
			expected: PageNumbering{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result PageNumbering

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error during unmarshaling: %v", err)
			}

			if result.Format != tt.expected.Format {
				t.Errorf("Expected Format %s but got %s", tt.expected.Format, result.Format)
			}
		})
	}
}
