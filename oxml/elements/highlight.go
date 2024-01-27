package elements

import "encoding/xml"

// Highlight represents the highlighting of a text or element.
type Highlight struct {
	Value ColorIndex
}

// NewHighlight creates a new Highlight.
func NewHighlight(value ColorIndex) *Highlight {
	return &Highlight{Value: value}
}

// MarshalXML marshals Highlight to XML.
func (h *Highlight) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:highlight"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(h.Value)})
	return e.EncodeElement("", start)
}

// UnmarshalXML unmarshals XML to Highlight.
func (h *Highlight) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var attr string
	for _, a := range start.Attr {
		if a.Name.Local == "val" {
			attr = a.Value
			break
		}
	}

	h.Value = ColorIndex(attr)
	return d.Skip() // Skipping the inner content
}
