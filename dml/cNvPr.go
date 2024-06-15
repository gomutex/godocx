package dml

import "encoding/xml"

// Non-Visual Drawing Properties
type CNvPr struct {
	ID          string
	Name        string
	Description string
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

func (c *CNvPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		switch a.Name.Local {
		case "descr":
			c.Description = a.Value
		case "name":
			c.Name = a.Value
		case "id":
			c.ID = a.Value
		}
	}

	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			}
		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
}
