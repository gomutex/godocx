package elemtypes

import (
	"encoding/xml"
)

// SingleStrVal - Generic Element that has only one string-type attribute
// And the String type does not have validation
// dont use this if the element requires validation
type SingleStrVal struct {
	Val string `xml:"val,attr"`
}

func NewSingleStrVal(value string) *SingleStrVal {
	return &SingleStrVal{
		Val: value,
	}
}

// MarshalXML implements the xml.Marshaler interface for the SingleStrVal type.
// It encodes the instance into XML using the "w:ELEMENT_NAME" element with a "w:val" attribute.
func (s *SingleStrVal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: s.Val})
	err := e.EncodeElement("", start)

	return err
}
