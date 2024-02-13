package elements

import (
	"encoding/xml"

	"github.com/gomutex/godocx/constants"
)

type TableProperty struct {
	Width         *TableWidth
	Justification *Justification
	Layout        *TableLayout
	Style         *TableStyle
	Indent        *TableIndent
	Margins       *TableCellMargins
	Borders       *TableBorders
}

func DefaultTableProperty() *TableProperty {
	return &TableProperty{}
}

func (t *TableProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:tblPr"

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	if t.Width != nil {
		if err = e.EncodeElement(t.Width, xml.StartElement{Name: xml.Name{Local: "w:tblW"}}); err != nil {
			return err
		}
	}

	if t.Indent != nil {
		if err = e.EncodeElement(t.Indent, xml.StartElement{Name: xml.Name{Local: "w:tblInd"}}); err != nil {
			return err
		}
	}

	if t.Justification != nil {
		if err = e.EncodeElement(t.Justification, xml.StartElement{Name: xml.Name{Local: "w:jc"}}); err != nil {
			return err
		}
	}

	if t.Layout != nil {
		if err = e.EncodeElement(t.Layout, xml.StartElement{Name: xml.Name{Local: "w:tblLayout"}}); err != nil {
			return err
		}
	}

	if t.Style != nil {
		if err = e.EncodeElement(t.Style, xml.StartElement{Name: xml.Name{Local: "w:tblStyle"}}); err != nil {
			return err
		}
	}

	if t.Margins != nil {
		if err = e.EncodeElement(t.Margins, xml.StartElement{Name: xml.Name{Local: "w:tblCellMar"}}); err != nil {
			return err
		}
	}

	if t.Borders != nil {
		if err = e.EncodeElement(t.Borders, xml.StartElement{Name: xml.Name{Local: "w:tblBorders"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (t *TableProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name {
			case xml.Name{Space: constants.WMLNamespace, Local: "tblW"}, xml.Name{Space: constants.AltWMLNamespace, Local: "tblW"}:
				tblWidth := TableWidth{}
				if err := d.DecodeElement(&tblWidth, &elem); err != nil {
					return err
				}
				t.Width = &tblWidth
			case xml.Name{Space: constants.WMLNamespace, Local: "tblInd"}, xml.Name{Space: constants.AltWMLNamespace, Local: "tblInd"}:
				tblIndent := TableIndent{}
				if err := d.DecodeElement(&tblIndent, &elem); err != nil {
					return err
				}
				t.Indent = &tblIndent
			case xml.Name{Space: constants.WMLNamespace, Local: "jc"}, xml.Name{Space: constants.AltWMLNamespace, Local: "jc"}:
				jc := Justification{}
				if err := d.DecodeElement(&jc, &elem); err != nil {
					return err
				}
				t.Justification = &jc
			case xml.Name{Space: constants.WMLNamespace, Local: "tblLayout"}, xml.Name{Space: constants.AltWMLNamespace, Local: "tblLayout"}:
				tl := TableLayout{}
				if err := d.DecodeElement(&tl, &elem); err != nil {
					return err
				}
				t.Layout = &tl
			case xml.Name{Space: constants.WMLNamespace, Local: "tblStyle"}, xml.Name{Space: constants.AltWMLNamespace, Local: "tblStyle"}:
				tblStyle := TableStyle{}
				if err := d.DecodeElement(&tblStyle, &elem); err != nil {
					return err
				}
				t.Style = &tblStyle
			case xml.Name{Space: constants.WMLNamespace, Local: "tblCellMar"}, xml.Name{Space: constants.AltWMLNamespace, Local: "tblCellMar"}:
				tblCellMargins := TableCellMargins{}
				if err := d.DecodeElement(&tblCellMargins, &elem); err != nil {
					return err
				}
				t.Margins = &tblCellMargins
			case xml.Name{Space: constants.WMLNamespace, Local: "tblBorders"}, xml.Name{Space: constants.AltWMLNamespace, Local: "tblBorders"}:
				tblBorders := TableBorders{}
				if err := d.DecodeElement(&tblBorders, &elem); err != nil {
					return err
				}
				t.Borders = &tblBorders
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
