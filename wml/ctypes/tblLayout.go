package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// TableLayout represents the layout of a table in a document.
type TableLayout struct {
	LayoutType *stypes.TableLayout `xml:"type,attr,omitempty"`
}

func DefaultTableLayout() *TableLayout {
	return &TableLayout{}
}

// NewTableLayout creates a new TableLayout instance.
func NewTableLayout(t stypes.TableLayout) *TableLayout {
	return &TableLayout{LayoutType: &t}
}

// MarshalXML implements the xml.Marshaler interface for TableLayout.
func (t TableLayout) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:tblLayout"
	if t.LayoutType != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:type"}, Value: string(*t.LayoutType)})
	}
	return e.EncodeElement("", start)
}
