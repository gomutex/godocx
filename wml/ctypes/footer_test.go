package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestFooterReference_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    FooterReference
		expected string
	}{
		{
			name: "Marshal with ID and Type",
			input: FooterReference{
				ID:   "rId1",
				Type: stypes.HdrFtrFirst,
			},
			expected: `<w:footerReference w:type="first" r:id="rId1"></w:footerReference>`,
		},
		{
			name: "Marshal with Type only",
			input: FooterReference{
				Type: stypes.HdrFtrEven,
			},
			expected: `<w:footerReference w:type="even"></w:footerReference>`,
		},
		{
			name: "Marshal with ID only",
			input: FooterReference{
				ID: "rId2",
			},
			expected: `<w:footerReference r:id="rId2"></w:footerReference>`,
		},
		{
			name:     "Marshal with neither ID nor Type",
			input:    FooterReference{},
			expected: `<w:footerReference></w:footerReference>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:titlePg"}}

			err := tt.input.MarshalXML(encoder, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			if result.String() != tt.expected {
				t.Errorf("Expected XML:\n%s\n\nGot:\n%s", tt.expected, result.String())
			}
		})
	}
}

func TestFooterReference_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected FooterReference
	}{
		{
			name:     "Unmarshal with ID and Type",
			inputXML: `<w:footerReference w:type="first" r:id="rId1"></w:footerReference>`,
			expected: FooterReference{
				ID:   "rId1",
				Type: stypes.HdrFtrFirst,
			},
		},
		{
			name:     "Unmarshal with Type only",
			inputXML: `<w:footerReference w:type="even"></w:footerReference>`,
			expected: FooterReference{
				Type: stypes.HdrFtrEven,
			},
		},
		{
			name:     "Unmarshal with ID only",
			inputXML: `<w:footerReference r:id="rId2"></w:footerReference>`,
			expected: FooterReference{
				ID: "rId2",
			},
		},
		{
			name:     "Unmarshal with neither ID nor Type",
			inputXML: `<w:footerReference></w:footerReference>`,
			expected: FooterReference{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result FooterReference

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error during unmarshaling: %v", err)
			}

			if result.ID != tt.expected.ID || result.Type != tt.expected.Type {
				t.Errorf("Expected %+v but got %+v", tt.expected, result)
			}
		})
	}
}
