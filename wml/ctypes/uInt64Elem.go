package ctypes

import (
	"encoding/xml"
	"strconv"
)

// Uint64Elem - Gomplex type that contains single val attribute which is type of uint64
// can be used where w:ST_UnsignedDecimalNumber is applicable
// example: ST_HpsMeasure
type Uint64Elem struct {
	Val uint64 `xml:"val,attr"`
}

func NewUint64Elem(value uint64) *Uint64Elem {
	return &Uint64Elem{
		Val: value,
	}
}

// MarshalXML implements the xml.Marshaler interface for the Uint64Elem type.
// It encodes the instance into XML using the "w:ELEMENT_NAME" element with a "w:val" attribute.
func (s *Uint64Elem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: strconv.FormatUint(s.Val, 10)})
	err := e.EncodeElement("", start)

	return err
}
