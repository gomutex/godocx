package ctypes

import (
	"encoding/xml"
	"strconv"

	"github.com/gomutex/godocx/wml/stypes"
)

// East Asian Typography Settings
type EALayout struct {
	ID           *int                    `xml:"id,attr,omitempty"`
	Combine      *stypes.OnOff           `xml:"combine,attr,omitempty"`
	CombineBrkts *stypes.CombineBrackets `xml:"combineBrackets,attr,omitempty"`
	Vert         *stypes.OnOff           `xml:"vert,attr,omitempty"`
	VertCompress *stypes.OnOff           `xml:"vertCompress,attr,omitempty"`
}

func (ea EALayout) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:eastAsianLayout"

	if ea.ID != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"}, Value: strconv.Itoa(*ea.ID)})
	}
	if ea.Combine != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:combine"}, Value: string(*ea.Combine)})
	}
	if ea.CombineBrkts != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:combineBrackets"}, Value: string(*ea.CombineBrkts)})
	}
	if ea.Vert != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:vert"}, Value: string(*ea.Vert)})
	}
	if ea.VertCompress != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:vertCompress"}, Value: string(*ea.VertCompress)})
	}

	return e.EncodeElement("", start)
}
