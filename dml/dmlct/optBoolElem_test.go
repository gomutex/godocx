package dmlct

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/dml/dmlst"
)

func TestOptBoolElem_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    OptBoolElem
		expected string
	}{
		{
			name:     "Valid true",
			input:    OptBoolElem{Val: dmlst.NewOptBool(true)},
			expected: `<w:b w:val="true"></w:b>`,
		},
		{
			name:     "Valid false",
			input:    OptBoolElem{Val: dmlst.NewOptBool(false)},
			expected: `<w:b w:val="false"></w:b>`,
		},
		{
			name:     "Invalid",
			input:    OptBoolElem{Val: dmlst.OptBool{Valid: false}},
			expected: `<w:b></w:b>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:b"}}

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

func TestOptBoolElem_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected OptBoolElem
	}{
		{
			name:     "Valid true",
			inputXML: `<w:b w:val="true"></w:b>`,
			expected: OptBoolElem{Val: dmlst.NewOptBool(true)},
		},
		{
			name:     "Valid false",
			inputXML: `<w:b w:val="false"></w:b>`,
			expected: OptBoolElem{Val: dmlst.NewOptBool(false)},
		},
		{
			name:     "Invalid",
			inputXML: `<w:b></w:b>>`,
			expected: OptBoolElem{Val: dmlst.OptBool{Valid: false}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result OptBoolElem

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if result.Val.Valid != tt.expected.Val.Valid || result.Val.Bool != tt.expected.Val.Bool {
				t.Errorf("Expected Val %v but got %v", tt.expected.Val, result.Val)
			}
		})
	}
}
