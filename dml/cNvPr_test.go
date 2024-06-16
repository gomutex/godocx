package dml

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestMarshalCNvPr(t *testing.T) {
	tests := []struct {
		cnvpr       *CNvPr
		expectedXML string
	}{
		{
			cnvpr: &CNvPr{
				ID:          "1",
				Name:        "Drawing1",
				Description: "Description of Drawing1",
			},
			expectedXML: `<pic:cNvPr id="1" name="Drawing1" descr="Description of Drawing1"></pic:cNvPr>`,
		},
		{
			cnvpr: &CNvPr{
				ID:   "2",
				Name: "Drawing2",
			},
			expectedXML: `<pic:cNvPr id="2" name="Drawing2"></pic:cNvPr>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.expectedXML, func(t *testing.T) {
			generatedXML, err := xml.Marshal(tt.cnvpr)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			if strings.TrimSpace(string(generatedXML)) != tt.expectedXML {
				t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expectedXML, generatedXML)
			}
		})
	}
}

func TestUnmarshalCNvPr(t *testing.T) {
	tests := []struct {
		inputXML      string
		expectedCNvPr CNvPr
	}{
		{
			inputXML: `<pic:cNvPr id="1" name="Drawing1" descr="Description of Drawing1"></pic:cNvPr>`,
			expectedCNvPr: CNvPr{
				ID:          "1",
				Name:        "Drawing1",
				Description: "Description of Drawing1",
			},
		},
		{
			inputXML: `<pic:cNvPr id="2" name="Drawing2"></pic:cNvPr>`,
			expectedCNvPr: CNvPr{
				ID:   "2",
				Name: "Drawing2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			var cnvpr CNvPr

			err := xml.Unmarshal([]byte(tt.inputXML), &cnvpr)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if cnvpr.ID != tt.expectedCNvPr.ID {
				t.Errorf("Expected ID %s, but got %s", tt.expectedCNvPr.ID, cnvpr.ID)
			}
			if cnvpr.Name != tt.expectedCNvPr.Name {
				t.Errorf("Expected Name %s, but got %s", tt.expectedCNvPr.Name, cnvpr.Name)
			}
			if cnvpr.Description != tt.expectedCNvPr.Description {
				t.Errorf("Expected Description %s, but got %s", tt.expectedCNvPr.Description, cnvpr.Description)
			}
		})
	}
}
