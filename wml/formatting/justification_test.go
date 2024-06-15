package formatting

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestJustification(t *testing.T) {
	testJustification := NewJustification("left")

	xmlData, err := xml.Marshal(testJustification)
	if err != nil {
		t.Fatalf("Error marshaling Justification to XML: %v", err)
	}

	var unmarshaledJustification Justification
	err = xml.Unmarshal(xmlData, &unmarshaledJustification)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to Justification: %v", err)
	}

	if testJustification.Value != unmarshaledJustification.Value {
		t.Errorf("Expected justification value %s, got %s", testJustification.Value, unmarshaledJustification.Value)
	}

	expectedXMLString := `<w:jc w:val="left">`
	if !strings.Contains(string(xmlData), expectedXMLString) {
		t.Errorf("Expected XML string %s, got %s", expectedXMLString, string(xmlData))
	}
}
