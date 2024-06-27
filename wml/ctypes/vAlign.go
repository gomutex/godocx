package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// Vertical Text Alignment on Page
// Complex ELement: CT_VerticalJc
type VAlign struct {
	//Subscript/Superscript Value
	Val stypes.VerticalJc `xml:"val,attr,omitempty"`
}

func (v VAlign) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:vAlign"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(v.Val)})

	return e.EncodeElement("", start)
}
