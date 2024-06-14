package txt

import "encoding/xml"

// ParagraphStyle represents the name of a paragraph style associated with a numbering level in a document.
// When a paragraph style includes a numbering definition, any numbering level defined by the numPr element (§17.3.1.19)
// shall be ignored. Instead, this element specifies the numbering level associated with that paragraph style.
type ParagraphStyle struct {
	Value string
}

// NewParagraphStyle creates a new ParagraphStyle.
func NewParagraphStyle(val string) *ParagraphStyle {
	return &ParagraphStyle{Value: val}
}

// DefaultParagraphStyle creates the default ParagraphStyle with the value "Normal".
func DefaultParagraphStyle() *ParagraphStyle {
	return &ParagraphStyle{Value: "Normal"}
}

func (p *ParagraphStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:pStyle"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: p.Value})
	return e.EncodeElement("", start)
}

func (p *ParagraphStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var attr string
	for _, a := range start.Attr {
		if a.Name.Local == "val" {
			attr = a.Value
			break
		}
	}

	p.Value = attr
	return d.Skip() // Skipping the inner content
}
