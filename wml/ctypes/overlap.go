package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// Floating Table Allows Other Tables to Overlap
type Overlap struct {
	//Floating Table Overlap Setting
	Val stypes.TblOverlap `xml:"val,attr"`
}

func (o Overlap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:tblOverlap"

	attr := xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(o.Val)}
	start.Attr = append(start.Attr, attr)

	return e.EncodeElement("", start)
}
