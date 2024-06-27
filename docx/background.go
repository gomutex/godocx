package docx

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// Specifies the background information for this document
//
// This background shall be displayed on all pages of the document, behind all other document content.
type Background struct {
	Color      *string            `xml:"color,attr,omitempty"`
	ThemeColor *stypes.ThemeColor `xml:"themeColor,attr,omitempty"`
	ThemeTint  *string            `xml:"themeTint,attr,omitempty"`
	ThemeShade *string            `xml:"themeShade,attr,omitempty"`
}

func NewBackground() *Background {
	return &Background{}
}
func (b Background) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:background"
	if b.Color != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:color"}, Value: *b.Color})
	}
	if b.ThemeColor != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeColor"}, Value: string(*b.ThemeColor)})
	}
	if b.ThemeTint != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeTint"}, Value: *b.ThemeTint})
	}
	if b.ThemeShade != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeShade"}, Value: *b.ThemeShade})
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	return e.EncodeToken(start.End())
}
