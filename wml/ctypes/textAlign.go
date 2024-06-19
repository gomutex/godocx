package ctypes

import (
	"encoding/xml"
	"errors"

	"github.com/gomutex/godocx/wml/stypes"
)

// TextAlign represents the text alignment settings in a Word document.
type TextAlign struct {
	Val stypes.TextAlign `xml:"val,attr,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for the TextAlign type.
// It encodes the TextAlign to its corresponding XML representation.
func (s *TextAlign) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if s.Val == "" {
		return errors.New("Invalid TextAlign")
	}

	start.Name.Local = "w:textAlignment"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(s.Val)})
	return e.EncodeElement("", start)
}
