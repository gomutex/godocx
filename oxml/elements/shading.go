package elements

import (
	"encoding/xml"
)

// ShadingType represents the shading type.
// ShadingType represents different shading types.
type ShadingType string

const (
	ShadingTypeNil                   ShadingType = "Nil"
	ShadingTypeClear                 ShadingType = "Clear"
	ShadingTypeSolid                 ShadingType = "Solid"
	ShadingTypeHorzStripe            ShadingType = "HorzStripe"
	ShadingTypeVertStripe            ShadingType = "VertStripe"
	ShadingTypeReverseDiagStripe     ShadingType = "ReverseDiagStripe"
	ShadingTypeDiagStripe            ShadingType = "DiagStripe"
	ShadingTypeHorzCross             ShadingType = "HorzCross"
	ShadingTypeDiagCross             ShadingType = "DiagCross"
	ShadingTypeThinHorzStripe        ShadingType = "ThinHorzStripe"
	ShadingTypeThinVertStripe        ShadingType = "ThinVertStripe"
	ShadingTypeThinReverseDiagStripe ShadingType = "ThinReverseDiagStripe"
	ShadingTypeThinDiagStripe        ShadingType = "ThinDiagStripe"
	ShadingTypeThinHorzCross         ShadingType = "ThinHorzCross"
	ShadingTypeThinDiagCross         ShadingType = "ThinDiagCross"
	ShadingTypePct5                  ShadingType = "Pct5"
	ShadingTypePct10                 ShadingType = "Pct10"
	ShadingTypePct12                 ShadingType = "Pct12"
	ShadingTypePct15                 ShadingType = "Pct15"
	ShadingTypePct20                 ShadingType = "Pct20"
	ShadingTypePct25                 ShadingType = "Pct25"
	ShadingTypePct30                 ShadingType = "Pct30"
	ShadingTypePct35                 ShadingType = "Pct35"
	ShadingTypePct37                 ShadingType = "Pct37"
	ShadingTypePct40                 ShadingType = "Pct40"
	ShadingTypePct45                 ShadingType = "Pct45"
	ShadingTypePct50                 ShadingType = "Pct50"
	ShadingTypePct55                 ShadingType = "Pct55"
	ShadingTypePct60                 ShadingType = "Pct60"
	ShadingTypePct62                 ShadingType = "Pct62"
	ShadingTypePct65                 ShadingType = "Pct65"
	ShadingTypePct70                 ShadingType = "Pct70"
	ShadingTypePct75                 ShadingType = "Pct75"
	ShadingTypePct80                 ShadingType = "Pct80"
	ShadingTypePct85                 ShadingType = "Pct85"
	ShadingTypePct87                 ShadingType = "Pct87"
	ShadingTypePct90                 ShadingType = "Pct90"
	ShadingTypePct95                 ShadingType = "Pct95"
)

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
