package wml

import (
	"encoding/xml"
)

// Strike represents the strike style of text or element.
type Strike struct {
	Val bool
}

// NewStrike creates a new Strike instance.
func NewStrike(value bool) *Strike {
	return &Strike{Val: value}
}

// Disable disables the strike style.
func (b *Strike) Disable() *Strike {
	b.Val = false
	return b
}

// MarshalXML implements the xml.Marshaler interface for the Strike type.
// It encodes the Strike instance into XML using the "w:b" element with a "w:val" attribute.
func (b *Strike) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:strike"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: map[bool]string{true: "true", false: "false"}[b.Val]})
	return e.EncodeElement("", start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for the Strike type.
// It decodes the XML representation of Strike, extracting the value from the "w:val" attribute.
// The inner content of the XML element is skipped.
func (b *Strike) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
