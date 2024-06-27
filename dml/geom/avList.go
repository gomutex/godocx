package geom

import (
	"encoding/xml"
)

// List of Shape Adjust Values
type AdjustValues struct {
	ShapeGuides []ShapeGuide `xml:"gd,omitempty"`
}

func (a AdjustValues) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "a:avLst"

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	for _, data := range a.ShapeGuides {
		err := data.MarshalXML(e, xml.StartElement{})
		if err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
