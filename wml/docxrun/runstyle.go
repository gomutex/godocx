package docxrun

import (
	"encoding/xml"
)

// RunStyle represents the style of a run within a document.
type RunStyle struct {
	Value string
}

// NewRunStyle creates a new RunStyle.
func NewRunStyle(value string) *RunStyle {
	return &RunStyle{Value: value}
}

// DefaultRunStyle creates the default RunStyle with the value "Normal".
func DefaultRunStyle() *RunStyle {
	return &RunStyle{Value: "Normal"}
}

// MarshalXML marshals RunStyle to XML.
func (r *RunStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:rStyle"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: r.Value})

	return e.EncodeElement("", start)
}

// UnmarshalXML unmarshals XML to RunStyle.
func (r *RunStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var attr string
	for _, a := range start.Attr {
		if a.Name.Local == "val" {
			attr = a.Value
			break
		}
	}

	r.Value = attr
	return d.Skip() // Skipping the inner content
}
