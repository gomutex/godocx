package shapes

import (
	"encoding/xml"
	"testing"

	"github.com/gomutex/godocx/dml/dmlct"
)

func TestMarshalStretch(t *testing.T) {
	tests := []struct {
		name        string
		stretch     *Stretch
		expectedXML string
	}{
		{
			name:        "With FillRect",
			stretch:     &Stretch{FillRect: &dmlct.RelativeRect{}},
			expectedXML: `<a:stretch><a:fillRect></a:fillRect></a:stretch>`,
		},
		{
			name:        "Without FillRect",
			stretch:     &Stretch{},
			expectedXML: `<a:stretch></a:stretch>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generatedXML, err := xml.Marshal(tt.stretch)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			if string(generatedXML) != tt.expectedXML {
				t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expectedXML, generatedXML)
			}
		})
	}
}

func TestUnmarshalStretch(t *testing.T) {
	tests := []struct {
		name           string
		inputXML       string
		expectedResult Stretch
	}{
		{
			name:     "With FillRect",
			inputXML: `<a:stretch><a:fillRect></a:fillRect></a:stretch>`,
			expectedResult: Stretch{
				FillRect: &dmlct.RelativeRect{},
			},
		},
		{
			name:           "Without FillRect",
			inputXML:       `<a:stretch></a:stretch>`,
			expectedResult: Stretch{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Stretch

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if (result.FillRect == nil) != (tt.expectedResult.FillRect == nil) {
				t.Errorf("Expected FillRect to be %v, but got %v", tt.expectedResult.FillRect, result.FillRect)
			}
		})
	}
}
