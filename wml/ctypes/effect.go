package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

type Effect struct {
	Val *stypes.TextEffect `xml:"val,attr,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for the Effect type.
// It encodes the instance into XML using the "w:XMLName" element with a "w:val" attribute.
func (eff Effect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if eff.Val != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(*eff.Val)})
	}
	err := e.EncodeElement("", start)

	return err
}
