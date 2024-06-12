package wml

import (
	"encoding/xml"

	"github.com/gomutex/godocx/common/constants"
)

type TableRow struct {
	Cells        []*TableCell
	Property     *TableRowProperty
	HasNumbering bool
}

func DefaultTableRow() *TableRow {
	return &TableRow{
		Cells:    []*TableCell{},
		Property: &TableRowProperty{},
	}
}

func (r *TableRow) AddCell() *TableCell {
	cell := DefaultCell()
	r.Cells = append(r.Cells, cell)
	return cell
}

// TODO  Implement Marshal and Unmarshal properly for all fields

func (tr *TableRow) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:tr"

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	for _, cell := range tr.Cells {
		if err = cell.MarshalXML(e, start); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (tr *TableRow) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name {
			case xml.Name{Space: constants.WMLNamespace, Local: "tc"}, xml.Name{Space: constants.AltWMLNamespace, Local: "tc"}:
				cell := TableCell{}
				if err := d.DecodeElement(&cell, &elem); err != nil {
					return err
				}
				tr.Cells = append(tr.Cells, &cell)
			default:
				if err = d.Skip(); err != nil {
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
