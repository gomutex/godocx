package ooxml

import "github.com/gomutex/godocx/ooxml/wml"

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

func (rd *RootDoc) AddTable() *wml.Table {
	tbl := wml.DefaultTable()
	tbl.Rows = []wml.TableChild{}
	tbl.Grid = wml.DefaultTableGrid()

	rd.Document.Body.Children = append(rd.Document.Body.Children, DocumentChild{
		Table: tbl,
	})

	return tbl
}
