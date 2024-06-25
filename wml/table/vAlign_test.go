package table

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestVertAlignMarshalXML(t *testing.T) {
	tests := []struct {
		name      string
		vertAlign VertAlign
		expected  string
	}{
		{
			name: "Top alignment",
			vertAlign: VertAlign{
				Val: stypes.VerticalJc("top"),
			},
			expected: `<w:vAlign w:val="top"></w:vAlign>`,
		},
		{
			name: "Center alignment",
			vertAlign: VertAlign{
				Val: stypes.VerticalJc("center"),
			},
			expected: `<w:vAlign w:val="center"></w:vAlign>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sb strings.Builder
			e := xml.NewEncoder(&sb)
			start := xml.StartElement{Name: xml.Name{Local: "w:vAlign"}}
			if err := tt.vertAlign.MarshalXML(e, start); err != nil {
				t.Errorf("MarshalXML error: %v", err)
			}
			output := sb.String()
			if output != tt.expected {
				t.Errorf("MarshalXML() got = %v, want %v", output, tt.expected)
			}
		})
	}
}

func TestVertAlignUnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected VertAlign
	}{
		{
			name:  "Top alignment",
			input: `<w:vAlign w:val="top"></w:vAlign>`,
			expected: VertAlign{
				Val: stypes.VerticalJc("top"),
			},
		},
		{
			name:  "Center alignment",
			input: `<w:vAlign w:val="center"></w:vAlign>`,
			expected: VertAlign{
				Val: stypes.VerticalJc("center"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var v VertAlign
			err := xml.Unmarshal([]byte(tt.input), &v)
			if err != nil {
				t.Errorf("UnmarshalXML error: %v", err)
			}
			if v.Val != tt.expected.Val {
				t.Errorf("UnmarshalXML() got = %v, want %v", v.Val, tt.expected.Val)
			}
		})
	}
}
