package sections

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestSectionType_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    SectionType
		expected string
	}{
		{
			name:     "Valid SectionType",
			input:    SectionType{Val: stypes.SectionMark("nextPage")},
			expected: `<w:type w:val="nextPage"></w:type>`,
		},
		{
			name:     "Empty SectionType",
			input:    SectionType{},
			expected: `<w:type></w:type>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:type"}}

			err := tt.input.MarshalXML(encoder, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			encoder.Flush()

			if result.String() != tt.expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", tt.expected, result.String())
			}
		})
	}
}

func TestSectionType_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected SectionType
	}{
		{
			name:     "Valid SectionType",
			inputXML: `<w:type w:val="nextPage"></w:type>`,
			expected: SectionType{Val: stypes.SectionMark("nextPage")},
		},
		{
			name:     "Empty SectionType",
			inputXML: `<w:type></w:type>`,
			expected: SectionType{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result SectionType

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error during unmarshaling: %v", err)
			}

			if result.Val != tt.expected.Val {
				t.Errorf("Expected Val %s but got %s", tt.expected.Val, result.Val)
			}
		})
	}
}
