package wml

import (
	"encoding/xml"
)

// Bold represents the bold style of text or element.
type Bold struct {
	Val      bool
	ValExist bool
}

// NewBold creates a new Bold instance.
func NewBold(value bool) *Bold {
	return &Bold{Val: value, ValExist: true}
}

// Disable disables the bold style.
func (b *Bold) Disable() *Bold {
	b.Val = false
	b.ValExist = true
	return b
}

// MarshalXML implements the xml.Marshaler interface for the Bold type.
// It encodes the Bold instance into XML using the "w:b" element with a "w:val" attribute.
func (b *Bold) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:b"
	if b.ValExist { // Add val attribute only if the val exists
		if b.Val {
			start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: "true"})
		} else {
			start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: "false"})
		}
	}
	return e.EncodeElement("", start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for the Bold type.
// It decodes the XML representation of Bold, extracting the value from the "w:val" attribute.
// The inner content of the XML element is skipped.
func (b *Bold) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "val" {
			// If value is "true", then set it to true
			b.Val = a.Value == "true"
			b.ValExist = true
			break
		}
	}

	return d.Skip() // Skipping the inner content
}
