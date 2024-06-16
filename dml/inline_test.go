package dml

import (
	"encoding/xml"
	"testing"

	"github.com/gomutex/godocx/common/constants"
	"github.com/gomutex/godocx/types"
)

func TestMarshalInline(t *testing.T) {
	var distTAttr uint = 2
	var distBAttr uint = 3
	var distLAttr uint = 4
	var distRAttr uint = 5

	tests := []struct {
		inline      *Inline
		expectedXML string
		xmlName     string
	}{
		{
			inline: &Inline{
				DistTAttr: &distTAttr,
				DistBAttr: &distBAttr,
				DistLAttr: &distLAttr,
				DistRAttr: &distRAttr,
				Extent: &Extent{
					Width:  100,
					Height: 200,
				},
				DocProp: &DocProp{
					ID:          1,
					Name:        "Document Property",
					Description: "This is a document property",
				},
				CNvGraphicFramePr: &NonVisualGraphicFrameProp{
					GraphicFrameLocks: &GraphicFrameLocks{
						NoChangeAspect: types.NewNullBool(true),
					},
				},
				Graphic: DefaultGraphic(),
			},
			expectedXML: `<wp:inline distT="2" distB="3" distL="4" distR="5"><wp:extent cx="100" cy="200"></wp:extent><wp:docPr id="1" name="Document Property" descr="This is a document property"></wp:docPr><wp:cNvGraphicFramePr><a:graphicFrameLocks xmlns:a="` + constants.DrawingMLMainNS + `" noChangeAspect="1"></a:graphicFrameLocks></wp:cNvGraphicFramePr><a:graphic xmlns:a="` + constants.DrawingMLMainNS + `"></a:graphic></wp:inline>`,
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
	var distTAttr uint = 2
	var distBAttr uint = 3
	var distLAttr uint = 4
	var distRAttr uint = 5

	tests := []struct {
		inputXML string
		expected Inline
	}{
		{
			inputXML: `<wp:inline xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:pic="http://schemas.openxmlformats.org/drawingml/2006/picture" distT="2" distB="3" distL="4" distR="5"><wp:extent cx="100" cy="200"></wp:extent><wp:docPr id="1" name="Document Property" descr="This is a document property"></wp:docPr><wp:cNvGraphicFramePr><a:graphicFrameLocks noChangeAspect="1"></a:graphicFrameLocks></wp:cNvGraphicFramePr><a:graphic xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main"></a:graphic></wp:inline>`,
			expected: Inline{
				DistTAttr: &distTAttr,
				DistBAttr: &distBAttr,
				DistLAttr: &distLAttr,
				DistRAttr: &distRAttr,
				Extent: &Extent{
					Width:  100,
					Height: 200,
				},
				DocProp: &DocProp{
					ID:          1,
					Name:        "Document Property",
					Description: "This is a document property",
				},
				CNvGraphicFramePr: &NonVisualGraphicFrameProp{
					GraphicFrameLocks: &GraphicFrameLocks{
						NoChangeAspect: types.NewNullBool(true),
					},
				},
				Graphic: DefaultGraphic(),
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

			if inline.DistTAttr == nil || *inline.DistTAttr != *tt.expected.DistTAttr {
				t.Errorf("Expected DistTAttr %d, but got %v", *tt.expected.DistTAttr, inline.DistTAttr)
			}
			if inline.DistBAttr == nil || *inline.DistBAttr != *tt.expected.DistBAttr {
				t.Errorf("Expected DistBAttr %d, but got %v", *tt.expected.DistBAttr, inline.DistBAttr)
			}
			if inline.DistLAttr == nil || *inline.DistLAttr != *tt.expected.DistLAttr {
				t.Errorf("Expected DistLAttr %d, but got %v", *tt.expected.DistLAttr, inline.DistLAttr)
			}
			if inline.DistRAttr == nil || *inline.DistRAttr != *tt.expected.DistRAttr {
				t.Errorf("Expected DistRAttr %d, but got %v", *tt.expected.DistRAttr, inline.DistRAttr)
			}

			if inline.Extent == nil || *inline.Extent != *tt.expected.Extent {
				t.Errorf("Expected Extent %+v, but got %+v", tt.expected.Extent, inline.Extent)
			}

			if inline.DocProp == nil || *inline.DocProp != *tt.expected.DocProp {
				t.Errorf("Expected DocProp %+v, but got %+v", tt.expected.DocProp, inline.DocProp)
			}

			if inline.CNvGraphicFramePr == nil || inline.CNvGraphicFramePr.GraphicFrameLocks == nil ||
				inline.CNvGraphicFramePr.GraphicFrameLocks.NoChangeAspect != tt.expected.CNvGraphicFramePr.GraphicFrameLocks.NoChangeAspect {
				t.Errorf("Expected CNvGraphicFramePr.GraphicFrameLocks.NoChangeAspect %v, but got %v",
					tt.expected.CNvGraphicFramePr.GraphicFrameLocks.NoChangeAspect,
					inline.CNvGraphicFramePr.GraphicFrameLocks.NoChangeAspect)
			}

			if inline.Graphic == nil || *inline.Graphic != *tt.expected.Graphic {
				t.Errorf("Expected Graphic %+v, but got %+v", tt.expected.Graphic, inline.Graphic)
			}
		})
	}
}
