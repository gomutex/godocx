package wml

import (
	"encoding/xml"
)

// Bold represents the bold style of text or element.
type Bold struct {
	Val bool
}

// NewBold creates a new Bold instance.
func NewBold(value bool) *Bold {
	return &Bold{Val: value}
}

// Disable disables the bold style.
func (b *Bold) Disable() *Bold {
	b.Val = false
	return b
}

// MarshalXML implements the xml.Marshaler interface for the Bold type.
// It encodes the Bold instance into XML using the "w:b" element with a "w:val" attribute.
func (b *Bold) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:b"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: map[bool]string{true: "true", false: "false"}[b.Val]})
	return e.EncodeElement("", start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for the Bold type.
// It decodes the XML representation of Bold, extracting the value from the "w:val" attribute.
// The inner content of the XML element is skipped.
func (b *Bold) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var attr string
	for _, a := range start.Attr {
		if a.Name.Local == "val" {
			attr = a.Value
			break
		}
	}

	b.Val = attr == "true"
	return d.Skip() // Skipping the inner content
}
