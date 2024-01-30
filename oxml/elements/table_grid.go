package elements

import (
	"encoding/xml"
	"strconv"
)

// TableGrid represents the grid columns of a table.
type TableGrid struct {
	Grid []uint64
}

// NewTableGrid creates a new TableGrid instance.
func NewTableGrid(grid []uint64) *TableGrid {
	return &TableGrid{Grid: grid}
}

func DefaultTableGrid() *TableGrid {
	return &TableGrid{Grid: []uint64{}}
}

// MarshalXML implements the xml.Marshaler interface for TableGrid.
func (t *TableGrid) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:tblGrid"
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	gridType := "dxa"

	for _, g := range t.Grid {
		gridCol := xml.StartElement{Name: xml.Name{Local: "w:gridCol"}}
		gridCol.Attr = append(gridCol.Attr, xml.Attr{Name: xml.Name{Local: "w:w"}, Value: strconv.FormatUint(uint64(g), 10)})
		gridCol.Attr = append(gridCol.Attr, xml.Attr{Name: xml.Name{Local: "w:type"}, Value: gridType})

		if err := e.EncodeElement("", gridCol); err != nil {
			return err
		}
	}

	if err := e.EncodeToken(start.End()); err != nil {
		return err
	}

	return nil
}

// UnmarshalXML implements the xml.Unmarshaler interface for TableGrid.
func (t *TableGrid) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		elem, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := elem.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "gridCol":
				var width string
				for _, a := range elem.Attr {
					if a.Name.Local == "w" {
						width = a.Value
					}
				}

				if width != "" {
					val, err := strconv.ParseUint(width, 10, 0)
					if err != nil {
						return err
					}
					t.Grid = append(t.Grid, val)
				}
				// Skip inner content of gridCol
				if err := d.Skip(); err != nil {
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
