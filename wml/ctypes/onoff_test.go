package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/internal"
	"github.com/gomutex/godocx/wml/stypes"
)

func TestOnOff_MarshalXML(t *testing.T) {

	tests := []struct {
		name     string
		input    OnOff
		expected string
	}{
		{
			name:     "With value",
			input:    OnOff{Val: internal.ToPtr(stypes.OnOffFalse)},
			expected: `<w:rStyle w:val="false"></w:rStyle>`,
		},
		{
			name:     "Empty value",
			input:    OnOff{Val: nil},
			expected: `<w:rStyle></w:rStyle>`,
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

func TestOnOff_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected OnOff
	}{
		{
			name:     "With value",
			inputXML: `<w:rStyle w:val="true"></w:rStyle>`,
			expected: OnOff{Val: internal.ToPtr(stypes.OnOffTrue)},
		},
		{
			name:     "Empty value",
			inputXML: `<w:rStyle></w:rStyle>`,
			expected: OnOff{Val: nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result OnOff

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if err = internal.ComparePtr("Val", tt.expected.Val, result.Val); err != nil {
				t.Error(err)
			}
		})
	}
}
