package table

import (
	"encoding/xml"

	"github.com/gomutex/godocx/common/constants"
	"github.com/gomutex/godocx/wml/formatting"
)

type TableCellProperty struct {
	Shading *formatting.Shading
}

func (t *TableCellProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:tcPr"

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	if t.Shading != nil {
		if err = e.EncodeElement(t.Shading, xml.StartElement{Name: xml.Name{Local: "w:shd"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (t *TableCellProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name {
			case xml.Name{Space: constants.WMLNamespace, Local: "shd"}, xml.Name{Space: constants.AltWMLNamespace, Local: "shd"}:
				shd := formatting.Shading{}
				if err = d.DecodeElement(&shd, &elem); err != nil {
					return err
				}
				t.Shading = &shd

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
