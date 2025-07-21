package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestPict_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Pict
		expected string
	}{
		{
			name:     "Empty Pict",
			input:    Pict{},
			expected: `<w:pict></w:pict>`,
		},
		{
			name: "Pict with Shape",
			input: Pict{
				Shape: &Shape{
					Type:  "rectangle",
					Style: "width:100pt;height:50pt",
				},
			},
			expected: `<w:pict><v:shape type="rectangle" style="width:100pt;height:50pt"></v:shape></w:pict>`,
		},
		{
			name: "Pict with Shape and ImageData",
			input: Pict{
				Shape: &Shape{
					Type:  "rectangle",
					Style: "width:100pt;height:50pt",
					ImageData: &ImageData{
						RId:   "rId5",
						Title: "image1.png",
					},
				},
			},
			expected: `<w:pict><v:shape type="rectangle" style="width:100pt;height:50pt"><v:imagedata r:id="rId5" o:title="image1.png"></v:imagedata></v:shape></w:pict>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:pict"}}

			err := tt.input.MarshalXML(encoder, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			if err = encoder.Flush(); err != nil {
				t.Fatalf("Error flushing XML encoder: %v", err)
			}

			if result.String() != tt.expected {
				t.Errorf("Expected XML:\n%s\nGot:\n%s", tt.expected, result.String())
			}
		})
	}
}
