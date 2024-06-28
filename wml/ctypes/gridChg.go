package ctypes

import (
	"encoding/xml"
	"strconv"
)

// Revision Information for Table Grid Column Definitions
type GridChange struct {
	ID int `xml:"id,attr,omitempty"` //Annotation Identifier
}

func (g GridChange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:tblGridChange"

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"}, Value: strconv.Itoa(g.ID)})

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
