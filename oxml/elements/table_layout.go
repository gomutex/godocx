package elements

import (
	"encoding/xml"
)

type LayoutType string

const (
	LayoutTypeFixed   LayoutType = "fixed"
	LayoutTypeAutoFit LayoutType = "autofit"
)

// TableLayout represents the layout of a table in a document.
type TableLayout struct {
	LayoutType LayoutType
}

func DefaultTableLayout() *TableLayout {
	return &TableLayout{}
}

// NewTableLayout creates a new TableLayout instance.
func NewTableLayout(t LayoutType) *TableLayout {
	return &TableLayout{LayoutType: t}
}

// MarshalXML implements the xml.Marshaler interface for TableLayout.
func (t *TableLayout) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:tblLayout"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:type"}, Value: string(t.LayoutType)})
	return e.EncodeElement("", start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for TableLayout.
func (t *TableLayout) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "type" {
			t.LayoutType = LayoutType(a.Value)
			break
		}
	}
	return d.Skip() // Skipping the inner content
}
