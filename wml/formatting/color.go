package formatting

import (
	"encoding/xml"
)

// Color represents the color of a text or element.
type Color struct {
	Val string
}

// NewColor creates a new Color instance with the specified color value.
func NewColor(value string) *Color {
	return &Color{Val: value}
}

// MarshalXML implements the xml.Marshaler interface for the Color type.
//
// It encodes the Color instance into XML using the "w:color" element with a "w:val" attribute.
func (c Color) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:color"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: c.Val})
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

	c.Val = attr
	return d.Skip() // Skipping the inner content
}

// ColorIndex represents different color indexes.

const (
	ColorIndexAuto        = "default"
	ColorIndexBlack       = "black"
	ColorIndexBlue        = "blue"
	ColorIndexBrightGreen = "green"
	ColorIndexDarkBlue    = "darkBlue"
	ColorIndexDarkRed     = "darkRed"
	ColorIndexDarkYellow  = "darkYellow"
	ColorIndexGray25      = "lightGray"
	ColorIndexGray50      = "darkGray"
	ColorIndexGreen       = "darkGreen"
	ColorIndexMagenta     = "magenta"
	ColorIndexRed         = "red"
	ColorIndexDarkCyan    = "darkCyan"
	ColorIndexCyan        = "cyan"
	ColorIndexDarkMagenta = "darkMagenta"
	ColorIndexWhite       = "white"
	ColorIndexYellow      = "yellow"
)
