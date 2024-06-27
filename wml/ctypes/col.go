package ctypes

import (
	"encoding/xml"
	"strconv"
)

// Grid Column Definition
type Column struct {
	Width *uint64 `xml:"w,attr,omitempty"` //Grid Column Width
}

func (c Column) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:gridCol"

	if c.Width != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:w"}, Value: strconv.FormatUint(*c.Width, 10)})
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
