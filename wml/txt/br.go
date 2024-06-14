package txt

import (
	"encoding/xml"
)

type Break struct {
	BreakType *BreakType
	Clear     *BreakClear
}

type BreakType string

const (
	BreakTypePage         BreakType = "page"         // Page Break
	BreakTypeColumn       BreakType = "column"       // Column Break
	BreakTypeTextWrapping BreakType = "textWrapping" // Line Break
	BreakTypeInvalid      BreakType = ""
)

type BreakClear string

const (
	BreakClearNone    BreakClear = "none"  // Restart On Next Line
	BreakClearLeft    BreakClear = "left"  // Restart In Next Text Region When In Leftmost Position
	BreakClearRight   BreakClear = "right" // Restart In Next Text Region When In Rightmost Position
	BreakClearAll     BreakClear = "all"   // Restart On Next Full Line
	BreakClearInvalid BreakClear = ""
)

// NewBreak creates a new Break element with the given break type.
func NewBreak(breakType BreakType) *Break {
	return &Break{
		BreakType: &breakType,
	}
}

// MarshalXML implements the xml.Marshaler interface.
func (b *Break) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:br"

	if b.BreakType != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "type"}, Value: string(*b.BreakType)})
	}

	if b.Clear != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "clear"}, Value: string(*b.Clear)})
	}

	return e.EncodeElement("", start)
}

// UnmarshalXML implements the xml.Unmarshaler interface.
func (b *Break) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "type":
			bt := BreakType(attr.Value)
			b.BreakType = &bt
		case "clear":
			bc := BreakClear(attr.Value)
			b.Clear = &bc
		}
	}

	return d.Skip() // Skipping the inner content
}
