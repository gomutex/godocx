package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// Shading represents the shading properties for a run in a WordprocessingML document.
type Shading struct {
	Val            stypes.Shading     `xml:"val,attr"`
	Color          *string            `xml:"color,attr,omitempty"`
	ThemeColor     *stypes.ThemeColor `xml:"themeColor,attr,omitempty"`
	ThemeFill      *stypes.ThemeColor `xml:"themeFill,attr,omitempty"`
	ThemeTint      *string            `xml:"themeTint,attr,omitempty"`
	ThemeShade     *string            `xml:"themeShade,attr,omitempty"`
	Fill           *string            `xml:"fill,attr,omitempty"`
	ThemeFillTint  *string            `xml:"themeFillTint,attr,omitempty"`
	ThemeFillShade *string            `xml:"themeFillShade,attr,omitempty"`
}

func (s Shading) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:shd"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(s.Val)})

	if s.Color != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:color"}, Value: *s.Color})
	}
	if s.ThemeColor != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeColor"}, Value: string(*s.ThemeColor)})
	}
	if s.ThemeFill != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeFill"}, Value: string(*s.ThemeFill)})
	}
	if s.ThemeTint != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeTint"}, Value: *s.ThemeTint})
	}
	if s.ThemeShade != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeShade"}, Value: *s.ThemeShade})
	}
	if s.Fill != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:fill"}, Value: *s.Fill})
	}
	if s.ThemeFillTint != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeFillTint"}, Value: *s.ThemeFillTint})
	}
	if s.ThemeFillShade != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeFillShade"}, Value: *s.ThemeFillShade})
	}

	return e.EncodeElement("", start)
}

// DefaultShading creates a new Shading with default values.
func DefaultShading() *Shading {
	color := "auto"
	fill := "FFFFFF"
	return &Shading{
		Val:   stypes.ShdClear,
		Color: &color,
		Fill:  &fill,
	}
}

// NewShading creates a new Shading.
func NewShading() *Shading {
	return DefaultShading()
}

// Color sets the color for the shading.
func (s *Shading) SetColor(color string) *Shading {
	s.Color = &color
	return s
}

// Fill sets the fill for the shading.
func (s *Shading) SetFill(fill string) *Shading {
	s.Fill = &fill
	return s
}

// ShadingType sets the shading type for the shading.
func (s *Shading) SetShadingType(shdType stypes.Shading) *Shading {
	s.Val = shdType
	return s
}
