package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// Emphasis Mark
type Em struct {
	Val stypes.Em `xml:"val,attr"`
}

func (f Em) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:em"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(f.Val)})

	return e.EncodeElement("", start)
}
