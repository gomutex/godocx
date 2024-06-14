package txt

import (
	"encoding/xml"
	"errors"
)

type Break struct {
	BreakType BreakType
}

type BreakType string

const (
	BreakTypePage         BreakType = "page"
	BreakTypeColumn       BreakType = "column"
	BreakTypeTextWrapping BreakType = "textWrapping"
	BreakTypeUnsupported  BreakType = "unsupported"
)

// NewBreak creates a new Break element with the given break type.
func NewBreak(breakType BreakType) *Break {
	return &Break{
		BreakType: breakType,
	}
}

// MarshalXML implements the xml.Marshaler interface.
func (b *Break) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "br"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "type"}, Value: string(b.BreakType)})
	return e.EncodeElement("", start)
}

// UnmarshalXML implements the xml.Unmarshaler interface.
func (b *Break) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Local != "br" {
		return errors.New("unexpected element name")
	}

	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			b.BreakType = BreakType(attr.Value)
			break
		}
	}

	return d.Skip() // Skipping the inner content
}
