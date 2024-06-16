package sections

import (
	"encoding/xml"

	"github.com/gomutex/godocx/internal/helpers"
	"github.com/gomutex/godocx/wml/simpletypes"
)

// PageNumbering represents the page numbering format in a Word document.
type PageNumbering struct {
	Format simpletypes.NumFmt
}

// MarshalXML implements the xml.Marshaler interface for the PageNumbering type.
// It encodes the PageNumbering to its corresponding XML representation.
func (p *PageNumbering) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:pgNumType"
	if p.Format != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:fmt"}, Value: string(p.Format)})
	}
	return e.EncodeElement("", start)
}

func (p *PageNumbering) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		switch a.Name.Local {
		case "fmt":
			format, err := simpletypes.NumFmtFromStr(a.Value)
			if err != nil {
				return nil
			}
			p.Format = format
		}

	}

	return helpers.SkipUntilEnd(d, start.End())
}
