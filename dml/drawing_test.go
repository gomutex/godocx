package dml

import (
	"encoding/xml"
	"testing"

	"github.com/gomutex/godocx/dml/dmlct"
	"github.com/gomutex/godocx/dml/dmlst"
)

func TestMarshalDrawing(t *testing.T) {
	simplePos := 1

	layoutInCell := 6
	allowOverlap := 7
	relativeHeight := 8
	behindDoc := 9
	locked := 10

	tests := []struct {
		drawing     *Drawing
		expectedXML string
		xmlName     string
	}{
		{
			drawing: &Drawing{
				Inline: []Inline{
					{
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
				Anchor: []*Anchor{
					{
						SimplePosAttr:  &simplePos,
						DistT:          2,
						DistB:          3,
						DistL:          4,
						DistR:          5,
						LayoutInCell:   layoutInCell,
						AllowOverlap:   allowOverlap,
						RelativeHeight: relativeHeight,
						BehindDoc:      behindDoc,
						Locked:         locked,
						Extent: dmlct.PSize2D{
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
						PositionH: PoistionH{
							RelativeFrom: dmlst.RelFromHColumn,
						},
						PositionV: PoistionV{
							RelativeFrom: dmlst.RelFromVLine,
						},
						DocProp: DocProp{
							ID:   1,
							Name: "test",
						},
					},
				},
			},
			expectedXML: `<w:drawing><wp:anchor behindDoc="9" distT="2" distB="3" distL="4" distR="5" simplePos="1" locked="10" layoutInCell="6" allowOverlap="7" relativeHeight="8"><wp:simplePos x="0" y="0"></wp:simplePos><wp:positionH relativeFrom="column"><wp:posOffset>0</wp:posOffset></wp:positionH><wp:positionV relativeFrom="line"><wp:posOffset>0</wp:posOffset></wp:positionV><wp:extent cx="100" cy="200"></wp:extent><wp:effectExtent l="1" t="2" r="3" b="4"></wp:effectExtent><wp:wrapNone></wp:wrapNone><wp:docPr id="1" name="test"></wp:docPr><a:graphic xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main"></a:graphic></wp:anchor><wp:inline xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:pic="http://schemas.openxmlformats.org/drawingml/2006/picture" distT="2" distB="3" distL="4" distR="5"><wp:extent cx="100" cy="200"></wp:extent><wp:docPr id="1" name="Document Property" descr="This is a document property"></wp:docPr><wp:cNvGraphicFramePr><a:graphicFrameLocks xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" noChangeAspect="1"></a:graphicFrameLocks></wp:cNvGraphicFramePr><a:graphic xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main"></a:graphic></wp:inline></w:drawing>`,
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
	var simplePos int = 1
	var layoutInCell int = 6
	var allowOverlap int = 7
	var relativeHeight int = 8
	var behindDoc int = 9
	var locked int = 10

	tests := []struct {
		inputXML        string
		expectedDrawing Drawing
	}{
		{
			inputXML: `<w:drawing xmlns:wp="http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"><wp:inline distT="2" distB="3" distL="4" distR="5"><wp:extent cx="100" cy="200"></wp:extent><wp:docPr id="1" name="Document Property" descr="This is a document property"></wp:docPr><wp:cNvGraphicFramePr><a:graphicFrameLocks  noChangeAspect="1"></a:graphicFrameLocks></wp:cNvGraphicFramePr><a:graphic xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main"></a:graphic></wp:inline><wp:anchor behindDoc="9" distT="2" distB="3" distL="4" distR="5" simplePos="1" locked="10" layoutInCell="6" allowOverlap="7" relativeHeight="8"><wp:extent cx="100" cy="200"></wp:extent><wp:effectExtent l="1" t="2" r="3" b="4"></wp:effectExtent><wp:wrapNone></wp:wrapNone></wp:anchor></w:drawing>`,
			expectedDrawing: Drawing{
				Inline: []Inline{
					{
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
				Anchor: []*Anchor{
					{
						SimplePosAttr:  &simplePos,
						DistT:          2,
						DistB:          3,
						DistL:          4,
						DistR:          5,
						LayoutInCell:   layoutInCell,
						AllowOverlap:   allowOverlap,
						RelativeHeight: relativeHeight,
						BehindDoc:      behindDoc,
						Locked:         locked,
						Extent: dmlct.PSize2D{
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
					if drawing.Inline[i].DistT != tt.expectedDrawing.Inline[i].DistT {
						t.Errorf("Expected Inline DistT %d, but got %v", tt.expectedDrawing.Inline[i].DistT, drawing.Inline[i].DistT)
					}
					if drawing.Inline[i].DistB != tt.expectedDrawing.Inline[i].DistB {
						t.Errorf("Expected Inline DistB %d, but got %v", tt.expectedDrawing.Inline[i].DistB, drawing.Inline[i].DistB)
					}
					if drawing.Inline[i].DistL != tt.expectedDrawing.Inline[i].DistL {
						t.Errorf("Expected Inline DistL %d, but got %v", tt.expectedDrawing.Inline[i].DistL, drawing.Inline[i].DistL)
					}
					if drawing.Inline[i].DistR != tt.expectedDrawing.Inline[i].DistR {
						t.Errorf("Expected Inline DistR %d, but got %v", tt.expectedDrawing.Inline[i].DistR, drawing.Inline[i].DistR)
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
					if drawing.Anchor[i].DistT != tt.expectedDrawing.Anchor[i].DistT {
						t.Errorf("Expected Anchor DistT %d, but got %v", tt.expectedDrawing.Anchor[i].DistT, drawing.Anchor[i].DistT)
					}
					if drawing.Anchor[i].DistB != tt.expectedDrawing.Anchor[i].DistB {
						t.Errorf("Expected Anchor DistB %d, but got %v", tt.expectedDrawing.Anchor[i].DistB, drawing.Anchor[i].DistB)
					}
					if drawing.Anchor[i].DistL != tt.expectedDrawing.Anchor[i].DistL {
						t.Errorf("Expected Anchor DistL %d, but got %v", tt.expectedDrawing.Anchor[i].DistL, drawing.Anchor[i].DistL)
					}
					if drawing.Anchor[i].DistR != tt.expectedDrawing.Anchor[i].DistR {
						t.Errorf("Expected Anchor DistR %d, but got %v", tt.expectedDrawing.Anchor[i].DistR, drawing.Anchor[i].DistR)
					}
					if drawing.Anchor[i].LayoutInCell != tt.expectedDrawing.Anchor[i].LayoutInCell {
						t.Errorf("Expected Anchor LayoutInCell %d, but got %v", tt.expectedDrawing.Anchor[i].LayoutInCell, drawing.Anchor[i].LayoutInCell)
					}
					if drawing.Anchor[i].AllowOverlap != tt.expectedDrawing.Anchor[i].AllowOverlap {
						t.Errorf("Expected Anchor AllowOverlap %d, but got %v", tt.expectedDrawing.Anchor[i].AllowOverlap, drawing.Anchor[i].AllowOverlap)
					}
					if drawing.Anchor[i].RelativeHeight != tt.expectedDrawing.Anchor[i].RelativeHeight {
						t.Errorf("Expected Anchor RelativeHeight %d, but got %v", tt.expectedDrawing.Anchor[i].RelativeHeight, drawing.Anchor[i].RelativeHeight)
					}
					if drawing.Anchor[i].BehindDoc != tt.expectedDrawing.Anchor[i].BehindDoc {
						t.Errorf("Expected Anchor BehindDoc %d, but got %v", tt.expectedDrawing.Anchor[i].BehindDoc, drawing.Anchor[i].BehindDoc)
					}
					if drawing.Anchor[i].Locked != tt.expectedDrawing.Anchor[i].Locked {
						t.Errorf("Expected Anchor Locked %d, but got %v", tt.expectedDrawing.Anchor[i].Locked, drawing.Anchor[i].Locked)
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
