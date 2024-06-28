package ctypes

import (
	"encoding/xml"
	"strconv"
)

type FitText struct {
	Val uint64 `xml:"val,attr"`
	ID  *int   `xml:"id,attr,omitempty"`
}

func (f FitText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:fitText"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: strconv.FormatUint(f.Val, 10)})

	if f.ID != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"}, Value: strconv.Itoa(*f.ID)})
	}

	return e.EncodeElement("", start)
}
