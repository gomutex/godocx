package hdrftr

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/simpletypes"
)

// Footer Reference
type FooterReference struct {
	Type simpletypes.HdrFtrType `xml:"type,attr"` //Footer or Footer Type
	ID   string                 `xml:"id,attr"`   //Relationship to Part
}

func (h *FooterReference) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:footerReference"

	if h.Type != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:type"}, Value: string(h.Type)})
	}

	if h.ID != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"}, Value: h.ID})
	}

	return e.EncodeElement("", start)
}
