package dml

import (
	"encoding/xml"
)

type PositionType struct {
	XAttr int `xml:"x,attr"`
	YAttr int `xml:"y,attr"`
}

type PoistionH struct {
	RelativeFrom RelativeFromHType
	PosOffset    int `xml:"posOffset"`
}

type PoistionV struct {
	RelativeFrom RelativeFromVType
	PosOffset    int `xml:"posOffset"`
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

func (p *PoistionH) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "relativeFrom" {
			p.RelativeFrom = RelativeFromHType(a.Value)
			break
		}
	}

	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			{
				switch elem.Name.Local {
				case "posOffset":
					if err = d.DecodeElement(&p.PosOffset, &elem); err != nil {
						return err
					}
				default:
					if err = d.Skip(); err != nil {
						return err
					}
				}
			}
		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
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

func (p *PoistionV) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "relativeFrom" {
			p.RelativeFrom = RelativeFromVType(a.Value)
			break
		}
	}

	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			{
				switch elem.Name.Local {
				case "posOffset":
					if err = d.DecodeElement(&p.PosOffset, &elem); err != nil {
						return err
					}
				default:
					if err = d.Skip(); err != nil {
						return err
					}
				}
			}
		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
}
