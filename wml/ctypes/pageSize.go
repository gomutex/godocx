package ctypes

import (
	"encoding/xml"
	"strconv"

	"github.com/gomutex/godocx/wml/stypes"
)

// Page Size : w:pgSz
type PageSize struct {
	Width  *uint64           `xml:"w,attr,omitempty"`
	Height *uint64           `xml:"h,attr,omitempty"`
	Orient stypes.PageOrient `xml:"orient,attr,omitempty"`
	Code   *int              `xml:"code,attr,omitempty"`
}

func (p PageSize) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:pgSz"

	if p.Width != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:w"}, Value: strconv.FormatUint(*p.Width, 10)})
	}

	if p.Height != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:h"}, Value: strconv.FormatUint(*p.Height, 10)})
	}

	if p.Orient != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:orient"}, Value: string(p.Orient)})
	}

	if p.Code != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:code"}, Value: strconv.Itoa(*p.Code)})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
