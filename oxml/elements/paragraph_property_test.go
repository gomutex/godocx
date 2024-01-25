package elements

import (
	"encoding/xml"
	"testing"
)

func areParagraphPropertiesEqual(p1, p2 ParagraphProperty) bool {
	return p1.Style.Value == p2.Style.Value &&
		p1.Justification.Value == p2.Justification.Value
}

func TestParagraphProperty(t *testing.T) {
	xmlString := `<w:pPr xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
  <w:pStyle w:val="Heading1"></w:pStyle>
  <w:jc w:val="center"></w:jc>
</w:pPr>`

	var parsedParagraphProperty ParagraphProperty
	err := xml.Unmarshal([]byte(xmlString), &parsedParagraphProperty)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to ParagraphProperty: %v", err)
	}

	expectedParagraphProperty := ParagraphProperty{
		Style:         NewParagraphStyle("Heading1"),
		Justification: NewJustification("center"),
	}

	if !areParagraphPropertiesEqual(expectedParagraphProperty, parsedParagraphProperty) {
		t.Errorf("Expected ParagraphProperty %v, got %v", expectedParagraphProperty, parsedParagraphProperty)
	}
}
