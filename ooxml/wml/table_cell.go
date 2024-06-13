package wml

import (
	"encoding/xml"

	"github.com/gomutex/godocx/common/constants"
)

type TableCell struct {
	Children     []*TableCellContent
	Property     TableCellProperty
	HasNumbering bool
}

type TableCellContent struct {
	Paragraph *Paragraph
	Table     *Table
	// StructuredDataTag *StructuredDataTag
	// TableOfContents   *TableOfContents
}

func DefaultCell() *TableCell {
	return &TableCell{
		Children: []*TableCellContent{},
	}
}

func (c *TableCell) AddParagraph(text string) *Paragraph {
	p := AddParagraph(text)
	tblContent := TableCellContent{
		Paragraph: p,
	}

	c.Children = append(c.Children, &tblContent)

	return p
}

// TODO  Implement Marshal and Unmarshal properly for all fields

func (tc *TableCell) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:tc"

	if err = e.EncodeToken(start); err != nil {
		return err
	}

	for _, data := range tc.Children {
		if data.Paragraph != nil {
			err := data.Paragraph.MarshalXML(e, start)
			if err != nil {
				return err
			}
		}

		if data.Table != nil {
			err := data.Table.MarshalXML(e, start)
			if err != nil {
				return err
			}
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (tc *TableCell) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name {
			case xml.Name{Space: constants.WMLNamespace, Local: "p"}, xml.Name{Space: constants.AltWMLNamespace, Local: "p"}:
				p := DefaultParagraph()
				if err = d.DecodeElement(p, &elem); err != nil {
					return err
				}

				tc.Children = append(tc.Children, &TableCellContent{Paragraph: p})
			case xml.Name{Space: constants.WMLNamespace, Local: "tbl"}, xml.Name{Space: constants.AltWMLNamespace, Local: "tbl"}:
				tbl := DefaultTable()
				if err = d.DecodeElement(tbl, &elem); err != nil {
					return err
				}

				tc.Children = append(tc.Children, &TableCellContent{Table: tbl})

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
