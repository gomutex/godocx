package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestEmpty_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Empty
		expected string
	}{
		{
			name:     "Empty element",
			input:    Empty{},
			expected: `<w:tab></w:tab>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:tab"}}

			err := tt.input.MarshalXML(encoder, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			encoder.Flush()

			if result.String() != tt.expected {
				t.Errorf("Expected XML:\n%s\nGot:\n%s", tt.expected, result.String())
			}
		})
	}
}

func TestEmpty_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
	}{
		{
			name:     "Empty element",
			inputXML: `<w:tab></w:tab>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Empty

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}
		})
	}
}
