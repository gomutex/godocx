package elements

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestTableStyle_MarshalXML(t *testing.T) {
	style := NewTableStyle("MyTableStyle")

	expected := `<w:tblStyle w:val="MyTableStyle"></w:tblStyle>`

	var builder strings.Builder
	encoder := xml.NewEncoder(&builder)
	err := encoder.Encode(style)
	if err != nil {
		t.Fatalf("Error encoding TableStyle: %v", err)
	}

	result := builder.String()
	if result != expected {
		t.Errorf("Unexpected XML. Expected: %s, Got: %s", expected, result)
	}
}

func TestTableStyle_UnmarshalXML(t *testing.T) {
	xmlData := `<w:tblStyle w:val="AnotherTableStyle"></w:tblStyle>`

	expected := "AnotherTableStyle"

	decoder := xml.NewDecoder(strings.NewReader(xmlData))
	var style TableStyle
	err := decoder.Decode(&style)
	if err != nil {
		t.Fatalf("Error decoding TableStyle: %v", err)
	}

	if style.Val != expected {
		t.Errorf("Unexpected value. Expected: %s, Got: %s", expected, style.Val)
	}
}
