package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestFontSize_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    FontSize
		expected string
	}{
		{
			name:     "With value",
			input:    *NewFontSize(24),
			expected: `<w:sz w:val="24"></w:sz>`,
		},
		{
			name:     "Without value",
			input:    FontSize{},
			expected: `<w:sz w:val="0"></w:sz>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:sz"}}

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

func TestFontSize_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected FontSize
	}{
		{
			name:     "With value",
			inputXML: `<w:sz w:val="24"></w:sz>`,
			expected: FontSize{Value: 24},
		},
		{
			name:     "Without value",
			inputXML: `<w:sz></w:sz>`,
			expected: FontSize{Value: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result FontSize

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if result.Value != tt.expected.Value {
				t.Errorf("Expected Value %d but got %d", tt.expected.Value, result.Value)
			}
		})
	}
}

func TestFontSizeCS_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    FontSizeCS
		expected string
	}{
		{
			name:     "With value",
			input:    *NewFontSizeCS(24),
			expected: `<w:szCs w:val="24"></w:szCs>`,
		},
		{
			name:     "Without value",
			input:    FontSizeCS{},
			expected: `<w:szCs w:val="0"></w:szCs>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:szCs"}}

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

func TestFontSizeCS_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected FontSizeCS
	}{
		{
			name:     "With value",
			inputXML: `<w:szCs w:val="24"></w:szCs>`,
			expected: FontSizeCS{Value: 24},
		},
		{
			name:     "Without value",
			inputXML: `<w:szCs></w:szCs>`,
			expected: FontSizeCS{Value: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result FontSizeCS

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if result.Value != tt.expected.Value {
				t.Errorf("Expected Value %d but got %d", tt.expected.Value, result.Value)
			}
		})
	}
}
