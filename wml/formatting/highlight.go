package formatting

import (
	"encoding/xml"
)

// Highlight represents the highlighting of a text or element.
type Highlight struct {
	Val string
}

// NewHighlight creates a new Highlight.
func NewHighlight(value string) *Highlight {
	return &Highlight{Val: value}
}

// MarshalXML marshals Highlight to XML.
func (h Highlight) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:highlight"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: h.Val})
	return e.EncodeElement("", start)
}

// UnmarshalXML unmarshals XML to Highlight.
func (h *Highlight) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var valStr string
	for _, a := range start.Attr {
		if a.Name.Local == "val" {
			valStr = a.Value
			break
		}
	}

	h.Val = valStr
	return d.Skip() // Skipping the inner content
}
