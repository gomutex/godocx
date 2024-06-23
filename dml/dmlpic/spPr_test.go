package dmlpic

import (
	"encoding/xml"
	"testing"
)

func TestMarshalPicShapeProp(t *testing.T) {
	bwMode := BlackWhiteModeGray

	tests := []struct {
		picShapeProp *PicShapeProp
		expectedXML  string
	}{
		{
			picShapeProp: &PicShapeProp{
				BwMode:         &bwMode,
				TransformGroup: &TransformGroup{},
				PresetGeometry: &PresetGeometry{
					Preset: "rect",
				},
			},
			expectedXML: `<pic:spPr bwMode="gray"><a:xfrm></a:xfrm><a:prstGeom prst="rect"></a:prstGeom></pic:spPr>`,
		},
		{
			picShapeProp: &PicShapeProp{},
			expectedXML:  `<pic:spPr></pic:spPr>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.expectedXML, func(t *testing.T) {
			generatedXML, err := xml.Marshal(tt.picShapeProp)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			if string(generatedXML) != tt.expectedXML {
				t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expectedXML, generatedXML)
			}
		})
	}
}

func TestUnmarshalPicShapeProp(t *testing.T) {
	bwMode := BlackWhiteModeGray

	tests := []struct {
		inputXML       string
		expectedResult PicShapeProp
	}{
		{
			inputXML: `<pic:spPr bwMode="gray"><a:xfrm></a:xfrm><a:prstGeom></a:prstGeom></pic:spPr>`,
			expectedResult: PicShapeProp{
				BwMode:         &bwMode,
				TransformGroup: &TransformGroup{},
				PresetGeometry: &PresetGeometry{},
			},
		},
		{
			inputXML:       `<pic:spPr></pic:spPr>`,
			expectedResult: PicShapeProp{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			var result PicShapeProp

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if result.BwMode == nil && tt.expectedResult.BwMode != nil {
				t.Errorf("Expected bwMode to be %v, but got nil", *tt.expectedResult.BwMode)
			} else if result.BwMode != nil && tt.expectedResult.BwMode == nil {
				t.Errorf("Expected bwMode to be nil, but got %v", *result.BwMode)
			} else if result.BwMode != nil && tt.expectedResult.BwMode != nil && *result.BwMode != *tt.expectedResult.BwMode {
				t.Errorf("Expected bwMode %v, but got %v", *tt.expectedResult.BwMode, *result.BwMode)
			}

			// Check TransformGroup and PresetGeometry
			if (result.TransformGroup == nil) != (tt.expectedResult.TransformGroup == nil) {
				t.Errorf("Expected TransformGroup to be %v, but got %v", tt.expectedResult.TransformGroup, result.TransformGroup)
			}
			if (result.PresetGeometry == nil) != (tt.expectedResult.PresetGeometry == nil) {
				t.Errorf("Expected PresetGeometry to be %v, but got %v", tt.expectedResult.PresetGeometry, result.PresetGeometry)
			}
		})
	}
}
