package docxpara

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestParagraphStyle(t *testing.T) {
	testParagraphStyle := NewParagraphStyle("Heading1")

	xmlData, err := xml.Marshal(testParagraphStyle)
	if err != nil {
		t.Fatalf("Error marshaling ParagraphStyle to XML: %v", err)
	}

	var unmarshaledParagraphStyle ParagraphStyle
	err = xml.Unmarshal(xmlData, &unmarshaledParagraphStyle)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to ParagraphStyle: %v", err)
	}

	if testParagraphStyle.Value != unmarshaledParagraphStyle.Value {
		t.Errorf("Expected paragraph style value %s, got %s", testParagraphStyle.Value, unmarshaledParagraphStyle.Value)
	}

	expectedXMLString := `<w:pStyle w:val="Heading1"></w:pStyle>`
	if !strings.Contains(string(xmlData), expectedXMLString) {
		t.Errorf("Expected XML string %s, got %s", expectedXMLString, string(xmlData))
	}
}
