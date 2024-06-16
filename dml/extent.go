package dml

import (
	"encoding/xml"
	"strconv"

	"github.com/gomutex/godocx/common/units"
)

type Extent struct {
	Width  uint64 `xml:"cx,attr,omitempty"`
	Height uint64 `xml:"cy,attr,omitempty"`
}

func NewExtent(width units.Emu, height units.Emu) *Extent {
	return &Extent{
		Height: uint64(height),
		Width:  uint64(width),
	}
}

func (x *Extent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cx"}, Value: strconv.FormatUint(x.Width, 10)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cy"}, Value: strconv.FormatUint(x.Height, 10)})

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
