package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// Run Fonts
type RunFonts struct {
	Hint          stypes.FontTypeHint `xml:"hint,attr,omitempty"`
	Ascii         string              `xml:"ascii,attr,omitempty"`
	HAnsi         string              `xml:"hAnsi,attr,omitempty"`
	EastAsia      string              `xml:"eastAsia,attr,omitempty"`
	CS            string              `xml:"cs,attr,omitempty"`
	AsciiTheme    stypes.ThemeFont    `xml:"asciiTheme,attr,omitempty"`
	HAnsiTheme    stypes.ThemeFont    `xml:"hAnsiTheme,attr,omitempty"`
	EastAsiaTheme stypes.ThemeFont    `xml:"eastAsiaTheme,attr,omitempty"`
	CSTheme       stypes.ThemeFont    `xml:"cstheme,attr,omitempty"`
}

func (rf RunFonts) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:rFonts"
	// start.Attr = []xml.Attr{}

	if rf.EastAsia != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:eastAsia"}, Value: rf.EastAsia})
	}

	if rf.Hint != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:hint"}, Value: string(rf.Hint)})
	}

	if rf.Ascii != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:ascii"}, Value: rf.Ascii})
	}

	if rf.HAnsi != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:hAnsi"}, Value: rf.HAnsi})
	}

	if rf.CS != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:cs"}, Value: rf.CS})
	}

	if rf.AsciiTheme != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:asciiTheme"}, Value: string(rf.AsciiTheme)})
	}

	if rf.HAnsiTheme != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:hAnsiTheme"}, Value: string(rf.HAnsiTheme)})
	}

	if rf.EastAsiaTheme != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:eastAsiaTheme"}, Value: string(rf.EastAsiaTheme)})
	}

	if rf.CSTheme != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:cstheme"}, Value: string(rf.CSTheme)})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
