package docxrun

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/simpletypes"
)

type TextBorder struct {
	Val        simpletypes.BorderStyle `xml:"val,attr"`
	Color      *string                 `xml:"color,attr,omitempty"`
	ThemeColor *string                 `xml:"themeColor,attr,omitempty"`
	ThemeTint  *string                 `xml:"themeTint,attr,omitempty"`
	ThemeShade *string                 `xml:"themeShade,attr,omitempty"`
	Space      *string                 `xml:"space,attr,omitempty"`
	Shadow     *simpletypes.OnOff      `xml:"shadow,attr,omitempty"`
	Frame      *simpletypes.OnOff      `xml:"frame,attr,omitempty"`
}

func (t *TextBorder) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(t.Val)})

	if t.Color != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:color"}, Value: *t.Color})
	}
	if t.ThemeColor != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeColor"}, Value: *t.ThemeColor})
	}
	if t.ThemeTint != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeTint"}, Value: *t.ThemeTint})
	}
	if t.ThemeShade != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeShade"}, Value: *t.ThemeShade})
	}
	if t.Space != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:space"}, Value: *t.Space})
	}

	if t.Shadow != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:shadow"}, Value: string(*t.Shadow)})
	}

	if t.Frame != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:frame"}, Value: string(*t.Frame)})
	}

	return e.EncodeElement("", start)
}
