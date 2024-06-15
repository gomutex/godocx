package formatting

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestUnderline(t *testing.T) {
	testUnderline := NewUnderline(UnderlineSingle)

	xmlData, err := xml.Marshal(testUnderline)
	if err != nil {
		t.Fatalf("Error marshaling Underline to XML: %v", err)
	}

	var unmarshaledUnderline Underline
	err = xml.Unmarshal(xmlData, &unmarshaledUnderline)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to Underline: %v", err)
	}

	if testUnderline.Val != unmarshaledUnderline.Val {
		t.Errorf("Expected underline value %s, got %s", testUnderline.Val, unmarshaledUnderline.Val)
	}

	expectedXMLString := `<w:u w:val="single"></w:u>`
	if !strings.Contains(string(xmlData), expectedXMLString) {
		t.Errorf("Expected XML string %s, got %s", expectedXMLString, string(xmlData))
	}
}
