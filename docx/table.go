package docx

import (
	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/stypes"
)

type Table struct {
	// Reverse inheriting the Rootdoc into Paragrah to access other elements
	Root *RootDoc

	// Paragraph Complex Type
	CT ctypes.Table
}

func NewTable(root *RootDoc) *Table {
	return &Table{
		Root: root,
	}
}

// AddTable adds a new table to the root document.
//
// It creates and initializes a new table, appends it to the root document's body, and returns a pointer to the created table.
// The table is initially empty, with no rows or cells. To add content to the table, use the provided methods on the returned
// table instance.
//
// Example usage:
//   docx := godocx.NewDocument()
//   table := docx.AddTable()
//   table.Style("LightList-Accent2")
//
//   // Add rows and cells to the table
//   row := table.AddRow()
//   cell := row.AddCell()
//   cell.AddParagraph("Hello, World!")
//
// Parameters:
//   - rd: A pointer to the RootDoc instance.
//
// Returns:
//   - *elements.Table: A pointer to the newly added table.

func (rd *RootDoc) AddTable() *Table {
	tbl := Table{}

	rd.Document.Body.Children = append(rd.Document.Body.Children, DocumentChild{
		Table: &tbl,
	})

	return &tbl
}

func (t *Table) AddRow() *ctypes.Row {
	row := ctypes.DefaultRow()
	t.CT.RowContents = append(t.CT.RowContents, ctypes.RowContent{
		Row: row,
	})
	return row
}

func (t *Table) ensureProp() {
}

func (t *Table) Indent(indent int) {
	t.CT.TableProp.Indent = ctypes.NewTableWidth(indent, stypes.TableWidthAuto)
}

func (t *Table) Style(value string) {
	t.CT.TableProp.Style = ctypes.NewTableStyle(value)
}
