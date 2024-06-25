package table

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// Table Cell Vertical Alignment
type VertAlign struct {
	//Vertical Alignment Setting
	Val stypes.VerticalJc `xml:"val,attr"`
}

func (v *VertAlign) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:vAlign"

	attr := xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(v.Val)}
	start.Attr = append(start.Attr, attr)

	return e.EncodeElement("", start)
}
