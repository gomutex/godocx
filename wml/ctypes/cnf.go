package ctypes

import (
	"encoding/xml"
)

type Cnf struct {
	Val string `xml:"val,attr"`
}

func (c *Cnf) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:cnfStyle"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: c.Val})
	err := e.EncodeElement("", start)

	return err
}
