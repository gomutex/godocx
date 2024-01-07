package elements

import (
	"encoding/xml"
	"strconv"
)

// SzCs represents the font size of a text or element.
type SzCs struct {
	Value int
}

// NewSzCs creates a new SzCs with the specified font size value.
func NewSzCs(value int) *SzCs {
	return &SzCs{Value: value}
}

func (s *SzCs) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:szCs"
	if s.Value != 0 {
		start.Attr = []xml.Attr{{Name: xml.Name{Local: "w:val"}, Value: strconv.Itoa(s.Value)}}
		return e.EncodeElement("", start)
	}
	return nil
}

func (s *SzCs) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "val" {
			valueStr := a.Value
			if valueStr != "" {
				value, err := strconv.ParseInt(valueStr, 10, 0)
				if err != nil {
					return err
				}
				s.Value = int(value)
			}
			break
		}
	}

	return d.Skip() // Skipping the inner content
}
