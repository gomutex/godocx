package dml

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/dml/dmlst"
)

func TestMarshalNonVisualGraphicFrameProp(t *testing.T) {
	tests := []struct {
		prop        *NonVisualGraphicFrameProp
		expectedXML string
	}{
		{
			prop: &NonVisualGraphicFrameProp{
				GraphicFrameLocks: &GraphicFrameLocks{
					NoChangeAspect: dmlst.NewOptBool(true),
				},
			},
			expectedXML: `<wp:cNvGraphicFramePr><a:graphicFrameLocks xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" noChangeAspect="1"></a:graphicFrameLocks></wp:cNvGraphicFramePr>`,
		},
		{
			prop:        &NonVisualGraphicFrameProp{},
			expectedXML: `<wp:cNvGraphicFramePr></wp:cNvGraphicFramePr>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.expectedXML, func(t *testing.T) {
			generatedXML, err := xml.Marshal(tt.prop)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			if strings.TrimSpace(string(generatedXML)) != tt.expectedXML {
				t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expectedXML, generatedXML)
			}
		})
	}
}

func TestUnmarshalNonVisualGraphicFrameProp(t *testing.T) {
	tests := []struct {
		inputXML string
		expected NonVisualGraphicFrameProp
	}{
		{
			inputXML: `<wp:cNvGraphicFramePr><a:graphicFrameLocks noChangeAspect="1"></a:graphicFrameLocks></wp:cNvGraphicFramePr>`,
			expected: NonVisualGraphicFrameProp{
				GraphicFrameLocks: &GraphicFrameLocks{
					NoChangeAspect: dmlst.NewOptBool(true),
				},
			},
		},
		{
			inputXML: `<wp:cNvGraphicFramePr></wp:cNvGraphicFramePr>`,
			expected: NonVisualGraphicFrameProp{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			var prop NonVisualGraphicFrameProp

			err := xml.Unmarshal([]byte(tt.inputXML), &prop)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if (prop.GraphicFrameLocks == nil && tt.expected.GraphicFrameLocks != nil) ||
				(prop.GraphicFrameLocks != nil && tt.expected.GraphicFrameLocks == nil) {
				t.Errorf("Expected GraphicFrameLocks to be %v, but got %v", tt.expected.GraphicFrameLocks, prop.GraphicFrameLocks)
			} else if prop.GraphicFrameLocks != nil && tt.expected.GraphicFrameLocks != nil {
				if prop.GraphicFrameLocks.NoChangeAspect != tt.expected.GraphicFrameLocks.NoChangeAspect {
					t.Errorf("Expected GraphicFrameLocks.NoChangeAspect %v, but got %v", tt.expected.GraphicFrameLocks.NoChangeAspect, prop.GraphicFrameLocks.NoChangeAspect)
				}
			}
		})
	}
}
