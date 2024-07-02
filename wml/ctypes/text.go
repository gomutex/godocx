package ctypes

import (
	"bytes"
	"encoding/xml"
	"strings"
)

type Text struct {
	Text  string
	Space *string
}

const (
	TextSpaceDefault  = "default"
	TextSpacePreserve = "preserve"
)

func NewText() *Text {
	return &Text{}
}

func TextFromString(text string) *Text {
	t := &Text{Text: text}
	if strings.TrimSpace(text) != text {
		xmlSpace := "preserve"
		t.Space = &xmlSpace
	}
	return t
}

func (t Text) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {

	if t.Space != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xml:space"}, Value: *t.Space})
	}

	if err = e.EncodeElement(t.Text, start); err != nil {
		return err
	}

	return nil
}

func (t *Text) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var buf bytes.Buffer

	for _, attr := range start.Attr {
		if attr.Name.Local == "space" {
			t.Space = &attr.Value
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
				t.Text = buf.String()
				return nil
			}
		}
	}
}
