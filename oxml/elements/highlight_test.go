package elements

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestHighlight(t *testing.T) {
	testHighlight := NewHighlight("yellow")

	xmlData, err := xml.Marshal(testHighlight)
	if err != nil {
		t.Fatalf("Error marshaling Highlight to XML: %v", err)
	}

	var unmarshaledHighlight Highlight
	err = xml.Unmarshal(xmlData, &unmarshaledHighlight)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to Highlight: %v", err)
	}

	if testHighlight.Value != unmarshaledHighlight.Value {
		t.Errorf("Expected highlight value %s, got %s", testHighlight.Value, unmarshaledHighlight.Value)
	}

	expectedXMLString := `<w:highlight w:val="yellow">`
	if !strings.Contains(string(xmlData), expectedXMLString) {
		t.Errorf("Expected XML string %s, got %s", expectedXMLString, string(xmlData))
	}
}
