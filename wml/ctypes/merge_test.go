package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestMerge(t *testing.T) {
	testMerge := NewMerge(stypes.MergeCellContinue)

	xmlData, err := xml.Marshal(testMerge)
	if err != nil {
		t.Fatalf("Error marshaling Merge to XML: %v", err)
	}

	var unmarshaledMerge Merge
	err = xml.Unmarshal(xmlData, &unmarshaledMerge)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to Merge: %v", err)
	}

	if testMerge.Val != unmarshaledMerge.Val {
		t.Errorf("Expected merge value %s, got %s", testMerge.Val, unmarshaledMerge.Val)
	}

	expectedXMLString := `<Merge w:val="continue"></Merge>`
	if !strings.Contains(string(xmlData), expectedXMLString) {
		t.Errorf("Expected XML string %s, got %s", expectedXMLString, string(xmlData))
	}
}
