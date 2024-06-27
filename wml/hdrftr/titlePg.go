package hdrftr

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// Different First Page Headers and Footers
type TitlePg struct {
	Val stypes.OnOff `xml:"val,attr,omitempty"`
}

func (t TitlePg) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:titlePg"

	if t.Val != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(t.Val)})
	}

	return e.EncodeElement("", start)
}
