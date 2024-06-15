package table

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestTableLayout_MarshalXML(t *testing.T) {
	layout := &TableLayout{LayoutType: LayoutTypeFixed}

	expected := `<w:tblLayout w:type="fixed"></w:tblLayout>`

	var builder strings.Builder
	encoder := xml.NewEncoder(&builder)
	err := encoder.Encode(layout)
	if err != nil {
		t.Fatalf("Error encoding TableLayout: %v", err)
	}

	result := builder.String()
	if result != expected {
		t.Errorf("Unexpected XML. Expected: %s, Got: %s", expected, result)
	}
}

func TestTableLayout_UnmarshalXML(t *testing.T) {
	xmlData := `<w:tblLayout w:type="autofit"></w:tblLayout>`

	expected := LayoutTypeAutoFit

	decoder := xml.NewDecoder(strings.NewReader(xmlData))
	var layout TableLayout
	err := decoder.Decode(&layout)
	if err != nil {
		t.Fatalf("Error decoding TableLayout: %v", err)
	}

	if layout.LayoutType != expected {
		t.Errorf("Unexpected layout type. Expected: %s, Got: %s", expected, layout.LayoutType)
	}
}
