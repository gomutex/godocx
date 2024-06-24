package runcontent

import (
	"bytes"
	"encoding/xml"
)

type Text struct {
	text  string
	space *string
}

const (
	TextSpaceDefault  = "default"
	TextSpacePreserve = "preserve"
)

func NewText() *Text {
	return &Text{}
}

func TextFromString(text string) *Text {
	return &Text{text: text}
}

func (t *Text) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {

	elem := xml.StartElement{Name: xml.Name{Local: "w:t"}}

	if t.space != nil {
		elem.Attr = append(elem.Attr, xml.Attr{Name: xml.Name{Local: "xml:space"}, Value: *t.space})
	}

	if err = e.EncodeElement(t.text, elem); err != nil {
		return err
	}

	return nil
}

func (t *Text) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var buf bytes.Buffer

	for _, attr := range start.Attr {
		if attr.Name.Local == "space" {
			t.space = &attr.Value
			break
		}
	}

	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch tokenElem := token.(type) {
		case xml.CharData:
			buf.Write([]byte(tokenElem))
		case xml.EndElement:
			if tokenElem == start.End() {
				t.text = buf.String()
				return nil
			}
		}
	}
}
