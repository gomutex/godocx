package formatting

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestStrike(t *testing.T) {
	testStrike := NewStrike(true)

	xmlData, err := xml.Marshal(testStrike)
	if err != nil {
		t.Fatalf("Error marshaling Strike to XML: %v", err)
	}

	var unmarshaledStrike Strike
	err = xml.Unmarshal(xmlData, &unmarshaledStrike)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to Strike: %v", err)
	}

	if testStrike.Val != unmarshaledStrike.Val {
		t.Errorf("Expected strike value %t, got %t", testStrike.Val, unmarshaledStrike.Val)
	}

	expectedXMLString := `<w:strike w:val="true"></w:strike>`
	if !strings.Contains(string(xmlData), expectedXMLString) {
		t.Errorf("Expected XML string %s, got %s", expectedXMLString, string(xmlData))
	}
}
