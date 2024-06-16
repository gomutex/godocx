package docxrun

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestRunStyle(t *testing.T) {
	testRunStyle := NewRunStyle("Heading")

	xmlData, err := xml.Marshal(testRunStyle)
	if err != nil {
		t.Fatalf("Error marshaling RunStyle to XML: %v", err)
	}

	expectedXMLString := `<w:rStyle w:val="Heading"></w:rStyle>`
	if !strings.Contains(string(xmlData), expectedXMLString) {
		t.Errorf("Expected XML string %s, got %s", expectedXMLString, string(xmlData))
	}
}
