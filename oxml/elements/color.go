package elements

import (
	"encoding/xml"
)

// Color represents the color of a text or element.
type Color struct {
	Value string
}

// NewColor creates a new Color.
func NewColor(value string) *Color {
	return &Color{Value: value}
}

func (c *Color) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:color"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: c.Value})
	return e.EncodeElement("", start)
}

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
