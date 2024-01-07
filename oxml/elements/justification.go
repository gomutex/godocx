package elements

import "encoding/xml"

// Justification represents the justification of a paragraph.
type Justification struct {
	Val string
}

// NewJustification creates a new Justification.
func NewJustification(val string) *Justification {
	return &Justification{Val: val}
}

// DefaultJustification creates the default Justification with the value "centerGroup".
func DefaultJustification() *Justification {
	return &Justification{Val: "centerGroup"}
}

func (j *Justification) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:jc"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: j.Val})
	return e.EncodeElement("", start)
}

func (j *Justification) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var attr string
	for _, a := range start.Attr {
		if a.Name.Local == "val" {
			attr = a.Value
			break
		}
	}

	j.Val = attr
	return d.Skip() // Skipping the inner content
}
