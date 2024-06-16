package elemtypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/types"
)

type NullBoolElem struct {
	Val types.NullBool
}

func NewNullBoolElem(value bool) *NullBoolElem {
	return &NullBoolElem{
		Val: types.NewNullBool(value),
	}
}

// Disable sets the value to false and valexists true
func (n *NullBoolElem) Disable() {
	n.Val = types.NewNullBool(false)
}

// MarshalXML implements the xml.Marshaler interface for the Bold type.
// It encodes the instance into XML using the "w:XMLName" element with a "w:val" attribute.
func (n *NullBoolElem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	if n.Val.Valid { // Add val attribute only if the val exists
		if n.Val.Bool {
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
func (n *NullBoolElem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "val" {
			// If value is "true", then set it to true
			n.Val = types.NewNullBool(a.Value == "true")
			break
		}
	}

	return d.Skip() // Skipping the inner content
}
