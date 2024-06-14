package elemtypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/types"
)

type NullBoolElem struct {
	Local string
	Val   types.NullBool
}

func DefaultNullBoolElem(local string) *NullBoolElem {
	return &NullBoolElem{
		Local: local,
	}
}

func NewNullBoolElem(local string, value bool) *NullBoolElem {
	return &NullBoolElem{
		Local: local,
		Val:   types.NewNullBool(value),
	}
}

// Disable sets the value to false and valexists true
func (self *NullBoolElem) Disable() {
	self.Val = types.NewNullBool(false)
}

// MarshalXML implements the xml.Marshaler interface for the Bold type.
// It encodes the instance into XML using the "w:XMLName" element with a "w:val" attribute.
func (self *NullBoolElem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name.Local = self.Local

	if self.Val.Valid { // Add val attribute only if the val exists
		if self.Val.Bool {
			start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: "true"})
		} else {
			start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: "false"})
		}
	}
	err := e.EncodeElement("", start)

	return err
}

// UnmarshalXML implements the xml.Unmarshaler interface for the type.
// It decodes the XML representation, extracting the value from the "w:val" attribute.
// The inner content of the XML element is skipped.
func (self *NullBoolElem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "val" {
			// If value is "true", then set it to true
			self.Val = types.NewNullBool(a.Value == "true")
			break
		}
	}

	return d.Skip() // Skipping the inner content
}
