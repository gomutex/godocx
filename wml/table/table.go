package table

import "encoding/xml"

type Table struct {
	Rows         []TableChild
	Grid         *TableGrid
	Property     *TableProperty
	HasNumbering bool
}

type TableChild struct {
	TableRow *TableRow
}

func DefaultTable() *Table {
	return &Table{}
}

func (t *Table) AddRow() *TableRow {
	row := DefaultTableRow()
	t.Rows = append(t.Rows, TableChild{
		TableRow: row,
	})
	return row
}

func (t *Table) Indent(indent int) {
	if t.Property == nil {
		t.Property = DefaultTableProperty()
	}
	t.Property.Indent = NewTableIndent(indent, WidthTypeAuto)
}

func (t *Table) Style(value string) {
	if t.Property == nil {
		t.Property = DefaultTableProperty()
	}
	t.Property.Style = NewTableStyle(value)
}

// TODO  Implement Marshal and Unmarshal properly for all fields

func (t *Table) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:tbl"

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	if t.Property != nil {
		err := t.Property.MarshalXML(e, xml.StartElement{})
		if err != nil {
			return err
		}
	}

	for _, row := range t.Rows {
		err := row.TableRow.MarshalXML(e, xml.StartElement{})
		if err != nil {
			return err
		}
	}

	if t.Grid != nil {
		err := t.Grid.MarshalXML(e, xml.StartElement{})
		if err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})

}

func (t *Table) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "tblPr":
				t.Property = &TableProperty{}
				if err := d.DecodeElement(t.Property, &elem); err != nil {
					return err
				}
			case "tr":
				row := TableRow{}
				if err := d.DecodeElement(&row, &elem); err != nil {
					return err
				}
				t.Rows = append(t.Rows, TableChild{
					TableRow: &row,
				})
			case "tblGrid":
				t.Grid = &TableGrid{}
				if err := d.DecodeElement(t.Grid, &elem); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
}
