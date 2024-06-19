package elemtypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// OptBinFlagElem helper struct that has only one optional field which is BinFlag type
type OptBinFlagElem struct {
	Val stypes.BinFlag `xml:"val,attr,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for the BinFlagElem type.
// It encodes the BinFlagElem to its corresponding XML representation.
func (s *OptBinFlagElem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if s.Val != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(s.Val)})
	}
	return e.EncodeElement("", start)
}
