package table

// TableRowProperty represents the properties of a table row.
type TableRowProperty struct {
	GridAfter   *uint32
	WidthAfter  *float32
	GridBefore  *uint32
	WidthBefore *float32
	RowHeight   *RowHeight
	// Del         *Delete // TODO: Implement Delete
	// Ins         *Insert // TODO: Implement Insert
}

type RowHeight struct {
	Val        *float32
	HeightRule *string
}

// NewTableRowProperty creates a new TableRowProperty instance.
func DefaultTableRowProperty() *TableRowProperty {
	return &TableRowProperty{}
}
