package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// PageNumbering represents the page numbering format in a Word document.
type PageNumbering struct {
	Format stypes.NumFmt `xml:"fmt,attr,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for the PageNumbering type.
// It encodes the PageNumbering to its corresponding XML representation.
func (p PageNumbering) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:pgNumType"
	if p.Format != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:fmt"}, Value: string(p.Format)})
	}
	return e.EncodeElement("", start)
}
