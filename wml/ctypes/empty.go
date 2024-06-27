package ctypes

import "encoding/xml"

type Empty struct {
}

func (s Empty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement("", start)
}
