package dmlct

import (
	"encoding/xml"
	"strconv"

	"github.com/gomutex/godocx/common/units"
)

// Complex Type: CT_PositiveSize2D
type PSize2D struct {
	Width  uint64 `xml:"cx,attr,omitempty"`
	Height uint64 `xml:"cy,attr,omitempty"`
}

func NewPostvSz2D(width units.Emu, height units.Emu) *PSize2D {
	return &PSize2D{
		Height: uint64(height),
		Width:  uint64(width),
	}
}

func (p PSize2D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cx"}, Value: strconv.FormatUint(p.Width, 10)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cy"}, Value: strconv.FormatUint(p.Height, 10)})

	return e.EncodeElement("", start)
}
