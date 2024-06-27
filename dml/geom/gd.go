package geom

import "encoding/xml"

type ShapeGuide struct {
	Name    string `xml:"name,attr,omitempty"`
	Formula string `xml:"fmla,attr,omitempty"`
}

func (s ShapeGuide) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
