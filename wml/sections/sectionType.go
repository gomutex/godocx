package sections

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

type SectionType struct {
	Val stypes.SectionMark `xml:"val,attr,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for the SectionType type.
// It encodes the SectionType to its corresponding XML representation.
func (s *SectionType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:type"
	if s.Val != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(s.Val)})
	}
	return e.EncodeElement("", start)
}
