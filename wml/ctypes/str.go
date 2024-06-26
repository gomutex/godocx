package ctypes

import (
	"encoding/xml"
)

// CTString - Generic Element that has only one string-type attribute
// And the String type does not have validation
// dont use this if the element requires validation
type CTString struct {
	Val string `xml:"val,attr"`
}

func NewCTString(value string) *CTString {
	return &CTString{
		Val: value,
	}
}

// MarshalXML implements the xml.Marshaler interface for the CTString type.
// It encodes the instance into XML using the "w:ELEMENT_NAME" element with a "w:val" attribute.
func (s *CTString) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: s.Val})
	err := e.EncodeElement("", start)

	return err
}
