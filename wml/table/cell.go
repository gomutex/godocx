package table

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/docxpara"
)

type Cell struct {
	// 1.Table Cell Properties
	Property *CellProperty

	// 2.1 Choice:
	Contents []TCBlockContent

	//TODO: Remaining choices
}

func DefaultCell() *Cell {
	return &Cell{}
}

func (c *Cell) AddParagraph(text string) *docxpara.Paragraph {
	p := docxpara.AddParagraph(text)
	tblContent := TCBlockContent{
		Paragraph: p,
	}

	c.Contents = append(c.Contents, tblContent)

	return p
}

func (c *Cell) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:tc"

	if err = e.EncodeToken(start); err != nil {
		return err
	}

	//1.Table Cell Properties
	if c.Property != nil {
		if err = c.Property.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	//2.1 Choice
	for _, elem := range c.Contents {
		if err = elem.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (c *Cell) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
loop:
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "tcPr":
				prop := CellProperty{}
				if err = d.DecodeElement(&prop, &elem); err != nil {
					return err
				}

				c.Property = &prop
			case "p":
				para := docxpara.Paragraph{}
				if err = d.DecodeElement(&para, &elem); err != nil {
					return err
				}

				c.Contents = append(c.Contents, TCBlockContent{
					Paragraph: &para,
				})
			case "tbl":
				tbl := Table{}
				if err = d.DecodeElement(&tbl, &elem); err != nil {
					return err
				}

				c.Contents = append(c.Contents, TCBlockContent{
					Table: &tbl,
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

// Table Cell - ContentBlockContent
type TCBlockContent struct {
	Paragraph *docxpara.Paragraph
	Table     *Table
}

func (t TCBlockContent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.Paragraph != nil {
		return t.Paragraph.MarshalXML(e, xml.StartElement{})
	}

	if t.Table != nil {
		return t.Table.MarshalXML(e, xml.StartElement{})
	}

	return nil
}
