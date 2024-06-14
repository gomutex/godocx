package formatting

import (
	"encoding/xml"
)

// UnderlineStyle represents different styles for underline.
type UnderlineStyle string

const (
	UnderlineNone            UnderlineStyle = "none"
	UnderlineSingle          UnderlineStyle = "single"
	UnderlineWords           UnderlineStyle = "words"
	UnderlineDouble          UnderlineStyle = "double"
	UnderlineDotted          UnderlineStyle = "dotted"
	UnderlineThick           UnderlineStyle = "thick"
	UnderlineDash            UnderlineStyle = "dash"
	UnderlineDotDash         UnderlineStyle = "dotDash"
	UnderlineDotDotDash      UnderlineStyle = "dotDotDash"
	UnderlineWavy            UnderlineStyle = "wavy"
	UnderlineDottedHeavy     UnderlineStyle = "dottedHeavy"
	UnderlineDashHeavy       UnderlineStyle = "dashHeavy"
	UnderlineDotDashHeavy    UnderlineStyle = "dotDashHeavy"
	UnderlineDotDotDashHeavy UnderlineStyle = "dotDotDashHeavy"
	UnderlineWavyHeavy       UnderlineStyle = "wavyHeavy"
	UnderlineDashLong        UnderlineStyle = "dashLong"
	UnderlineWavyDouble      UnderlineStyle = "wavyDouble"
	UnderlineDashLongHeavy   UnderlineStyle = "dashLongHeavy"
)

// Underline represents the underline style of text or element.
type Underline struct {
	Val UnderlineStyle
}

// NewUnderline creates a new Underline instance.
func NewUnderline(val UnderlineStyle) *Underline {
	return &Underline{Val: val}
}

// MarshalXML implements the xml.Marshaler interface for the Underline type.
// It encodes the Underline instance into XML using the "w:u" element with a "w:val" attribute.
func (u *Underline) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:u"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(u.Val)})
	return e.EncodeElement("", start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for the Underline type.
// It decodes the XML representation of Underline, extracting the value from the "w:val" attribute.
// The inner content of the XML element is skipped.
func (u *Underline) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var attr string
	for _, a := range start.Attr {
		if a.Name.Local == "val" {
			attr = a.Value
			break
		}
	}

	u.Val = UnderlineStyle(attr)
	return d.Skip() // Skipping the inner content
}
