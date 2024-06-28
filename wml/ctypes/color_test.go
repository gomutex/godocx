package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestColor(t *testing.T) {
	testColor := NewColor("FF0000")

	xmlData, err := xml.Marshal(testColor)
	if err != nil {
		t.Fatalf("Error marshaling Color to XML: %v", err)
	}

	var unmarshaledColor Color
	err = xml.Unmarshal(xmlData, &unmarshaledColor)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to Color: %v", err)
	}

	if testColor.Val != unmarshaledColor.Val {
		t.Errorf("Expected color value %s, got %s", testColor.Val, unmarshaledColor.Val)
	}

	expectedXMLString := `<w:color w:val="FF0000"></w:color>`
	if !strings.Contains(string(xmlData), expectedXMLString) {
		t.Errorf("Expected XML string %s, got %s", expectedXMLString, string(xmlData))
	}
}
