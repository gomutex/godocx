package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/internal"
)

func TestColumn_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Column
		expected string
	}{
		{
			name:     "With Width",
			input:    Column{Width: internal.ToPtr(uint64(500))},
			expected: `<w:gridCol w:w="500"></w:gridCol>`,
		},
		{
			name:     "Without Width",
			input:    Column{},
			expected: `<w:gridCol></w:gridCol>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:gridCol"}}

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

func TestColumn_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected Column
	}{
		{
			name:     "With Width",
			inputXML: `<w:gridCol w:w="750"></w:gridCol>`,
			expected: Column{Width: internal.ToPtr(uint64(750))},
		},
		{
			name:     "Without Width",
			inputXML: `<w:gridCol></w:gridCol>`,
			expected: Column{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Column

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if result.Width == nil && tt.expected.Width == nil {
				// Both are nil, which is fine
			} else if result.Width == nil || tt.expected.Width == nil || *result.Width != *tt.expected.Width {
				t.Errorf("Expected Width %v but got %v", tt.expected.Width, result.Width)
			}
		})
	}
}
