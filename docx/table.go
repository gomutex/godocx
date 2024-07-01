package docx

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/stypes"
)

type Table struct {
	// Reverse inheriting the Rootdoc into paragraph to access other elements
	root *RootDoc

	// Table Complex Type
	ct ctypes.Table
}

func (t *Table) unmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return t.ct.UnmarshalXML(d, start)
}

func NewTable(root *RootDoc) *Table {
	return &Table{
		root: root,
	}
}

// AddTable adds a new table to the root document.
//
// It creates and initializes a new table, appends it to the root document's body, and returns a pointer to the created table.
// The table is initially empty, with no rows or cells. To add content to the table, use the provided methods on the returned
// table instance.
//
// Example usage:
//   document := godocx.NewDocument()
//   table := document.AddTable()
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
	tbl := Table{
		root: rd,
		ct:   *ctypes.DefaultTable(),
	}

	rd.Document.Body.Children = append(rd.Document.Body.Children, DocumentChild{
		Table: &tbl,
	})

	return &tbl
}

// AddRow adds a new row to the table.
//
// It creates a new row and appends it to the table's row contents. Use this method to construct the structure
// of the table by sequentially adding rows and cells.
//
// Returns:
//   - *ctypes.Row: A pointer to the newly added row.

func (t *Table) AddRow() *Row {
	row := Row{
		root: t.root,
		ct:   *ctypes.DefaultRow(),
	}

	t.ct.RowContents = append(t.ct.RowContents, ctypes.RowContent{
		Row: &row.ct,
	})

	return &row
}

func (t *Table) ensureProp() {
}

// Indent sets the indent width for the table.
//
// Parameters:
//   - indent: An integer specifying the indent width
func (t *Table) Indent(indent int) {
	t.ct.TableProp.Indent = ctypes.NewTableWidth(indent, stypes.TableWidthAuto)
}

// Style sets the style for the table.
//
// TableStyle represents the style of a table in a document.
// This is applicable when creating a new document. When using this style in a new document, you need to ensure
// that the specified style ID exists in your document's style base or is manually created through the library.
//
// Some examples of predefined style IDs in the docx template that can be used are:
//
//   - "LightShading"
//   - "LightShading-Accent1"
//   - "LightShading-Accent2"
//   - "LightShading-Accent3"
//   - "LightShading-Accent4"
//   - "LightShading-Accent5"
//   - "LightShading-Accent6"
//   - "LightList"
//   - "LightList-Accent1"..."LightList-Accent6"
//   - "LightGrid"
//   - "LightGrid-Accent1"..."LightGrid-Accent6"
//   - "MediumShading"
//   - "MediumShading-Accent1"..."MediumShading-Accent6"
//   - "MediumShading2"
//   - "MediumShading2-Accent1"..."MediumShading2-Accent6"
//   - "MediumList1"
//   - "MediumList1-Accent1"..."MediumList1-Accent6"
//   - "MediumList2"
//   - "MediumList2-Accent1"..."MediumList2-Accent6"
//   - "TableGrid"
//   - "MediumGrid1"
//   - "MediumGrid1-Accent1"..."MediumGrid1-Accent6"
//   - "MediumGrid2"
//   - "MediumGrid2-Accent1"..."MediumGrid2-Accent6"
//   - "MediumGrid3"
//   - "MediumGrid3-Accent1"..."MediumGrid3-Accent6"
//   - "DarkList"
//   - "DarkList-Accent1"..."DarkList-Accent6"
//   - "ColorfulShading"
//   - "ColorfulShading-Accent1"..."ColorfulShading-Accent6"
//   - "ColorfulList"
//   - "ColorfulList-Accent1"..."ColorfulList-Accent6"
//   - "ColorfulGrid"
//   - "ColorfulGrid-Accent1"..."ColorfulGrid-Accent6"
//
// Parameters:
//   - value: A string representing the style value. It should match a valid table style defined in the WordprocessingML specification.
func (t *Table) Style(value string) {
	t.ct.TableProp.Style = ctypes.NewCTString(value)
}

type Row struct {
	// Reverse inheriting the Rootdoc into paragraph to access other elements
	root *RootDoc

	// Row Complex Type
	ct ctypes.Row
}

// func (r *Row) unmarshalXML(d *xml.Decoder, start xml.StartElement) error {
// 	return r.ct.UnmarshalXML(d, start)
// }

type RunPropOpts func(*Run)

func WithFontSize(size uint64) RunPropOpts {
	return func(r *Run) {
		r.getProp().Size = ctypes.NewFontSize(size * 2)
	}
}

func (r *Row) AddTextCell(text string, opts ...RunPropOpts) *Cell {
	// Wrapper cell
	cell := Cell{
		root: r.root,
		ct:   *ctypes.DefaultCell(),
	}

	// add paragraph with text & get Run obj
	p := newParagraph(r.root)
	run := p.AddText(text)

	// Table cell block content
	tblContent := ctypes.TCBlockContent{
		Paragraph: &p.ct,
	}

	cell.ct.Contents = append(cell.ct.Contents, tblContent)

	r.ct.Contents = append(r.ct.Contents, ctypes.TRCellContent{
		Cell: &cell.ct,
	})

	for _, opt := range opts {
		opt(run)
	}

	return &cell
}

func (r *Row) AddCell() *Cell {
	cell := Cell{
		root: r.root,
		ct:   *ctypes.DefaultCell(),
	}

	r.ct.Contents = append(r.ct.Contents, ctypes.TRCellContent{
		Cell: &cell.ct,
	})

	return &cell
}

type Cell struct {
	// Reverse inheriting the Rootdoc into paragraph to access other elements
	root *RootDoc

	// Cell Complex Type
	ct ctypes.Cell
}

func (c *Cell) AddParagraph(text string) *Paragraph {
	p := newParagraph(c.root, paraWithText(text))
	tblContent := ctypes.TCBlockContent{
		Paragraph: &p.ct,
	}

	c.ct.Contents = append(c.ct.Contents, tblContent)

	return p
}
