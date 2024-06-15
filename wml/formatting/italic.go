package formatting

import (
	"encoding/xml"
)

// Italic represents the italic style of text or element.
type Italic struct {
	Val bool
}

// NewItalic creates a new Italic instance.
func NewItalic(value bool) *Italic {
	return &Italic{Val: value}
}

// Disable disables the italic style.
func (i *Italic) Disable() *Italic {
	i.Val = false
	return i
}

// MarshalXML implements the xml.Marshaler interface for the Italic type.
// It encodes the Italic instance into XML using the "w:i" element with a "w:val" attribute.
func (i *Italic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:i"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: map[bool]string{true: "true", false: "false"}[i.Val]})
	return e.EncodeElement("", start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for the Italic type.
// It decodes the XML representation of Italic, extracting the value from the "w:val" attribute.
// The inner content of the XML element is skipped.
func (i *Italic) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var attr string
	for _, a := range start.Attr {
		if a.Name.Local == "val" {
			attr = a.Value
			break
		}
	}

	i.Val = attr == "true"
	return d.Skip() // Skipping the inner content
}
