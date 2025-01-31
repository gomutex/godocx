package ctypes

import (
	"encoding/xml"
	"strconv"

	"github.com/gomutex/godocx/wml/stypes"
)

type Border struct {
	Val        stypes.BorderStyle `xml:"val,attr"`
	Color      *string            `xml:"color,attr,omitempty"`
	ThemeColor *stypes.ThemeColor `xml:"themeColor,attr,omitempty"`
	ThemeTint  *string            `xml:"themeTint,attr,omitempty"`
	ThemeShade *string            `xml:"themeShade,attr,omitempty"`
	Space      *string            `xml:"space,attr,omitempty"`
	Shadow     *stypes.OnOff      `xml:"shadow,attr,omitempty"`
	Frame      *stypes.OnOff      `xml:"frame,attr,omitempty"`
	Size       *int               `xml:"sz,attr,omitempty"`
}

/*
new cell border
@param style: border style
@param color: border color; @see github.com/gomutex/godocx/wml/color
@param space: the gap between the border and the cell content
@param size:  border size
*/
func NewCellBorder(style stypes.BorderStyle, color string, space string, size int) *Border {
	if size < 0 {
		size = 0
	}
	return &Border{
		Val:   style,
		Color: &color,
		Space: &space,
		Size:  &size,
	}
}

func (t Border) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(t.Val)})

	if t.Color != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:color"}, Value: *t.Color})
	}
	if t.ThemeColor != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeColor"}, Value: string(*t.ThemeColor)})
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

	if t.Size != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:sz"}, Value: strconv.Itoa(*t.Size)})
	}

	return e.EncodeElement("", start)
}
