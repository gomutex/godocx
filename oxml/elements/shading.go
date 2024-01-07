package elements

import (
	"encoding/xml"
)

// ShadingType represents the shading type.
type ShadingType string

// Shading represents the shading properties.
type Shading struct {
	ShdType ShadingType
	Color   string
	Fill    string
}

// DefaultShading creates a new Shading with default values.
func DefaultShading() *Shading {
	return &Shading{
		ShdType: "Clear",
		Color:   "auto",
		Fill:    "FFFFFF",
	}
}

// NewShading creates a new Shading.
func NewShading() *Shading {
	return DefaultShading()
}

// Color sets the color for the shading.
func (s *Shading) SetColor(color string) *Shading {
	s.Color = color
	return s
}

// Fill sets the fill for the shading.
func (s *Shading) SetFill(fill string) *Shading {
	s.Fill = fill
	return s
}

// ShadingType sets the shading type for the shading.
func (s *Shading) SetShadingType(shdType ShadingType) *Shading {
	s.ShdType = shdType
	return s
}

func (s *Shading) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if s.Fill == "" && s.Color == "" && s.ShdType == "" {
		return nil
	}

	start.Name.Local = "w:shd"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "w:fill"}, Value: s.Fill},
		{Name: xml.Name{Local: "w:val"}, Value: string(s.ShdType)},
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if err := e.EncodeToken(start.End()); err != nil {
		return err
	}

	return nil
}

// UnmarshalXML unmarshals XML to Shading.
func (s *Shading) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Extracting attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "fill":
			s.Fill = attr.Value
		case "val":
			s.ShdType = ShadingType(attr.Value)
		}
	}

	return d.Skip() // Skipping the inner content
}
