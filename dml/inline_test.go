package dml

import (
	"encoding/xml"
	"testing"

	"github.com/gomutex/godocx/common/constants"
	"github.com/gomutex/godocx/dml/dmlct"
	"github.com/gomutex/godocx/dml/dmlst"
)

func TestMarshalInline(t *testing.T) {

	tests := []struct {
		inline      *Inline
		expectedXML string
		xmlName     string
	}{
		{
			inline: &Inline{
				DistT: 2,
				DistB: 3,
				DistL: 4,
				DistR: 5,
				Extent: dmlct.PSize2D{
					Width:  100,
					Height: 200,
				},
				DocProp: DocProp{
					ID:          1,
					Name:        "Document Property",
					Description: "This is a document property",
				},
				CNvGraphicFramePr: &NonVisualGraphicFrameProp{
					GraphicFrameLocks: &GraphicFrameLocks{
						NoChangeAspect: dmlst.NewOptBool(true),
					},
				},
				Graphic: *DefaultGraphic(),
			},
			expectedXML: `<wp:inline xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:pic="http://schemas.openxmlformats.org/drawingml/2006/picture" distT="2" distB="3" distL="4" distR="5"><wp:extent cx="100" cy="200"></wp:extent><wp:docPr id="1" name="Document Property" descr="This is a document property"></wp:docPr><wp:cNvGraphicFramePr><a:graphicFrameLocks xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" noChangeAspect="1"></a:graphicFrameLocks></wp:cNvGraphicFramePr><a:graphic xmlns:a="` + constants.DrawingMLMainNS + `"></a:graphic></wp:inline>`,
			xmlName:     "wp:inline",
		},
	}

	for _, tt := range tests {
		t.Run(tt.xmlName, func(t *testing.T) {
			generatedXML, err := xml.Marshal(tt.inline)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			if string(generatedXML) != tt.expectedXML {
				t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expectedXML, generatedXML)
			}
		})
	}
}

func TestUnmarshalInline(t *testing.T) {

	tests := []struct {
		inputXML string
		expected Inline
	}{
		{
			inputXML: `<wp:inline xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:pic="http://schemas.openxmlformats.org/drawingml/2006/picture" distT="2" distB="3" distL="4" distR="5"><wp:extent cx="100" cy="200"></wp:extent><wp:docPr id="1" name="Document Property" descr="This is a document property"></wp:docPr><wp:cNvGraphicFramePr><a:graphicFrameLocks xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" noChangeAspect="1"></a:graphicFrameLocks></wp:cNvGraphicFramePr><a:graphic xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main"></a:graphic></wp:inline>`,
			expected: Inline{
				DistT: 2,
				DistB: 3,
				DistL: 4,
				DistR: 5,
				Extent: dmlct.PSize2D{
					Width:  100,
					Height: 200,
				},
				DocProp: DocProp{
					ID:          1,
					Name:        "Document Property",
					Description: "This is a document property",
				},
				CNvGraphicFramePr: &NonVisualGraphicFrameProp{
					GraphicFrameLocks: &GraphicFrameLocks{
						NoChangeAspect: dmlst.NewOptBool(true),
					},
				},
				Graphic: *DefaultGraphic(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			var inline Inline

			err := xml.Unmarshal([]byte(tt.inputXML), &inline)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if inline.DistT != tt.expected.DistT {
				t.Errorf("Expected DistT %d, but got %v", tt.expected.DistT, inline.DistT)
			}
			if inline.DistB != tt.expected.DistB {
				t.Errorf("Expected DistB %d, but got %v", tt.expected.DistB, inline.DistB)
			}
			if inline.DistL != tt.expected.DistL {
				t.Errorf("Expected DistL %d, but got %v", tt.expected.DistL, inline.DistL)
			}
			if inline.DistR != tt.expected.DistR {
				t.Errorf("Expected DistR %d, but got %v", tt.expected.DistR, inline.DistR)
			}

			if inline.Extent != tt.expected.Extent {
				t.Errorf("Expected Extent %+v, but got %+v", tt.expected.Extent, inline.Extent)
			}

			if inline.DocProp != tt.expected.DocProp {
				t.Errorf("Expected DocProp %+v, but got %+v", tt.expected.DocProp, inline.DocProp)
			}

			if inline.CNvGraphicFramePr == nil || inline.CNvGraphicFramePr.GraphicFrameLocks == nil ||
				inline.CNvGraphicFramePr.GraphicFrameLocks.NoChangeAspect != tt.expected.CNvGraphicFramePr.GraphicFrameLocks.NoChangeAspect {
				t.Errorf("Expected CNvGraphicFramePr.GraphicFrameLocks.NoChangeAspect %v, but got %v",
					tt.expected.CNvGraphicFramePr.GraphicFrameLocks.NoChangeAspect,
					inline.CNvGraphicFramePr.GraphicFrameLocks.NoChangeAspect)
			}

			if inline.Graphic != tt.expected.Graphic {
				t.Errorf("Expected Graphic %+v, but got %+v", tt.expected.Graphic, inline.Graphic)
			}
		})
	}
}
