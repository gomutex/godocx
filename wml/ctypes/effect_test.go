package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestEffect_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Effect
		expected string
	}{
		{
			name:     "With value",
			input:    Effect{Val: TextEffectPtr(stypes.TextEffectBlinkBackground)},
			expected: `<w:effect w:val="blinkBackground"></w:effect>`,
		},
		{
			name:     "Without value",
			input:    Effect{},
			expected: `<w:effect></w:effect>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:effect"}}

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

func TestEffect_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected Effect
	}{
		{
			name:     "With value",
			inputXML: `<w:effect w:val="blinkBackground"></w:effect>`,
			expected: Effect{Val: TextEffectPtr(stypes.TextEffectBlinkBackground)},
		},
		{
			name:     "Without value",
			inputXML: `<w:effect></w:effect>`,
			expected: Effect{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Effect

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if tt.expected.Val != nil {
				if result.Val == nil {
					t.Errorf("Expected Value %s but got nil", *tt.expected.Val)
				} else if *tt.expected.Val != *result.Val {
					t.Errorf("Expected Value %s but got %s", *tt.expected.Val, *result.Val)
				}
			} else {
				if result.Val != nil {
					t.Errorf("Expected nil but got %s", *result.Val)
				}
			}

		})
	}
}

func TextEffectPtr(value stypes.TextEffect) *stypes.TextEffect {
	return &value
}
