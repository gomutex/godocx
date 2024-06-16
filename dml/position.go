package dml

import (
	"encoding/xml"
)

type PositionType struct {
	XAttr int `xml:"x,attr"`
	YAttr int `xml:"y,attr"`
}

type PoistionH struct {
	RelativeFrom RelativeFromHType `xml:"relativeFrom,attr"`
	PosOffset    int               `xml:"posOffset"`
}

type PoistionV struct {
	RelativeFrom RelativeFromVType `xml:"relativeFrom,attr"`
	PosOffset    int               `xml:"posOffset"`
}

func (p *PoistionH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:positionH"

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "relativeFrom"}, Value: string(p.RelativeFrom)})

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	offsetElem := xml.StartElement{Name: xml.Name{Local: "wp:posOffset"}}
	if err = e.EncodeElement(p.PosOffset, offsetElem); err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (p *PoistionV) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:positionV"

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "relativeFrom"}, Value: string(p.RelativeFrom)})

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	offsetElem := xml.StartElement{Name: xml.Name{Local: "wp:posOffset"}}
	if err = e.EncodeElement(p.PosOffset, offsetElem); err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
