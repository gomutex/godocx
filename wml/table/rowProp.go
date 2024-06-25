package table

import "encoding/xml"

// RowProperty represents the properties of a table row.
type RowProperty struct {
	// GridAfter   *uint32
	// WidthAfter  *float32
	// GridBefore  *uint32
	// WidthBefore *float32
	// RowHeight   *RowHeight
	// Del         *Delete // TODO: Implement Delete
	// Ins         *Insert // TODO: Implement Insert
}

type RowHeight struct {
	Val        *float32
	HeightRule *string
}

// NewRowProperty creates a new RowProperty instance.
func DefaultRowProperty() *RowProperty {
	return &RowProperty{}
}

func (r *RowProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:trPr"

	// TODO:
	return nil
}
