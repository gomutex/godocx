package dml

import (
	"encoding/xml"
	"strconv"
)

type DocProp struct {
	ID          uint64 `xml:"id,attr,omitempty"`
	Name        string `xml:"name,attr,omitempty"`
	Description string `xml:"descr,attr,omitempty"`

	//TODO: Remaining attrs & child elements
}

func (d DocProp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:docPr"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "id"}, Value: strconv.FormatUint(d.ID, 10)},
		{Name: xml.Name{Local: "name"}, Value: d.Name},
	}

	if d.Description != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "descr"}, Value: d.Description})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
