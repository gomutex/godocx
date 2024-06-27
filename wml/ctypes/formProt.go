package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// FormProt represents the text direction settings in a Word document.
type FormProt struct {
	Val stypes.OnOff `xml:"val,attr,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for the FormProt type.
// It encodes the FormProt to its corresponding XML representation.
func (s FormProt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:formProt"
	if s.Val != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(s.Val)})
	}
	return e.EncodeElement("", start)
}
