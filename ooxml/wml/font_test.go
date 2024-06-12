package wml

import (
	"encoding/xml"
	"testing"
)

func TestFont_UnmarshalXML(t *testing.T) {
	xmlString := `<w:font w:name="Arial"  xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main" xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
	<w:charset w:val="00" />
	<w:family w:val="swiss" />
	<w:pitch w:val="variable" />
  </w:font>`

	var parsedFont Font
	err := xml.Unmarshal([]byte(xmlString), &parsedFont)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to Font: %v", err)
	}

	expectedFont := Font{
		Name:    "Arial",
		Charset: Charset{Value: "00"},
		Family:  Family{Value: "swiss"},
		Pitch:   Pitch{Value: "variable"},
	}

	// Compare the parsedFont with the expectedFont
	if parsedFont.Name != expectedFont.Name ||
		parsedFont.Charset.Value != expectedFont.Charset.Value ||
		parsedFont.Family.Value != expectedFont.Family.Value ||
		parsedFont.Pitch.Value != expectedFont.Pitch.Value {
		t.Errorf("Expected Font %v, got %v", expectedFont, parsedFont)
	}
}
