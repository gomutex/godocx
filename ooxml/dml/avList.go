package dml

import (
	"encoding/xml"
)

type ShapeGuide struct {
	Name    string
	Formula string
}

func (s *ShapeGuide) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:gd"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "name"}, Value: s.Name},
		{Name: xml.Name{Local: "fmla"}, Value: s.Formula},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (s *ShapeGuide) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "name" {
			s.Name = a.Value
		} else if a.Name.Local == "fmla" {
			s.Formula = a.Value
		}
	}

	return nil
}

// List of Shape Adjust Values
type AdjustValues struct {
	ShapeGuides []ShapeGuide
}

func (a *AdjustValues) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "a:avLst"

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	for _, data := range a.ShapeGuides {
		err := data.MarshalXML(e, start)
		if err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (a *AdjustValues) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {

loop:
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "t":
				gd := ShapeGuide{}
				if err := d.DecodeElement(&gd, &elem); err != nil {
					return err
				}

				a.ShapeGuides = append(a.ShapeGuides, gd)

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