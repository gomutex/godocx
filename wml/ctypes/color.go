package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// Color represents the color of a text or element.
type Color struct {
	//Run Content Color
	Val string `xml:"val,attr"`

	//Run Content Theme Color
	ThemeColor *stypes.ThemeColor `xml:"themeColor,attr,omitempty"`

	//Run Content Theme Color Tint
	ThemeTint *string `xml:"themeTint,attr,omitempty"`

	//Run Content Theme Color Shade
	ThemeShade *string `xml:"themeShade,attr,omitempty"`
}

// NewColor creates a new Color instance with the specified color value.
func NewColor(value string) *Color {
	return &Color{Val: value}
}

// MarshalXML implements the xml.Marshaler interface for the Color type.
func (c Color) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:color"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: c.Val})

	if c.ThemeColor != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeColor"}, Value: string(*c.ThemeColor)})
	}

	if c.ThemeTint != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeTint"}, Value: *c.ThemeTint})
	}

	if c.ThemeShade != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeShade"}, Value: *c.ThemeShade})
	}

	return e.EncodeElement("", start)
}
