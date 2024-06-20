package dml

import (
	"encoding/xml"
	"testing"

	"github.com/gomutex/godocx/common/constants"
	"github.com/gomutex/godocx/dml/dmlpic"
)

func TestMarshalGraphic(t *testing.T) {
	tests := []struct {
		graphic     *Graphic
		expectedXML string
		xmlName     string
	}{
		{
			graphic:     NewPicGraphic(&dmlpic.Pic{}),
			expectedXML: `<a:graphic xmlns:a="` + constants.DrawingMLMainNS + `"><a:graphicData uri="` + constants.DrawingMLPicNS + `"><pic:pic xmlns:pic="` + constants.DrawingMLPicNS + `"></pic:pic></a:graphicData></a:graphic>`,
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
