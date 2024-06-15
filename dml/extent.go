package dml

import (
	"encoding/xml"
	"strconv"

	"github.com/gomutex/godocx/common/units"
)

type Extent struct {
	XMLName string
	Width   uint64 // cx
	Height  uint64 // cy
}

func NewExtent(width units.Emu, height units.Emu) *Extent {
	return &Extent{
		Height: uint64(height),
		Width:  uint64(width),
	}
}

func (x *Extent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// start.Name.Local = x.XMLName
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "cx"}, Value: strconv.FormatUint(x.Width, 10)},
		{Name: xml.Name{Local: "cy"}, Value: strconv.FormatUint(x.Height, 10)},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (x *Extent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		switch a.Name.Local {
		case "cx":
			cx, err := strconv.ParseUint(a.Value, 10, 32)
			if err != nil {
				return nil
			}
			x.Width = cx
		case "cy":
			cy, err := strconv.ParseUint(a.Value, 10, 32)
			if err != nil {
				return nil
			}
			x.Height = cy
		}
	}

	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
}
