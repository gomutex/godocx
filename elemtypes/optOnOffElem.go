package elemtypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// OptOnOffElem helper struct that has only one optional field which is OnOff type
type OptOnOffElem struct {
	Val stypes.OnOff `xml:"val,attr,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for the OnOffElem type.
// It encodes the OnOffElem to its corresponding XML representation.
func (s *OptOnOffElem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if s.Val != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(s.Val)})
	}
	return e.EncodeElement("", start)
}
