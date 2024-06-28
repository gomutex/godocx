package dmlpic

import (
	"encoding/xml"
	"fmt"
)

const (
	BlackWhiteModeClr        = "clr"
	BlackWhiteModeAuto       = "auto"
	BlackWhiteModeGray       = "gray"
	BlackWhiteModeLtGray     = "ltGray"
	BlackWhiteModeInvGray    = "invGray"
	BlackWhiteModeGrayWhite  = "grayWhite"
	BlackWhiteModeBlackGray  = "blackGray"
	BlackWhiteModeBlackWhite = "blackWhite"
	BlackWhiteModeBlack      = "black"
	BlackWhiteModeWhite      = "white"
	BlackWhiteModeHidden     = "hidden"
)

type PicShapeProp struct {
	// -- Attributes --
	//Black and White Mode
	BwMode *string `xml:"bwMode,attr,omitempty"`

	// -- Child Elements --
	//1.2D Transform for Individual Objects
	TransformGroup *TransformGroup `xml:"xfrm,omitempty"`

	// 2. Choice
	//TODO: Modify it as Geometry choice
	PresetGeometry *PresetGeometry `xml:"prstGeom,omitempty"`

	//TODO: Remaining sequcence of elements
}

type PicShapePropOption func(*PicShapeProp)

func WithTransformGroup(options ...TFGroupOption) PicShapePropOption {
	return func(p *PicShapeProp) {
		p.TransformGroup = NewTransformGroup(options...)
	}
}

func WithPrstGeom(preset string) PicShapePropOption {
	return func(p *PicShapeProp) {
		p.PresetGeometry = NewPresetGeom(preset)
	}
}

func NewPicShapeProp(options ...PicShapePropOption) *PicShapeProp {
	p := &PicShapeProp{}

	for _, opt := range options {
		opt(p)
	}

	return p
}

func (p PicShapeProp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "pic:spPr"

	if p.BwMode != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "bwMode"}, Value: *p.BwMode})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	//1. Transform
	if p.TransformGroup != nil {
		if err = p.TransformGroup.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "a:xfrm"},
		}); err != nil {
			return fmt.Errorf("marshalling TransformGroup: %w", err)
		}
	}

	//2. Geometry
	if p.PresetGeometry != nil {

		if err = p.PresetGeometry.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "a:prstGeom"},
		}); err != nil {
			return fmt.Errorf("marshalling PresetGeometry: %w", err)
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
