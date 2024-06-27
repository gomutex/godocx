package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// Subscript/Superscript Text
// Complex ELement: CT_VerticalAlignRun
type VertAlignRun struct {
	//Subscript/Superscript Value
	Val stypes.VerticalAlignRun `xml:"val,attr,omitempty"`
}

func (v VertAlignRun) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:vertAlign"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(v.Val)})

	return e.EncodeElement("", start)
}
