package dml

import (
	"encoding/xml"
	"strconv"

	"github.com/gomutex/godocx/common/constants"
	"github.com/gomutex/godocx/common/units"
)

type Pic struct {
	// ID string

	NonVisualPicProp *NonVisualPicProp
	BlipFill         *BlipFill
	PicShapeProp     *PicShapeProp
}

func NewPic(rID string, width units.Emu, height units.Emu) *Pic {
	shapeProp := NewPicShapeProp(
		WithTransformGroup(
			WithTFExtent(width, height),
		),
	)

	return &Pic{
		BlipFill:     NewBlipFill(rID),
		PicShapeProp: shapeProp,
	}
}

func (p *Pic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "pic:pic"

	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "xmlns:pic"}, Value: constants.DrawingMLPicNS},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if p.NonVisualPicProp != nil {
		if err := e.EncodeElement(p.NonVisualPicProp, xml.StartElement{Name: xml.Name{Local: "pic:nvPicPr"}}); err != nil {
			return err
		}
	}

	if p.BlipFill != nil {
		if err := e.EncodeElement(p.BlipFill, xml.StartElement{Name: xml.Name{Local: "pic:blipFill"}}); err != nil {
			return err
		}
	}

	if p.PicShapeProp != nil {
		if err := e.EncodeElement(p.PicShapeProp, xml.StartElement{Name: xml.Name{Local: "pic:spPr"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (p *Pic) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "nvPicPr":
				nvp := &NonVisualPicProp{}
				if err := d.DecodeElement(nvp, &elem); err != nil {
					return err
				}
				p.NonVisualPicProp = nvp
			case "blipFill":
				blipFill := &BlipFill{}
				if err := d.DecodeElement(blipFill, &elem); err != nil {
					return err
				}
				p.BlipFill = blipFill
			case "spPr":
				psp := &PicShapeProp{}
				if err := d.DecodeElement(psp, &elem); err != nil {
					return err
				}
				p.PicShapeProp = psp
			}
		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
}

type BlipFill struct {
	Blip    *Blip
	Stretch *Stretch
}

// NewBlipFill creates a new BlipFill with the given relationship ID (rID)
// The rID is used to reference the image in the presentation.
func NewBlipFill(rID string) *BlipFill {
	return &BlipFill{
		Blip: &Blip{
			EmbedID: rID,
		},
	}
}

func (b *BlipFill) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "pic:blipFill"

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if b.Blip != nil {
		if err := e.EncodeElement(b.Blip, xml.StartElement{Name: xml.Name{Local: "a:blip"}}); err != nil {
			return err
		}
	}

	if b.Stretch != nil {
		if err := e.EncodeElement(b.Stretch, xml.StartElement{Name: xml.Name{Local: "a:stretch"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (b *BlipFill) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "blip":
				blip := &Blip{}
				if err := d.DecodeElement(blip, &elem); err != nil {
					return err
				}
				b.Blip = blip
			case "stretch":
				stretch := &Stretch{}
				if err := d.DecodeElement(stretch, &elem); err != nil {
					return err
				}
				b.Stretch = stretch
			}

		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
}

// Binary large image or picture
type Blip struct {
	EmbedID string
}

func (b *Blip) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:blip"

	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "r:embed"}, Value: b.EmbedID},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (b *Blip) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "embed" {
			b.EmbedID = a.Value
		}
	}

	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			}
		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
}

type Stretch struct {
	FillRect *FillRect
}

func (s *Stretch) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:stretch"

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if s.FillRect != nil {
		if err := e.EncodeElement(s.FillRect, xml.StartElement{Name: xml.Name{Local: "a:fillRect"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (s *Stretch) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "fillRect":
				fr := &FillRect{}
				if err := d.DecodeElement(fr, &elem); err != nil {
					return err
				}
				s.FillRect = fr
			}
		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
}

type FillRect struct{}

func (f *FillRect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:fillRect"

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (f *FillRect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			}
		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
}

type TransformGroup struct {
	Extent *Extent
	Offset *Offset
}

type TFGroupOption func(*TransformGroup)

func NewTransformGroup(options ...TFGroupOption) *TransformGroup {
	tf := &TransformGroup{}

	for _, opt := range options {
		opt(tf)
	}

	return tf
}

func WithTFExtent(width units.Emu, height units.Emu) TFGroupOption {
	return func(tf *TransformGroup) {
		tf.Extent = NewExtent(width, height)
	}
}

func (t *TransformGroup) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:xfrm"

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if t.Offset != nil {
		if err := e.EncodeElement(t.Offset, xml.StartElement{Name: xml.Name{Local: "a:off"}}); err != nil {
			return err
		}
	}

	if t.Extent != nil {
		if err := e.EncodeElement(t.Extent, xml.StartElement{Name: xml.Name{Local: "a:ext"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (t *TransformGroup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "off":
				off := &Offset{}
				if err := d.DecodeElement(off, &elem); err != nil {
					return err
				}
				t.Offset = off
			case "ext":
				ex := &Extent{}
				if err := d.DecodeElement(ex, &elem); err != nil {
					return err
				}
				t.Extent = ex
			}
		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
}

type Offset struct {
	X uint64
	Y uint64
}

func (o *Offset) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:off"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "x"}, Value: strconv.FormatUint(o.X, 10)},
		{Name: xml.Name{Local: "y"}, Value: strconv.FormatUint(o.Y, 10)},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (o *Offset) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "x" {
			x, err := strconv.ParseUint(a.Value, 10, 32)
			if err != nil {
				return nil
			}
			o.X = x
		} else if a.Name.Local == "y" {
			y, err := strconv.ParseUint(a.Value, 10, 32)
			if err != nil {
				return nil
			}
			o.Y = y
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

type PresetGeometry struct {
	Preset       string
	AdjustValues *AdjustValues
}

func (p *PresetGeometry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:prstGeom"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "prst"}, Value: p.Preset},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if p.AdjustValues != nil {
		if err := e.EncodeElement(p.AdjustValues, xml.StartElement{Name: xml.Name{Local: "a:avLst"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (p *PresetGeometry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "prst" {
			p.Preset = a.Value
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
			case "avLst":
				avList := &AdjustValues{}
				if err := d.DecodeElement(avList, &elem); err != nil {
					return err
				}
				p.AdjustValues = avList
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
