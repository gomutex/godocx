package docxrun

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/simpletypes"
)

type VertAlign struct {
	Val simpletypes.VerticalAlignRun `xml:"val,attr,omitempty"`
}

func (v *VertAlign) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:vertAlign"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(v.Val)})

	return e.EncodeElement("", start)
}
