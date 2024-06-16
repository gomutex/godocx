package dml

import (
	"encoding/xml"
	"testing"
)

func TestMarshalAnchor(t *testing.T) {
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
		anchor      *Anchor
		expectedXML string
		xmlName     string
	}{
		{
			anchor: &Anchor{
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
			expectedXML: `<wp:anchor behindDoc="9" distT="2" distB="3" distL="4" distR="5" simplePos="1" locked="10" layoutInCell="6" allowOverlap="7" relativeHeight="8"><wp:extent cx="100" cy="200"></wp:extent><wp:effectExtent l="1" t="2" r="3" b="4"></wp:effectExtent><wp:wrapNone></wp:wrapNone></wp:anchor>`,
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
		inputXML string
		expected Anchor
	}{
		{
			inputXML: `<wp:anchor behindDoc="9" distT="2" distB="3" distL="4" distR="5" simplePos="1" locked="10" layoutInCell="6" allowOverlap="7" relativeHeight="8"><wp:extent cx="100" cy="200"></wp:extent><wp:effectExtent l="1" t="2" r="3" b="4"></wp:effectExtent><wp:wrapNone></wp:wrapNone></wp:anchor>`,
			expected: Anchor{
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
		{
			inputXML: `<wp:anchor distT="2" distB="3" distL="4" distR="5" simplePos="1" locked="10" layoutInCell="6" allowOverlap="7" relativeHeight="8"><wp:extent cx="100" cy="200"></wp:extent><wp:effectExtent l="1" t="2" r="3" b="4"></wp:effectExtent></wp:anchor>`,
			expected: Anchor{
				SimplePosAttr:      &simplePosAttr,
				DistTAttr:          &distTAttr,
				DistBAttr:          &distBAttr,
				DistLAttr:          &distLAttr,
				DistRAttr:          &distRAttr,
				LayoutInCellAttr:   &layoutInCellAttr,
				AllowOverlapAttr:   &allowOverlapAttr,
				RelativeHeightAttr: &relativeHeightAttr,
				BehindDocAttr:      nil,
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
				WrapNone: nil,
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
				t.Errorf("Expected SimplePosAttr %d, but got %d", *tt.expected.SimplePosAttr, *anchor.SimplePosAttr)
			}
			if *anchor.DistTAttr != *tt.expected.DistTAttr {
				t.Errorf("Expected DistTAttr %d, but got %d", *tt.expected.DistTAttr, *anchor.DistTAttr)
			}
			if *anchor.DistBAttr != *tt.expected.DistBAttr {
				t.Errorf("Expected DistBAttr %d, but got %d", *tt.expected.DistBAttr, *anchor.DistBAttr)
			}
			if *anchor.DistLAttr != *tt.expected.DistLAttr {
				t.Errorf("Expected DistLAttr %d, but got %d", *tt.expected.DistLAttr, *anchor.DistLAttr)
			}
			if *anchor.DistRAttr != *tt.expected.DistRAttr {
				t.Errorf("Expected DistRAttr %d, but got %d", *tt.expected.DistRAttr, *anchor.DistRAttr)
			}
			if *anchor.LayoutInCellAttr != *tt.expected.LayoutInCellAttr {
				t.Errorf("Expected LayoutInCellAttr %d, but got %d", *tt.expected.LayoutInCellAttr, *anchor.LayoutInCellAttr)
			}
			if *anchor.AllowOverlapAttr != *tt.expected.AllowOverlapAttr {
				t.Errorf("Expected AllowOverlapAttr %d, but got %d", *tt.expected.AllowOverlapAttr, *anchor.AllowOverlapAttr)
			}
			if *anchor.RelativeHeightAttr != *tt.expected.RelativeHeightAttr {
				t.Errorf("Expected RelativeHeightAttr %d, but got %d", *tt.expected.RelativeHeightAttr, *anchor.RelativeHeightAttr)
			}

			if tt.expected.BehindDocAttr != nil {
				if anchor.BehindDocAttr == nil {
					t.Errorf("Expected BehindDocAttr %d, but got nil", *tt.expected.BehindDocAttr)
				}
				if *anchor.BehindDocAttr != *tt.expected.BehindDocAttr {
					t.Errorf("Expected BehindDocAttr %d, but got %d", *tt.expected.BehindDocAttr, *anchor.BehindDocAttr)
				}
			} else {
				if anchor.BehindDocAttr != nil {
					t.Errorf("Expected BehindDocAttr nil, but got %d", *anchor.BehindDocAttr)
				}
			}

			if *anchor.LockedAttr != *tt.expected.LockedAttr {
				t.Errorf("Expected LockedAttr %d, but got %d", *tt.expected.LockedAttr, *anchor.LockedAttr)
			}

			// Validate nested structs
			if anchor.Extent != nil {
				if anchor.Extent.Width != tt.expected.Extent.Width {
					t.Errorf("Expected Extent.Width %d, but got %d", tt.expected.Extent.Width, anchor.Extent.Width)
				}
				if anchor.Extent.Height != tt.expected.Extent.Height {
					t.Errorf("Expected Extent.Height %d, but got %d", tt.expected.Extent.Height, anchor.Extent.Height)
				}
			} else if tt.expected.Extent != nil {
				t.Errorf("Expected Extent %v, but got nil", tt.expected.Extent)
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
