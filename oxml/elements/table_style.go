package elements

import (
	"encoding/xml"
)

// TableStyle represents the style of a table in a document.
type TableStyle struct {
	Val string
}

// NewTableStyle creates a new TableStyle instance.
func NewTableStyle(val string) *TableStyle {
	return &TableStyle{Val: val}
}

// MarshalXML implements the xml.Marshaler interface for TableStyle.
func (t *TableStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:tblStyle"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: t.Val})
	return e.EncodeElement("", start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for TableStyle.
func (t *TableStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "val" {
			t.Val = a.Value
			break
		}
	}
	return d.Skip() // Skipping the inner content
}
