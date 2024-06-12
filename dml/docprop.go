package dml

import (
	"encoding/xml"
	"strconv"
)

type DocProp struct {
	ID   uint64 // cx
	Name string // cy
}

func (d *DocProp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:docPr"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "id"}, Value: strconv.FormatUint(d.ID, 10)},
		{Name: xml.Name{Local: "name"}, Value: d.Name},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (d *DocProp) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "id" {
			id, err := strconv.ParseUint(a.Value, 10, 32)
			if err != nil {
				return nil
			}
			d.ID = id
		} else if a.Name.Local == "name" {
			d.Name = a.Value
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
