package elemtypes

import (
	"encoding/xml"
	"strconv"
)

// SingleIntVal - Generic Element that has only one string-type attribute
// And the String type does not have validation
// dont use this if the element requires validation
type SingleIntVal struct {
	Val uint64 `xml:"val,attr"`
}

func NewSingleIntVal(value uint64) *SingleIntVal {
	return &SingleIntVal{
		Val: value,
	}
}

// MarshalXML implements the xml.Marshaler interface for the SingleIntVal type.
// It encodes the instance into XML using the "w:ELEMENT_NAME" element with a "w:val" attribute.
func (s *SingleIntVal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: strconv.FormatUint(s.Val, 10)})
	err := e.EncodeElement("", start)

	return err
}
