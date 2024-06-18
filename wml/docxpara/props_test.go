package docxpara

import (
	"encoding/xml"
	"testing"

	"github.com/gomutex/godocx/wml/formatting"
)

func areParagraphPropertiesEqual(p1, p2 ParagraphProperty) bool {
	return p1.Style.Val == p2.Style.Val &&
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
		Justification: formatting.NewJustification("center"),
	}

	if !areParagraphPropertiesEqual(expectedParagraphProperty, parsedParagraphProperty) {
		t.Errorf("Expected ParagraphProperty %v, got %v", expectedParagraphProperty, parsedParagraphProperty)
	}
}

func TestNewParagraphStyle(t *testing.T) {
	expected := "TestStyle"
	style := NewParagraphStyle(expected)

	if style.Val != expected {
		t.Errorf("NewParagraphStyle() = %s; want %s", style.Val, expected)
	}
}

func TestDefaultParagraphStyle(t *testing.T) {
	expected := "Normal"
	style := DefaultParagraphStyle()

	if style.Val != expected {
		t.Errorf("DefaultParagraphStyle() = %s; want %s", style.Val, expected)
	}
}
