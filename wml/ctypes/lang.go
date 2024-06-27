package ctypes

import (
	"encoding/xml"
)

// Languages for Run Content
type Lang struct {
	Val      *string `xml:"val,attr,omitempty"`
	EastAsia *string `xml:"eastAsia,attr,omitempty"`
	Bidi     *string `xml:"bidi,attr,omitempty"`
}

func (l Lang) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:lang"
	if l.Val != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: *l.Val})
	}
	if l.EastAsia != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:eastAsia"}, Value: *l.EastAsia})
	}
	if l.Bidi != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:bidi"}, Value: *l.Bidi})
	}
	return e.EncodeElement("", start)
}
