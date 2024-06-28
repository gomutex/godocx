package dml

import (
	"encoding/xml"
	"testing"

	"github.com/gomutex/godocx/dml/dmlct"
	"github.com/gomutex/godocx/dml/dmlst"
)

func TestMarshalAnchor(t *testing.T) {
	simplePos := 1
	layoutInCell := 6
	allowOverlap := 7
	relativeHeight := 8
	behindDoc := 9
	locked := 10

	tests := []struct {
		anchor      *Anchor
		expectedXML string
		xmlName     string
	}{
		{
			anchor: &Anchor{
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
			expectedXML: `<wp:anchor behindDoc="9" distT="2" distB="3" distL="4" distR="5" simplePos="1" locked="10" layoutInCell="6" allowOverlap="7" relativeHeight="8"><wp:simplePos x="0" y="0"></wp:simplePos><wp:positionH relativeFrom="column"><wp:posOffset>0</wp:posOffset></wp:positionH><wp:positionV relativeFrom="line"><wp:posOffset>0</wp:posOffset></wp:positionV><wp:extent cx="100" cy="200"></wp:extent><wp:effectExtent l="1" t="2" r="3" b="4"></wp:effectExtent><wp:wrapNone></wp:wrapNone><wp:docPr id="1" name="test"></wp:docPr><a:graphic xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main"></a:graphic></wp:anchor>`,
			xmlName:     "wp:anchor",
		},
	}

	for _, tt := range tests {
		t.Run(tt.xmlName, func(t *testing.T) {
			generatedXML, err := xml.Marshal(tt.anchor)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			if string(generatedXML) != tt.expectedXML {
				t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expectedXML, generatedXML)
			}
		})
	}
}

func TestUnmarshalAnchor(t *testing.T) {

	simplePos := 1
	layoutInCell := 6
	allowOverlap := 7
	relativeHeight := 8
	behindDoc := 9
	locked := 10

	tests := []struct {
		inputXML string
		expected Anchor
	}{
		{
			inputXML: `<wp:anchor behindDoc="9" distT="2" distB="3" distL="4" distR="5" simplePos="1" locked="10" layoutInCell="6" allowOverlap="7" relativeHeight="8"><wp:simplePos x="0" y="0"></wp:simplePos><wp:positionH relativeFrom="column"><wp:posOffset>0</wp:posOffset></wp:positionH><wp:positionV relativeFrom="line"><wp:posOffset>0</wp:posOffset></wp:positionV><wp:extent cx="100" cy="200"></wp:extent><wp:effectExtent l="1" t="2" r="3" b="4"></wp:effectExtent><wp:wrapNone></wp:wrapNone><wp:docPr></wp:docPr><a:graphic></a:graphic></wp:anchor>`,
			expected: Anchor{
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
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			var anchor Anchor

			err := xml.Unmarshal([]byte(tt.inputXML), &anchor)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if *anchor.SimplePosAttr != *tt.expected.SimplePosAttr {
				t.Errorf("Expected SimplePos %d, but got %d", *tt.expected.SimplePosAttr, *anchor.SimplePosAttr)
			}
			if anchor.DistT != tt.expected.DistT {
				t.Errorf("Expected DistT %d, but got %d", tt.expected.DistT, anchor.DistT)
			}
			if anchor.DistB != tt.expected.DistB {
				t.Errorf("Expected DistB %d, but got %d", tt.expected.DistB, anchor.DistB)
			}
			if anchor.DistL != tt.expected.DistL {
				t.Errorf("Expected DistL %d, but got %d", tt.expected.DistL, anchor.DistL)
			}
			if anchor.DistR != tt.expected.DistR {
				t.Errorf("Expected DistR %d, but got %d", tt.expected.DistR, anchor.DistR)
			}
			if anchor.LayoutInCell != tt.expected.LayoutInCell {
				t.Errorf("Expected LayoutInCell %d, but got %d", tt.expected.LayoutInCell, anchor.LayoutInCell)
			}
			if anchor.AllowOverlap != tt.expected.AllowOverlap {
				t.Errorf("Expected AllowOverlap %d, but got %d", tt.expected.AllowOverlap, anchor.AllowOverlap)
			}
			if anchor.RelativeHeight != tt.expected.RelativeHeight {
				t.Errorf("Expected RelativeHeight %d, but got %d", tt.expected.RelativeHeight, anchor.RelativeHeight)
			}

			if anchor.BehindDoc != tt.expected.BehindDoc {
				t.Errorf("Expected BehindDoc %d, but got %d", tt.expected.BehindDoc, anchor.BehindDoc)
			}

			if anchor.Locked != tt.expected.Locked {
				t.Errorf("Expected Locked %d, but got %d", tt.expected.Locked, anchor.Locked)
			}

			// Validate nested structs
			if anchor.Extent.Width != tt.expected.Extent.Width {
				t.Errorf("Expected Extent.Width %d, but got %d", tt.expected.Extent.Width, anchor.Extent.Width)
			}
			if anchor.Extent.Height != tt.expected.Extent.Height {
				t.Errorf("Expected Extent.Height %d, but got %d", tt.expected.Extent.Height, anchor.Extent.Height)
			}

			if anchor.EffectExtent != nil {
				if anchor.EffectExtent.LeftEdge != tt.expected.EffectExtent.LeftEdge {
					t.Errorf("Expected EffectExtent.LeftEdge %d, but got %d", tt.expected.EffectExtent.LeftEdge, anchor.EffectExtent.LeftEdge)
				}
				if anchor.EffectExtent.TopEdge != tt.expected.EffectExtent.TopEdge {
					t.Errorf("Expected EffectExtent.TopEdge %d, but got %d", tt.expected.EffectExtent.TopEdge, anchor.EffectExtent.TopEdge)
				}
				if anchor.EffectExtent.RightEdge != tt.expected.EffectExtent.RightEdge {
					t.Errorf("Expected EffectExtent.RightEdge %d, but got %d", tt.expected.EffectExtent.RightEdge, anchor.EffectExtent.RightEdge)
				}
				if anchor.EffectExtent.BottomEdge != tt.expected.EffectExtent.BottomEdge {
					t.Errorf("Expected EffectExtent.BottomEdge %d, but got %d", tt.expected.EffectExtent.BottomEdge, anchor.EffectExtent.BottomEdge)
				}
			} else if tt.expected.EffectExtent != nil {
				t.Errorf("Expected EffectExtent %v, but got nil", tt.expected.EffectExtent)
			}
		})
	}
}
