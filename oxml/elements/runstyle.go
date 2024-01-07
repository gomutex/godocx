package elements

import (
	"encoding/xml"
)

// RunStyle represents the style of a run within a document.
type RunStyle struct {
	Val string
}

// NewRunStyle creates a new RunStyle.
func NewRunStyle(val string) *RunStyle {
	return &RunStyle{Val: val}
}

// DefaultRunStyle creates the default RunStyle with the value "Normal".
func DefaultRunStyle() *RunStyle {
	return &RunStyle{Val: "Normal"}
}

// MarshalXML marshals RunStyle to XML.
func (r *RunStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:rStyle"
	return e.EncodeElement(r.Val, start)
}

// UnmarshalXML unmarshals XML to RunStyle.
func (r *RunStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return d.DecodeElement(&r.Val, &start)
}
