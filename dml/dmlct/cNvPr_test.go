package dmlct

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
				ID:          1,
				Name:        "Drawing1",
				Description: "Description of Drawing1",
			},
			expectedXML: `<pic:cNvPr id="1" name="Drawing1" descr="Description of Drawing1"></pic:cNvPr>`,
		},
		{
			cnvpr: &CNvPr{
				ID:   2,
				Name: "Drawing2",
			},
			expectedXML: `<pic:cNvPr id="2" name="Drawing2" descr=""></pic:cNvPr>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.expectedXML, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)

			start := xml.StartElement{Name: xml.Name{Local: "pic:cNvPr"}}
			err := tt.cnvpr.MarshalXML(encoder, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			err = encoder.Flush()
			if err != nil {
				t.Errorf("Error flushing encoder: %v", err)
			}

			if result.String() != tt.expectedXML {
				t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expectedXML, result.String())
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
				ID:          1,
				Name:        "Drawing1",
				Description: "Description of Drawing1",
			},
		},
		{
			inputXML: `<pic:cNvPr id="2" name="Drawing2"></pic:cNvPr>`,
			expectedCNvPr: CNvPr{
				ID:   2,
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
				t.Errorf("Expected ID %d, but got %d", tt.expectedCNvPr.ID, cnvpr.ID)
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
