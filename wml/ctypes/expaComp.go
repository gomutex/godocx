package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// This element specifies the amount by which each character shall be expanded or when the character is rendered in the document
//
// This property has an of stretching or compressing each character in the run, as opposed to the spacing element (§2.3.2.33) which expands/compresses the text by adding additional character pitch but not changing the width of the actual characters displayed on the line.
type ExpaComp struct {
	Val *stypes.TextScale `xml:"val,attr,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for the ExpaComp type.
// It encodes the instance into XML using the "w:XMLName" element with a "w:val" attribute.
func (ec ExpaComp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if ec.Val != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: ec.Val.ToStr()})
	}
	err := e.EncodeElement("", start)

	return err
}
