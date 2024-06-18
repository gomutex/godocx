package sections

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// TextDirection represents the text direction settings in a Word document.
type TextDirection struct {
	Val stypes.TextDirection `xml:"val,attr,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for the TextDirection type.
// It encodes the TextDirection to its corresponding XML representation.
func (s *TextDirection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:textDirection"
	if s.Val != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(s.Val)})
	}
	return e.EncodeElement("", start)
}
