package ctypes

import "encoding/xml"

type RowContent struct {
	Row *Row `xml:"tr,omitempty"`
}

func (r RowContent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Row != nil {
		return r.Row.MarshalXML(e, xml.StartElement{})
	}
	return nil
}
