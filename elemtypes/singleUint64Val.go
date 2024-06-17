package elemtypes

import (
	"encoding/xml"
	"strconv"
)

// SingleUint64Val - Generic Element that has only one string-type attribute
// And the String type does not have validation
// dont use this if the element requires validation
type SingleUint64Val struct {
	Val uint64 `xml:"val,attr"`
}

func NewSingleUint64Val(value uint64) *SingleUint64Val {
	return &SingleUint64Val{
		Val: value,
	}
}

// MarshalXML implements the xml.Marshaler interface for the SingleUint64Val type.
// It encodes the instance into XML using the "w:ELEMENT_NAME" element with a "w:val" attribute.
func (s *SingleUint64Val) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: strconv.FormatUint(s.Val, 10)})
	err := e.EncodeElement("", start)

	return err
}
