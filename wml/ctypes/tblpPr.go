package ctypes

import (
	"encoding/xml"
	"strconv"

	"github.com/gomutex/godocx/wml/stypes"
)

type FloatPos struct {
	LeftFromText   *uint64 `xml:"leftFromText,attr,omitempty"`
	RightFromText  *uint64 `xml:"rightFromText,attr,omitempty"`
	TopFromText    *uint64 `xml:"topFromText,attr,omitempty"`
	BottomFromText *uint64 `xml:"bottomFromText,attr,omitempty"`

	//Frame Horizontal Positioning Base
	HAnchor *stypes.Anchor `xml:"hAnchor,attr,omitempty"`

	//Frame Vertical Positioning Base
	VAnchor *stypes.Anchor `xml:"vAnchor,attr,omitempty"`

	//Relative Horizontal Alignment From Anchor
	XAlign *stypes.XAlign `xml:"tblpXSpec,attr,omitempty"`

	//Relative Vertical Alignment from Anchor
	YAlign *stypes.YAlign `xml:"tblpYSpec,attr,omitempty"`

	//Absolute Horizontal Distance From Anchor
	AbsXDist *int `xml:"tblpX,attr,omitempty"`

	// Absolute Vertical Distance From Anchor
	AbsYDist *int `xml:"tblpY,attr,omitempty"`
}

func (t FloatPos) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:tblpPr"

	if t.LeftFromText != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:leftFromText"}, Value: strconv.FormatUint(*t.LeftFromText, 10)})
	}
	if t.RightFromText != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rightFromText"}, Value: strconv.FormatUint(*t.RightFromText, 10)})
	}
	if t.TopFromText != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:topFromText"}, Value: strconv.FormatUint(*t.TopFromText, 10)})
	}
	if t.BottomFromText != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:bottomFromText"}, Value: strconv.FormatUint(*t.BottomFromText, 10)})
	}
	if t.HAnchor != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:hAnchor"}, Value: string(*t.HAnchor)})
	}
	if t.VAnchor != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:vAnchor"}, Value: string(*t.VAnchor)})
	}
	if t.XAlign != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:tblpXSpec"}, Value: string(*t.XAlign)})
	}
	if t.YAlign != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:tblpYSpec"}, Value: string(*t.YAlign)})
	}
	if t.AbsXDist != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:tblpX"}, Value: strconv.Itoa(*t.AbsXDist)})
	}
	if t.AbsYDist != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:tblpY"}, Value: strconv.Itoa(*t.AbsYDist)})
	}

	return e.EncodeElement("", start)
}
