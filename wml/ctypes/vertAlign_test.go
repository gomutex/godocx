package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestVertAlignRun_MarshalXML(t *testing.T) {
	tests := []struct {
		name      string
		vertAlign VertAlignRun
		expected  string
	}{
		{
			name:      "Valid Value",
			vertAlign: VertAlignRun{Val: stypes.VerticalAlignRunBaseline},
			expected:  `<w:vertAlign w:val="baseline"></w:vertAlign>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:effect"}}

			err := tt.vertAlign.MarshalXML(encoder, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			encoder.Flush()
			if result.String() != tt.expected {
				t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expected, result.String())
			}
		})
	}
}

func TestVertAlignRun_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected VertAlignRun
	}{
		{
			name:     "Valid Value",
			inputXML: `<w:vertAlign w:val="baseline"></w:vertAlign>`,
			expected: VertAlignRun{Val: stypes.VerticalAlignRunBaseline},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var v VertAlignRun
			err := xml.Unmarshal([]byte(tt.inputXML), &v)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if v.Val != tt.expected.Val {
				t.Errorf("Expected val %s but got %s", tt.expected.Val, v.Val)
			}
		})
	}
}
