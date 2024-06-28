package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestHeaderReference_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    HeaderReference
		expected string
	}{
		{
			name: "Marshal with ID and Type",
			input: HeaderReference{
				ID:   "rId1",
				Type: stypes.HdrFtrFirst,
			},
			expected: `<w:headerReference w:type="first" r:id="rId1"></w:headerReference>`,
		},
		{
			name: "Marshal with Type only",
			input: HeaderReference{
				Type: stypes.HdrFtrEven,
			},
			expected: `<w:headerReference w:type="even"></w:headerReference>`,
		},
		{
			name: "Marshal with ID only",
			input: HeaderReference{
				ID: "rId2",
			},
			expected: `<w:headerReference r:id="rId2"></w:headerReference>`,
		},
		{
			name:     "Marshal with neither ID nor Type",
			input:    HeaderReference{},
			expected: `<w:headerReference></w:headerReference>`,
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

func TestHeaderReference_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected HeaderReference
	}{
		{
			name:     "Unmarshal with ID and Type",
			inputXML: `<w:headerReference w:type="first" r:id="rId1"></w:headerReference>`,
			expected: HeaderReference{
				ID:   "rId1",
				Type: stypes.HdrFtrFirst,
			},
		},
		{
			name:     "Unmarshal with Type only",
			inputXML: `<w:headerReference w:type="even"></w:headerReference>`,
			expected: HeaderReference{
				Type: stypes.HdrFtrEven,
			},
		},
		{
			name:     "Unmarshal with ID only",
			inputXML: `<w:headerReference r:id="rId2"></w:headerReference>`,
			expected: HeaderReference{
				ID: "rId2",
			},
		},
		{
			name:     "Unmarshal with neither ID nor Type",
			inputXML: `<w:headerReference></w:headerReference>`,
			expected: HeaderReference{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result HeaderReference

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
