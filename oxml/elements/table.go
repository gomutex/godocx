package elements

import "encoding/xml"

type Table struct {
	Rows         []TableChild
	Grid         *TableGrid
	HasNumbering bool
	Property     *TableProperty
}

type TableChild struct {
	TableCell TableCell
}

func DefaultTable() *Table {
	return &Table{}
}

// MarshalXML implements the xml.Marshaler interface for Table.
func (t *Table) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:tbl"
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// Serialize table properties
	if t.Property != nil {
		if err := e.EncodeElement(&t.Property, xml.StartElement{Name: xml.Name{Local: "w:tblPr"}}); err != nil {
			return err
		}
	}

	// Serialize table grid
	if t.Grid != nil {
		if err := e.EncodeElement(&t.Grid, xml.StartElement{Name: xml.Name{Local: "w:tblGrid"}}); err != nil {
			return err
		}
	}

	// Serialize table rows
	for _, child := range t.Rows {
		if err := e.EncodeElement(&child, xml.StartElement{Name: xml.Name{Local: "w:tr"}}); err != nil {
			return err
		}
	}

	if err := e.EncodeToken(start.End()); err != nil {
		return err
	}

	return nil
}

// UnmarshalXML implements the xml.Unmarshaler interface for Table.
func (t *Table) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	t.Rows = nil
	t.HasNumbering = false

	for {
		elem, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := elem.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "tblPr":
				t.Property = DefaultTableProperty()
				if err := d.DecodeElement(&t.Property, &elem); err != nil {
					return err
				}
			case "tblGrid":
				t.Grid = DefaultTableGrid()
				if err := d.DecodeElement(&t.Grid, &elem); err != nil {
					return err
				}
			case "tr":
				child := TableChild{}
				if err := d.DecodeElement(&child, &elem); err != nil {
					return err
				}
				t.Rows = append(t.Rows, child)
			}
		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
}

// MarshalXML implements the xml.Marshaler interface for TableChild.
func (tc *TableChild) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:tr"
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// Serialize table cell
	if err := e.EncodeElement(&tc.TableCell, xml.StartElement{Name: xml.Name{Local: "w:tc"}}); err != nil {
		return err
	}

	if err := e.EncodeToken(start.End()); err != nil {
		return err
	}

	return nil
}

// UnmarshalXML implements the xml.Unmarshaler interface for TableChild.
func (tc *TableChild) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tc.TableCell = TableCell{}

	for {
		elem, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := elem.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "tc":
				if err := d.DecodeElement(&tc.TableCell, &elem); err != nil {
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
