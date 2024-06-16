package dml

import (
	"encoding/xml"
)

// Non-Visual Drawing Properties
type CNvPr struct {
	ID          string `xml:"id,attr,omitempty"`
	Name        string `xml:"name,attr,omitempty"`
	Description string `xml:"descr,attr,omitempty"`
}

func (c *CNvPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "pic:cNvPr"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "id"}, Value: c.ID},
		{Name: xml.Name{Local: "name"}, Value: c.Name},
	}

	if c.Description != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "descr"}, Value: c.Description})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
