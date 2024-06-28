package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestBreakMarshaling(t *testing.T) {
	breakType := stypes.BreakTypePage
	br := NewBreak(breakType)

	expectedXML := `<w:br w:type="page"></w:br>`

	xmlData, err := xml.Marshal(br)
	if err != nil {
		t.Fatalf("Error marshaling Break to XML: %v", err)
	}

	if strings.TrimSpace(string(xmlData)) != expectedXML {
		t.Errorf("Unexpected XML output. Expected:\n%s\nGot:\n%s", expectedXML, string(xmlData))
	}

	var unmarshalledBreak Break
	err = xml.Unmarshal(xmlData, &unmarshalledBreak)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to Break: %v", err)
	}

	if *unmarshalledBreak.BreakType != *br.BreakType {
		t.Errorf("Expected BreakType %s, got %s", *br.BreakType, *unmarshalledBreak.BreakType)
	}
}
