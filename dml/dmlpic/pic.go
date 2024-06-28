package dmlpic

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/gomutex/godocx/common/constants"
	"github.com/gomutex/godocx/common/units"
	"github.com/gomutex/godocx/dml/dmlct"
	"github.com/gomutex/godocx/dml/geom"
	"github.com/gomutex/godocx/dml/shapes"
)

type Pic struct {
	// 1. Non-Visual Picture Properties
	NonVisualPicProp NonVisualPicProp `xml:"nvPicPr,omitempty"`

	// 2.Picture Fill
	BlipFill BlipFill `xml:"blipFill,omitempty"`

	// 3.Shape Properties
	PicShapeProp PicShapeProp `xml:"spPr,omitempty"`
}

func NewPic(rID string, imgCount uint, width units.Emu, height units.Emu) *Pic {
	shapeProp := NewPicShapeProp(
		WithTransformGroup(
			WithTFExtent(width, height),
		),
		WithPrstGeom("rect"),
	)

	nvPicProp := DefaultNVPicProp(imgCount, fmt.Sprintf("Image%v", imgCount))

	blipFill := NewBlipFill(rID)

	blipFill.FillModeProps = FillModeProps{
		Stretch: &shapes.Stretch{
			FillRect: &dmlct.RelativeRect{},
		},
	}

	return &Pic{
		BlipFill:         blipFill,
		NonVisualPicProp: nvPicProp,
		PicShapeProp:     *shapeProp,
	}
}

func (p Pic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "pic:pic"

	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "xmlns:pic"}, Value: constants.DrawingMLPicNS},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	// 1. nvPicPr
	if err = p.NonVisualPicProp.MarshalXML(e, xml.StartElement{
		Name: xml.Name{Local: "pic:nvPicPr"},
	}); err != nil {
		return fmt.Errorf("marshalling NonVisualPicProp: %w", err)
	}

	// 2. BlipFill
	if err = p.BlipFill.MarshalXML(e, xml.StartElement{
		Name: xml.Name{Local: "pic:blipFill"},
	}); err != nil {
		return fmt.Errorf("marshalling BlipFill: %w", err)
	}

	// 3. spPr
	if err = p.PicShapeProp.MarshalXML(e, xml.StartElement{
		Name: xml.Name{Local: "pic:spPr"},
	}); err != nil {
		return fmt.Errorf("marshalling PicShapeProp: %w", err)
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

type TransformGroup struct {
	Extent *dmlct.PSize2D `xml:"ext,omitempty"`
	Offset *Offset        `xml:"off,omitempty"`
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
		tf.Extent = dmlct.NewPostvSz2D(width, height)
	}
}

func (t TransformGroup) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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

func (o Offset) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	Preset       string             `xml:"prst,attr,omitempty"`
	AdjustValues *geom.AdjustValues `xml:"avLst,omitempty"`
}

func NewPresetGeom(preset string) *PresetGeometry {
	return &PresetGeometry{
		Preset: preset,
	}
}

func (p PresetGeometry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:prstGeom"
	start.Attr = []xml.Attr{}

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "prst"}, Value: p.Preset})

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
