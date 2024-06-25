package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

type Merge struct {
	Val stypes.MergeCell `xml:"val,attr"`
}

func NewMerge(value stypes.MergeCell) *Merge {
	return &Merge{
		Val: value,
	}
}

func (m *Merge) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(m.Val)})

	return e.EncodeElement("", start)
}
