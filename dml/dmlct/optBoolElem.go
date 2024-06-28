package dmlct

import (
	"encoding/xml"

	"github.com/gomutex/godocx/dml/dmlst"
)

// Optional Bool Element: Helper element that only has one attribute which is optional
type OptBoolElem struct {
	Val dmlst.OptBool
}

func NewOptBoolElem(value bool) *OptBoolElem {
	return &OptBoolElem{
		Val: dmlst.NewOptBool(value),
	}
}

// Disable sets the value to false and valexists true
func (n *OptBoolElem) Disable() {
	n.Val = dmlst.NewOptBool(false)
}

// MarshalXML implements the xml.Marshaler interface for the Bold type.
// It encodes the instance into XML using the "w:XMLName" element with a "w:val" attribute.
func (n OptBoolElem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	if n.Val.Valid { // Add val attribute only if the val exists
		if n.Val.Bool {
			start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: "true"})
		} else {
			start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: "false"})
		}
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})

	// return e.EncodeElement("", start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for the type.
// It decodes the XML representation, extracting the value from the "w:val" attribute.
// The inner content of the XML element is skipped.
func (n *OptBoolElem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "val" {
			// If value is "true", then set it to true
			n.Val = dmlst.NewOptBool(a.Value == "true")
			break
		}
	}

	return d.Skip() // Skipping the inner content
}
