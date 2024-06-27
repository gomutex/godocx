package ctypes

import (
	"encoding/xml"
	"strconv"
)

// FontSize represents the font size of a text or element.
type FontSize struct {
	Value uint64 `xml:"val,attr,omitempty"`
}

// NewFontSize creates a new FontSize with the specified font size value.
func NewFontSize(value uint64) *FontSize {
	return &FontSize{Value: value}
}

func (s FontSize) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:sz"
	start.Attr = []xml.Attr{{Name: xml.Name{Local: "w:val"}, Value: strconv.FormatUint(s.Value, 10)}}
	return e.EncodeElement("", start)
}

// FontSizeCs represents the font size of a text or element.
type FontSizeCS struct {
	Value uint64 `xml:"val,attr,omitempty"`
}

// NewFontSizeCs creates a new FontSizeCs with the specified font size value.
func NewFontSizeCS(value uint64) *FontSizeCS {
	return &FontSizeCS{Value: value}
}

func (s FontSizeCS) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:szCs"
	start.Attr = []xml.Attr{{Name: xml.Name{Local: "w:val"}, Value: strconv.FormatUint(s.Value, 10)}}
	return e.EncodeElement("", start)
}
