package dml

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestMarshalDocProp(t *testing.T) {
	tests := []struct {
		docProp     *DocProp
		expectedXML string
	}{
		{
			docProp: &DocProp{
				ID:          1,
				Name:        "Document1",
				Description: "Description of Document1",
			},
			expectedXML: `<wp:docPr id="1" name="Document1" descr="Description of Document1"></wp:docPr>`,
		},
		{
			docProp: &DocProp{
				ID:   2,
				Name: "Document2",
			},
			expectedXML: `<wp:docPr id="2" name="Document2"></wp:docPr>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.expectedXML, func(t *testing.T) {
			generatedXML, err := xml.Marshal(tt.docProp)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			if strings.TrimSpace(string(generatedXML)) != tt.expectedXML {
				t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expectedXML, generatedXML)
			}
		})
	}
}

func TestUnmarshalDocProp(t *testing.T) {
	tests := []struct {
		inputXML        string
		expectedDocProp DocProp
	}{
		{
			inputXML: `<wp:docPr id="1" name="Document1" descr="Description of Document1"></wp:docPr>`,
			expectedDocProp: DocProp{
				ID:          1,
				Name:        "Document1",
				Description: "Description of Document1",
			},
		},
		{
			inputXML: `<wp:docPr id="2" name="Document2"></wp:docPr>`,
			expectedDocProp: DocProp{
				ID:   2,
				Name: "Document2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			var docProp DocProp

			err := xml.Unmarshal([]byte(tt.inputXML), &docProp)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if docProp.ID != tt.expectedDocProp.ID {
				t.Errorf("Expected ID %d, but got %d", tt.expectedDocProp.ID, docProp.ID)
			}
			if docProp.Name != tt.expectedDocProp.Name {
				t.Errorf("Expected Name %s, but got %s", tt.expectedDocProp.Name, docProp.Name)
			}
			if docProp.Description != tt.expectedDocProp.Description {
				t.Errorf("Expected Description %s, but got %s", tt.expectedDocProp.Description, docProp.Description)
			}
		})
	}
}
