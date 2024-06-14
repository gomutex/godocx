package dml

import (
	"encoding/xml"
	"strconv"
)

type DocProp struct {
	ID          uint64
	Name        string
	Description string
}

func (d *DocProp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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

func (d *DocProp) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		switch a.Name.Local {
		case "descr":
			d.Description = a.Value
		case "name":
			d.Name = a.Value
		case "id":
			id, err := strconv.ParseUint(a.Value, 10, 32)
			if err != nil {
				return nil
			}
			d.ID = id
		}
	}

	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {

			default:
				if err = decoder.Skip(); err != nil {
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
