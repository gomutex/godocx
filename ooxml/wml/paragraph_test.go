package wml

import (
	"bytes"
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/gomutex/godocx/common/constants"
)

func TestParagraphXML(t *testing.T) {
	// Create a sample paragraph
	para := NewParagraph()
	// para.Style("Heading1")
	// para.Numbering(1, 0)
	para.AddText("This is a sample paragraph.")

	// Marshal the paragraph to XML
	var buf bytes.Buffer

	encoder := xml.NewEncoder(&buf)
	start := xml.StartElement{Name: xml.Name{Local: "fake"}}
	if err := para.MarshalXML(encoder, start); err != nil {
		t.Errorf("Error during MarshalXML: %v", err)
	}

	err := encoder.Flush()
	if err != nil {
		t.Errorf("Error flushing encoder: %v", err)
	}

	// Unmarshal the XML back to a paragraph
	var paraUnmarshaled Paragraph
	ns := map[string]string{
		"w": constants.WMLNamespace,
	}
	decoder := xml.NewDecoder(&buf)
	decoder.DefaultSpace = constants.WMLNamespace
	decoder.Entity = ns

	if err := decoder.Decode(&paraUnmarshaled); err != nil {
		t.Errorf("Error unmarshaling XML to paragraph: %v", err)
		return
	}

	if !reflect.DeepEqual(*para, paraUnmarshaled) {
		t.Errorf("Original and unmarshaled paragraphs are not equal.")
	}
}
