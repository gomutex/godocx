package elements

import (
	"encoding/xml"
)

// Color represents the color of a text or element.
type Color struct {
	Value string
}

// NewColor creates a new Color instance with the specified color value.
func NewColor(value string) *Color {
	return &Color{Value: value}
}

// MarshalXML implements the xml.Marshaler interface for the Color type.
//
// It encodes the Color instance into XML using the "w:color" element with a "w:val" attribute.
func (c *Color) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:color"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: c.Value})
	return e.EncodeElement("", start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for the Color type.
//
// It decodes the XML representation of Color, extracting the value from the "val" attribute.
// The inner content of the XML element is skipped.
func (c *Color) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var attr string
	for _, a := range start.Attr {
		if a.Name.Local == "val" {
			attr = a.Value
			break
		}
	}

	c.Value = attr
	return d.Skip() // Skipping the inner content
}
