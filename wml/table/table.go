package table

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/stypes"
)

// Table
type Table struct {
	//1.Choice: RangeMarkupElements
	RngMarkupElems []RngMarkupElem

	//2. Table Properties
	Property Property `xml:"tblPr,omitempty"`

	//3. Table Grid
	Grid Grid `xml:"tblGrid,omitempty"`

	//4.1 Choice:
	RowContents []RowContent

	//4.2 TODO: Remaining choices
}

func DefaultTable() *Table {
	return &Table{}
}

func (t *Table) AddRow() *Row {
	row := DefaultRow()
	t.RowContents = append(t.RowContents, RowContent{
		Row: row,
	})
	return row
}

func (t *Table) Indent(indent int) {
	t.Property = *DefaultProperty()
	t.Property.Indent = ctypes.NewTableWidth(indent, stypes.TableWidthAuto)
}

func (t *Table) Style(value string) {
	t.Property = *DefaultProperty()
	t.Property.Style = NewTableStyle(value)
}

func (t *Table) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:tbl"

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	//1.Choice: RangeMarkupElements
	for _, rme := range t.RngMarkupElems {
		err := rme.MarshalXML(e, xml.StartElement{})
		if err != nil {
			return err
		}
	}

	//2. Table Properties
	if err = t.Property.MarshalXML(e, xml.StartElement{}); err != nil {
		return err
	}

	//3. Table Grid
	if err = t.Grid.MarshalXML(e, xml.StartElement{}); err != nil {
		return err
	}

	// 4. Choice: RowContents
	for _, rc := range t.RowContents {
		err := rc.MarshalXML(e, xml.StartElement{})
		if err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (t *Table) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
loop:
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "tblPr":
				prop := Property{}
				if err = d.DecodeElement(&prop, &elem); err != nil {
					return err
				}

				t.Property = prop
			case "tblGrid":
				grid := Grid{}
				if err = d.DecodeElement(&grid, &elem); err != nil {
					return err
				}

				t.Grid = grid
			case "tr":
				row := Row{}
				if err = d.DecodeElement(&row, &elem); err != nil {
					return err
				}

				t.RowContents = append(t.RowContents, RowContent{
					Row: &row,
				})

			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break loop
		}
	}

	return nil
}
