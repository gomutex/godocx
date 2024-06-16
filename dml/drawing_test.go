package dml

import (
	"encoding/xml"
	"testing"

	"github.com/gomutex/godocx/common/constants"
	"github.com/gomutex/godocx/types"
)

func TestMarshalDrawing(t *testing.T) {
	simplePosAttr := 1
	var distTAttr uint = 2
	var distBAttr uint = 3
	var distLAttr uint = 4
	var distRAttr uint = 5
	layoutInCellAttr := 6
	allowOverlapAttr := 7
	relativeHeightAttr := 8
	behindDocAttr := 9
	lockedAttr := 10

	tests := []struct {
		drawing     *Drawing
		expectedXML string
		xmlName     string
	}{
		{
			drawing: &Drawing{
				Inline: []*Inline{
					{
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
				Anchor: []*Anchor{
					{
						SimplePosAttr:      &simplePosAttr,
						DistTAttr:          &distTAttr,
						DistBAttr:          &distBAttr,
						DistLAttr:          &distLAttr,
						DistRAttr:          &distRAttr,
						LayoutInCellAttr:   &layoutInCellAttr,
						AllowOverlapAttr:   &allowOverlapAttr,
						RelativeHeightAttr: &relativeHeightAttr,
						BehindDocAttr:      &behindDocAttr,
						LockedAttr:         &lockedAttr,
						Extent: &Extent{
							Width:  100,
							Height: 200,
						},
						EffectExtent: &EffectExtent{
							LeftEdge:   1,
							TopEdge:    2,
							RightEdge:  3,
							BottomEdge: 4,
						},
						WrapNone: &WrapNone{},
					},
				},
			},
			expectedXML: `<w:drawing><wp:anchor behindDoc="9" distT="2" distB="3" distL="4" distR="5" simplePos="1" locked="10" layoutInCell="6" allowOverlap="7" relativeHeight="8"><wp:extent cx="100" cy="200"></wp:extent><wp:effectExtent l="1" t="2" r="3" b="4"></wp:effectExtent><wp:wrapNone></wp:wrapNone></wp:anchor><wp:inline distT="2" distB="3" distL="4" distR="5"><wp:extent cx="100" cy="200"></wp:extent><wp:docPr id="1" name="Document Property" descr="This is a document property"></wp:docPr><wp:cNvGraphicFramePr><a:graphicFrameLocks xmlns:a="` + constants.DrawingMLMainNS + `" noChangeAspect="1"></a:graphicFrameLocks></wp:cNvGraphicFramePr><a:graphic xmlns:a="` + constants.DrawingMLMainNS + `"></a:graphic></wp:inline></w:drawing>`,
			xmlName:     "w:drawing",
		},
	}

	for _, tt := range tests {
		t.Run(tt.xmlName, func(t *testing.T) {
			generatedXML, err := xml.Marshal(tt.drawing)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			if string(generatedXML) != tt.expectedXML {
				t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expectedXML, generatedXML)
			}
		})
	}
}

func TestUnmarshalDrawing(t *testing.T) {
	var distTAttr uint = 2
	var distBAttr uint = 3
	var distLAttr uint = 4
	var distRAttr uint = 5
	var simplePosAttr int = 1
	var layoutInCellAttr int = 6
	var allowOverlapAttr int = 7
	var relativeHeightAttr int = 8
	var behindDocAttr int = 9
	var lockedAttr int = 10

	tests := []struct {
		inputXML        string
		expectedDrawing Drawing
	}{
		{
			inputXML: `<w:drawing xmlns:wp="http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"><wp:inline distT="2" distB="3" distL="4" distR="5"><wp:extent cx="100" cy="200"></wp:extent><wp:docPr id="1" name="Document Property" descr="This is a document property"></wp:docPr><wp:cNvGraphicFramePr><a:graphicFrameLocks xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" noChangeAspect="1"></a:graphicFrameLocks></wp:cNvGraphicFramePr><a:graphic xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main"></a:graphic></wp:inline><wp:anchor behindDoc="9" distT="2" distB="3" distL="4" distR="5" simplePos="1" locked="10" layoutInCell="6" allowOverlap="7" relativeHeight="8"><wp:extent cx="100" cy="200"></wp:extent><wp:effectExtent l="1" t="2" r="3" b="4"></wp:effectExtent><wp:wrapNone></wp:wrapNone></wp:anchor></w:drawing>`,
			expectedDrawing: Drawing{
				Inline: []*Inline{
					{
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
				Anchor: []*Anchor{
					{
						SimplePosAttr:      &simplePosAttr,
						DistTAttr:          &distTAttr,
						DistBAttr:          &distBAttr,
						DistLAttr:          &distLAttr,
						DistRAttr:          &distRAttr,
						LayoutInCellAttr:   &layoutInCellAttr,
						AllowOverlapAttr:   &allowOverlapAttr,
						RelativeHeightAttr: &relativeHeightAttr,
						BehindDocAttr:      &behindDocAttr,
						LockedAttr:         &lockedAttr,
						Extent: &Extent{
							Width:  100,
							Height: 200,
						},
						EffectExtent: &EffectExtent{
							LeftEdge:   1,
							TopEdge:    2,
							RightEdge:  3,
							BottomEdge: 4,
						},
						WrapNone: &WrapNone{},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			var drawing Drawing

			err := xml.Unmarshal([]byte(tt.inputXML), &drawing)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if len(drawing.Inline) != len(tt.expectedDrawing.Inline) {
				t.Errorf("Expected %d Inline elements, but got %d", len(tt.expectedDrawing.Inline), len(drawing.Inline))
			} else {
				for i := range drawing.Inline {
					if drawing.Inline[i].DistTAttr == nil || *drawing.Inline[i].DistTAttr != *tt.expectedDrawing.Inline[i].DistTAttr {
						t.Errorf("Expected Inline DistTAttr %d, but got %v", *tt.expectedDrawing.Inline[i].DistTAttr, drawing.Inline[i].DistTAttr)
					}
					if drawing.Inline[i].DistBAttr == nil || *drawing.Inline[i].DistBAttr != *tt.expectedDrawing.Inline[i].DistBAttr {
						t.Errorf("Expected Inline DistBAttr %d, but got %v", *tt.expectedDrawing.Inline[i].DistBAttr, drawing.Inline[i].DistBAttr)
					}
					if drawing.Inline[i].DistLAttr == nil || *drawing.Inline[i].DistLAttr != *tt.expectedDrawing.Inline[i].DistLAttr {
						t.Errorf("Expected Inline DistLAttr %d, but got %v", *tt.expectedDrawing.Inline[i].DistLAttr, drawing.Inline[i].DistLAttr)
					}
					if drawing.Inline[i].DistRAttr == nil || *drawing.Inline[i].DistRAttr != *tt.expectedDrawing.Inline[i].DistRAttr {
						t.Errorf("Expected Inline DistRAttr %d, but got %v", *tt.expectedDrawing.Inline[i].DistRAttr, drawing.Inline[i].DistRAttr)
					}
					if drawing.Inline[i].Extent.Width != tt.expectedDrawing.Inline[i].Extent.Width {
						t.Errorf("Expected Inline Extent Width %d, but got %d", tt.expectedDrawing.Inline[i].Extent.Width, drawing.Inline[i].Extent.Width)
					}
					if drawing.Inline[i].Extent.Height != tt.expectedDrawing.Inline[i].Extent.Height {
						t.Errorf("Expected Inline Extent Height %d, but got %d", tt.expectedDrawing.Inline[i].Extent.Height, drawing.Inline[i].Extent.Height)
					}
					if drawing.Inline[i].DocProp.ID != tt.expectedDrawing.Inline[i].DocProp.ID {
						t.Errorf("Expected Inline DocProp ID %d, but got %d", tt.expectedDrawing.Inline[i].DocProp.ID, drawing.Inline[i].DocProp.ID)
					}
					if drawing.Inline[i].DocProp.Name != tt.expectedDrawing.Inline[i].DocProp.Name {
						t.Errorf("Expected Inline DocProp Name %s, but got %s", tt.expectedDrawing.Inline[i].DocProp.Name, drawing.Inline[i].DocProp.Name)
					}
					if drawing.Inline[i].DocProp.Description != tt.expectedDrawing.Inline[i].DocProp.Description {
						t.Errorf("Expected Inline DocProp Description %s, but got %s", tt.expectedDrawing.Inline[i].DocProp.Description, drawing.Inline[i].DocProp.Description)
					}
					if drawing.Inline[i].CNvGraphicFramePr.GraphicFrameLocks.NoChangeAspect != tt.expectedDrawing.Inline[i].CNvGraphicFramePr.GraphicFrameLocks.NoChangeAspect {
						t.Errorf("Expected Inline CNvGraphicFramePr GraphicFrameLocks NoChangeAspect %v, but got %v", tt.expectedDrawing.Inline[i].CNvGraphicFramePr.GraphicFrameLocks.NoChangeAspect, drawing.Inline[i].CNvGraphicFramePr.GraphicFrameLocks.NoChangeAspect)
					}
				}
			}

			if len(drawing.Anchor) != len(tt.expectedDrawing.Anchor) {
				t.Errorf("Expected %d Anchor elements, but got %d", len(tt.expectedDrawing.Anchor), len(drawing.Anchor))
			} else {
				for i := range drawing.Anchor {
					if drawing.Anchor[i].SimplePosAttr == nil || *drawing.Anchor[i].SimplePosAttr != *tt.expectedDrawing.Anchor[i].SimplePosAttr {
						t.Errorf("Expected Anchor SimplePosAttr %d, but got %v", *tt.expectedDrawing.Anchor[i].SimplePosAttr, drawing.Anchor[i].SimplePosAttr)
					}
					if drawing.Anchor[i].DistTAttr == nil || *drawing.Anchor[i].DistTAttr != *tt.expectedDrawing.Anchor[i].DistTAttr {
						t.Errorf("Expected Anchor DistTAttr %d, but got %v", *tt.expectedDrawing.Anchor[i].DistTAttr, drawing.Anchor[i].DistTAttr)
					}
					if drawing.Anchor[i].DistBAttr == nil || *drawing.Anchor[i].DistBAttr != *tt.expectedDrawing.Anchor[i].DistBAttr {
						t.Errorf("Expected Anchor DistBAttr %d, but got %v", *tt.expectedDrawing.Anchor[i].DistBAttr, drawing.Anchor[i].DistBAttr)
					}
					if drawing.Anchor[i].DistLAttr == nil || *drawing.Anchor[i].DistLAttr != *tt.expectedDrawing.Anchor[i].DistLAttr {
						t.Errorf("Expected Anchor DistLAttr %d, but got %v", *tt.expectedDrawing.Anchor[i].DistLAttr, drawing.Anchor[i].DistLAttr)
					}
					if drawing.Anchor[i].DistRAttr == nil || *drawing.Anchor[i].DistRAttr != *tt.expectedDrawing.Anchor[i].DistRAttr {
						t.Errorf("Expected Anchor DistRAttr %d, but got %v", *tt.expectedDrawing.Anchor[i].DistRAttr, drawing.Anchor[i].DistRAttr)
					}
					if drawing.Anchor[i].LayoutInCellAttr == nil || *drawing.Anchor[i].LayoutInCellAttr != *tt.expectedDrawing.Anchor[i].LayoutInCellAttr {
						t.Errorf("Expected Anchor LayoutInCellAttr %d, but got %v", *tt.expectedDrawing.Anchor[i].LayoutInCellAttr, drawing.Anchor[i].LayoutInCellAttr)
					}
					if drawing.Anchor[i].AllowOverlapAttr == nil || *drawing.Anchor[i].AllowOverlapAttr != *tt.expectedDrawing.Anchor[i].AllowOverlapAttr {
						t.Errorf("Expected Anchor AllowOverlapAttr %d, but got %v", *tt.expectedDrawing.Anchor[i].AllowOverlapAttr, drawing.Anchor[i].AllowOverlapAttr)
					}
					if drawing.Anchor[i].RelativeHeightAttr == nil || *drawing.Anchor[i].RelativeHeightAttr != *tt.expectedDrawing.Anchor[i].RelativeHeightAttr {
						t.Errorf("Expected Anchor RelativeHeightAttr %d, but got %v", *tt.expectedDrawing.Anchor[i].RelativeHeightAttr, drawing.Anchor[i].RelativeHeightAttr)
					}
					if drawing.Anchor[i].BehindDocAttr == nil || *drawing.Anchor[i].BehindDocAttr != *tt.expectedDrawing.Anchor[i].BehindDocAttr {
						t.Errorf("Expected Anchor BehindDocAttr %d, but got %v", *tt.expectedDrawing.Anchor[i].BehindDocAttr, drawing.Anchor[i].BehindDocAttr)
					}
					if drawing.Anchor[i].LockedAttr == nil || *drawing.Anchor[i].LockedAttr != *tt.expectedDrawing.Anchor[i].LockedAttr {
						t.Errorf("Expected Anchor LockedAttr %d, but got %v", *tt.expectedDrawing.Anchor[i].LockedAttr, drawing.Anchor[i].LockedAttr)
					}
					if drawing.Anchor[i].Extent.Width != tt.expectedDrawing.Anchor[i].Extent.Width {
						t.Errorf("Expected Anchor Extent Width %d, but got %d", tt.expectedDrawing.Anchor[i].Extent.Width, drawing.Anchor[i].Extent.Width)
					}
					if drawing.Anchor[i].Extent.Height != tt.expectedDrawing.Anchor[i].Extent.Height {
						t.Errorf("Expected Anchor Extent Height %d, but got %d", tt.expectedDrawing.Anchor[i].Extent.Height, drawing.Anchor[i].Extent.Height)
					}
					if drawing.Anchor[i].EffectExtent.LeftEdge != tt.expectedDrawing.Anchor[i].EffectExtent.LeftEdge {
						t.Errorf("Expected Anchor EffectExtent LeftEdge %d, but got %d", tt.expectedDrawing.Anchor[i].EffectExtent.LeftEdge, drawing.Anchor[i].EffectExtent.LeftEdge)
					}
					if drawing.Anchor[i].EffectExtent.TopEdge != tt.expectedDrawing.Anchor[i].EffectExtent.TopEdge {
						t.Errorf("Expected Anchor EffectExtent TopEdge %d, but got %d", tt.expectedDrawing.Anchor[i].EffectExtent.TopEdge, drawing.Anchor[i].EffectExtent.TopEdge)
					}
					if drawing.Anchor[i].EffectExtent.RightEdge != tt.expectedDrawing.Anchor[i].EffectExtent.RightEdge {
						t.Errorf("Expected Anchor EffectExtent RightEdge %d, but got %d", tt.expectedDrawing.Anchor[i].EffectExtent.RightEdge, drawing.Anchor[i].EffectExtent.RightEdge)
					}
					if drawing.Anchor[i].EffectExtent.BottomEdge != tt.expectedDrawing.Anchor[i].EffectExtent.BottomEdge {
						t.Errorf("Expected Anchor EffectExtent BottomEdge %d, but got %d", tt.expectedDrawing.Anchor[i].EffectExtent.BottomEdge, drawing.Anchor[i].EffectExtent.BottomEdge)
					}
				}
			}
		})
	}
}
