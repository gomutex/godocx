package dml

import (
	"encoding/xml"
	"testing"

	"github.com/gomutex/godocx/common/constants"
	"github.com/gomutex/godocx/dml/dmlct"
	"github.com/gomutex/godocx/dml/dmlpic"
	"github.com/gomutex/godocx/dml/dmlprops"
	"github.com/gomutex/godocx/dml/dmlst"
	"github.com/gomutex/godocx/dml/shapes"
)

func TestMarshalGraphic(t *testing.T) {
	tests := []struct {
		graphic     *Graphic
		expectedXML string
		xmlName     string
	}{
		{
			graphic: NewPicGraphic(&dmlpic.Pic{
				NonVisualPicProp: dmlpic.NonVisualPicProp{
					CNvPr: dmlct.CNvPr{
						ID:          1,
						Name:        "Pic 1",
						Description: "Description",
					},
					CNvPicPr: dmlpic.CNvPicPr{
						PicLocks: &dmlprops.PicLocks{
							NoChangeAspect:     dmlst.NewOptBool(true),
							NoChangeArrowheads: dmlst.NewOptBool(true),
						},
					},
				},
				BlipFill: dmlpic.BlipFill{
					Blip: &dmlpic.Blip{
						EmbedID: "rId1",
					},
					FillModeProps: dmlpic.FillModeProps{
						Stretch: &shapes.Stretch{
							FillRect: &dmlct.RelativeRect{},
						},
					},
				},
				PicShapeProp: dmlpic.PicShapeProp{
					TransformGroup: &dmlpic.TransformGroup{
						Offset: &dmlpic.Offset{
							X: 0,
							Y: 0,
						},
						Extent: &dmlct.PSize2D{
							Width:  100000,
							Height: 100000,
						},
					},
					PresetGeometry: &dmlpic.PresetGeometry{
						Preset: "rect",
					},
				},
			}),
			expectedXML: `<a:graphic xmlns:a="` + constants.DrawingMLMainNS + `"><a:graphicData uri="` + constants.DrawingMLPicNS + `"><pic:pic xmlns:pic="` + constants.DrawingMLPicNS + `"><pic:nvPicPr><pic:cNvPr id="1" name="Pic 1" descr="Description"></pic:cNvPr><pic:cNvPicPr><a:picLocks noChangeAspect="1" noChangeArrowheads="1"></a:picLocks></pic:cNvPicPr></pic:nvPicPr><pic:blipFill><a:blip r:embed="rId1"></a:blip><a:stretch><a:fillRect></a:fillRect></a:stretch></pic:blipFill><pic:spPr><a:xfrm><a:off x="0" y="0"></a:off><a:ext cx="100000" cy="100000"></a:ext></a:xfrm><a:prstGeom prst="rect"></a:prstGeom></pic:spPr></pic:pic></a:graphicData></a:graphic>`,
			xmlName:     "a:graphic",
		},
		{
			graphic:     DefaultGraphic(),
			expectedXML: `<a:graphic xmlns:a="` + constants.DrawingMLMainNS + `"></a:graphic>`,
			xmlName:     "a:graphic",
		},
	}

	for _, tt := range tests {
		t.Run(tt.xmlName, func(t *testing.T) {
			generatedXML, err := xml.Marshal(tt.graphic)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			if string(generatedXML) != tt.expectedXML {
				t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expectedXML, generatedXML)
			}
		})
	}
}

func TestUnmarshalGraphic(t *testing.T) {
	tests := []struct {
		inputXML        string
		expectedGraphic Graphic
	}{
		{
			inputXML: `<a:graphic xmlns:a="` + constants.DrawingMLMainNS + `"><a:graphicData uri="` + constants.DrawingMLPicNS + `"><pic:pic></pic:pic></a:graphicData></a:graphic>`,
			expectedGraphic: Graphic{
				Data: &GraphicData{
					URI: constants.DrawingMLPicNS,
					Pic: &dmlpic.Pic{},
				},
			},
		},
		{
			inputXML: `<a:graphic xmlns:a="` + constants.DrawingMLMainNS + `"></a:graphic>`,
			expectedGraphic: Graphic{
				Data: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			var graphic Graphic

			err := xml.Unmarshal([]byte(tt.inputXML), &graphic)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if graphic.Data == nil && tt.expectedGraphic.Data != nil {
				t.Errorf("Expected Data to be %v, but got nil", tt.expectedGraphic.Data)
			} else if graphic.Data != nil && tt.expectedGraphic.Data == nil {
				t.Errorf("Expected Data to be nil, but got %v", graphic.Data)
			} else if graphic.Data != nil && tt.expectedGraphic.Data != nil {
				if graphic.Data.URI != tt.expectedGraphic.Data.URI {
					t.Errorf("Expected URI %s, but got %s", tt.expectedGraphic.Data.URI, graphic.Data.URI)
				}
				if graphic.Data.Pic == nil && tt.expectedGraphic.Data.Pic != nil {
					t.Errorf("Expected Pic to be %v, but got nil", tt.expectedGraphic.Data.Pic)
				} else if graphic.Data.Pic != nil && tt.expectedGraphic.Data.Pic == nil {
					t.Errorf("Expected Pic to be nil, but got %v", graphic.Data.Pic)
				}
			}
		})
	}
}
