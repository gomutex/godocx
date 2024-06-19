package docxpara

import (
	"encoding/xml"
	"strconv"
)

// Revision Information for Paragraph Properties
type PPrChange struct {
	ID       int            `xml:"id,attr"`
	Author   string         `xml:"author,attr"`
	Date     *string        `xml:"date,attr,omitempty"`
	ParaProp *ParagraphProp `xml:"pPr"`
}

func (p PPrChange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:pPrChange"

	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "id"}, Value: strconv.Itoa(p.ID)},
		{Name: xml.Name{Local: "author"}, Value: p.Author},
	}

	if p.Date != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "date"}, Value: *p.Date})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if p.ParaProp != nil {
		if err := p.ParaProp.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})

}
