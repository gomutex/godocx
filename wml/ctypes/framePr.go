package ctypes

import (
	"encoding/xml"
	"strconv"

	"github.com/gomutex/godocx/wml/stypes"
)

type FrameProp struct {
	Width  *int64 `xml:"w,attr,omitempty"`
	Height *int64 `xml:"h,attr,omitempty"`

	//Drop Cap Frame
	DropCap *stypes.DropCap `xml:"dropCap,attr,omitempty"`

	//Drop Cap Vertical Height in Lines
	Lines *int `xml:"lines,attr,omitempty"`

	//Frame Padding
	VSpace *int64 `xml:"vSpace,attr,omitempty"`
	HSpace *int64 `xml:"hSpace,attr,omitempty"`

	//Text Wrapping Around Frame
	Wrap *stypes.Wrap `xml:"wrap,attr,omitempty"`

	//Frame Horizontal Positioning Base
	HAnchor *stypes.Anchor `xml:"hAnchor,attr,omitempty"`

	//Frame Vertical Positioning Base
	VAnchor *stypes.Anchor `xml:"vAnchor,attr,omitempty"`

	//Absolute Horizontal Position
	AbsHPos *int `xml:"x,attr,omitempty"`

	//Absolute Vertical Position
	AbsVPos *int `xml:"y,attr,omitempty"`

	//Relative Horizontal Position
	XAlign *stypes.XAlign `xml:"xAlign,attr,omitempty"`

	//Relative Vertical Position
	YAlign *stypes.YAlign `xml:"yAlign,attr,omitempty"`

	//Frame Height Type
	HRule *stypes.HeightRule `xml:"hRule,attr,omitempty"`

	//Lock Frame Anchor to Paragraph
	AnchorLock *stypes.OnOff `xml:"anchorLock,attr,omitempty"`
}

func (f FrameProp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:framePr"

	start.Attr = []xml.Attr{}

	if f.Width != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:w"}, Value: strconv.FormatInt(*f.Width, 10)})
	}

	if f.Height != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:h"}, Value: strconv.FormatInt(*f.Height, 10)})
	}

	if f.DropCap != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:dropCap"}, Value: string(*f.DropCap)})
	}

	if f.Lines != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:lines"}, Value: strconv.Itoa(*f.Lines)})
	}

	if f.HSpace != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:hSpace"}, Value: strconv.FormatInt(*f.HSpace, 10)})
	}

	if f.VSpace != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:vSpace"}, Value: strconv.FormatInt(*f.VSpace, 10)})
	}

	if f.Wrap != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:wrap"}, Value: string(*f.Wrap)})
	}

	if f.HAnchor != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:hAnchor"}, Value: string(*f.HAnchor)})
	}

	if f.VAnchor != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:vAnchor"}, Value: string(*f.VAnchor)})
	}

	if f.AbsHPos != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:x"}, Value: strconv.Itoa(*f.AbsHPos)})
	}

	if f.AbsVPos != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:y"}, Value: strconv.Itoa(*f.AbsVPos)})
	}

	if f.XAlign != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:xAlign"}, Value: string(*f.XAlign)})
	}

	if f.YAlign != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:yAlign"}, Value: string(*f.YAlign)})
	}

	if f.HRule != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:hRule"}, Value: string(*f.HRule)})
	}

	if f.AnchorLock != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:anchorLock"}, Value: string(*f.AnchorLock)})
	}

	return e.EncodeElement("", start)

}
