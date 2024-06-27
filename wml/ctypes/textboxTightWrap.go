package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// TextboxTightWrap represents the tight wrap settings in a Word document text box.
type TextboxTightWrap struct {
	Val stypes.TextboxTightWrap `xml:"val,attr,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for the TextboxTightWrap type.
// It encodes the TextboxTightWrap to its corresponding XML representation.
func (s TextboxTightWrap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:textboxTightWrap"
	if s.Val != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(s.Val)})
	}
	return e.EncodeElement("", start)
}
