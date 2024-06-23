package dmlct

import (
	"encoding/xml"
	"strconv"
)

// RelativeRect represents a Relative Rectangle structure with abbreviated attributes.
type RelativeRect struct {
	Top    *int `xml:"t,attr,omitempty"` // Top margin
	Left   *int `xml:"l,attr,omitempty"` // Left margin
	Bottom *int `xml:"b,attr,omitempty"` // Bottom margin
	Right  *int `xml:"r,attr,omitempty"` // Right margin
}

// MarshalXML implements the xml.Marshaler interface for RelativeRect.
func (r RelativeRect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = []xml.Attr{}

	if r.Top != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "t"}, Value: strconv.Itoa(*r.Top)})
	}

	if r.Left != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "l"}, Value: strconv.Itoa(*r.Left)})
	}

	if r.Bottom != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "b"}, Value: strconv.Itoa(*r.Bottom)})
	}

	if r.Right != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"}, Value: strconv.Itoa(*r.Right)})
	}

	return e.EncodeElement("", start)
}
