package wml

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestItalic(t *testing.T) {
	testItalic := NewItalic(true)

	xmlData, err := xml.Marshal(testItalic)
	if err != nil {
		t.Fatalf("Error marshaling Italic to XML: %v", err)
	}

	var unmarshaledItalic Italic
	err = xml.Unmarshal(xmlData, &unmarshaledItalic)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to Italic: %v", err)
	}

	if testItalic.Val != unmarshaledItalic.Val {
		t.Errorf("Expected italic value %t, got %t", testItalic.Val, unmarshaledItalic.Val)
	}

	expectedXMLString := `<w:i w:val="true"></w:i>`
	if !strings.Contains(string(xmlData), expectedXMLString) {
		t.Errorf("Expected XML string %s, got %s", expectedXMLString, string(xmlData))
	}
}
