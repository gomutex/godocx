package ctypes

import "encoding/xml"

// Range Markup elements
type RngMarkupElem struct {
}

func (r RngMarkupElem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// TODO:
	return nil
}
