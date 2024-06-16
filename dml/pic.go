package dml

import (
	"encoding/xml"
	"strconv"

	"github.com/gomutex/godocx/common/constants"
	"github.com/gomutex/godocx/common/units"
)

type Pic struct {
	// ID string
	NonVisualPicProp *NonVisualPicProp `xml:"nvPicPr,omitempty"`
	BlipFill         *BlipFill         `xml:"blipFill,omitempty"`
	PicShapeProp     *PicShapeProp     `xml:"spPr,omitempty"`
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

type BlipFill struct {
	Blip    *Blip    `xml:"blip,omitempty"`
	Stretch *Stretch `xml:"stretch,omitempty"`
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

// Binary large image or picture
type Blip struct {
	EmbedID string `xml:"embed,attr,omitempty"`
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

type Stretch struct {
	FillRect *FillRect `xml:"fillRect,omitempty"`
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

type FillRect struct{}

func (f *FillRect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:fillRect"

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

type TransformGroup struct {
	Extent *Extent `xml:"ext,omitempty"`
	Offset *Offset `xml:"off,omitempty"`
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

type Offset struct {
	X uint64 `xml:"x,attr,omitempty"`
	Y uint64 `xml:"y,attr,omitempty"`
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

type PresetGeometry struct {
	Preset       *string       `xml:"prst,attr,omitempty"`
	AdjustValues *AdjustValues `xml:"avLst,omitempty"`
}

func (p *PresetGeometry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:prstGeom"
	start.Attr = []xml.Attr{}

	if p.Preset != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "prst"}, Value: *p.Preset})
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
