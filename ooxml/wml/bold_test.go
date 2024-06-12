package wml

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestBold(t *testing.T) {
	testBold := NewBold(true)

	xmlData, err := xml.Marshal(testBold)
	if err != nil {
		t.Fatalf("Error marshaling Bold to XML: %v", err)
	}

	var unmarshaledBold Bold
	err = xml.Unmarshal(xmlData, &unmarshaledBold)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to Bold: %v", err)
	}

	if testBold.Val != unmarshaledBold.Val {
		t.Errorf("Expected bold value %t, got %t", testBold.Val, unmarshaledBold.Val)
	}

	expectedXMLString := `<w:b w:val="true"></w:b>`
	if !strings.Contains(string(xmlData), expectedXMLString) {
		t.Errorf("Expected XML string %s, got %s", expectedXMLString, string(xmlData))
	}
}
