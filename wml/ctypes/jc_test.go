package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestJustification(t *testing.T) {
	testJustification, err := NewJustification("left")
	if err != nil {
		t.Fatal(err)
	}

	xmlData, err := xml.Marshal(testJustification)
	if err != nil {
		t.Fatalf("Error marshaling Justification to XML: %v", err)
	}

	var unmarshaledJustification Justification
	err = xml.Unmarshal(xmlData, &unmarshaledJustification)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to Justification: %v", err)
	}

	if testJustification.Val != unmarshaledJustification.Val {
		t.Errorf("Expected justification value %s, got %s", testJustification.Val, unmarshaledJustification.Val)
	}

	expectedXMLString := `<Justification w:val="left">`
	if !strings.Contains(string(xmlData), expectedXMLString) {
		t.Errorf("Expected XML string %s, got %s", expectedXMLString, string(xmlData))
	}
}
