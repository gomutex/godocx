package dmlpic

import "encoding/xml"

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
	BwMode         *string         `xml:"bwMode,attr,omitempty"`
	TransformGroup *TransformGroup `xml:"xfrm,omitempty"`
	PresetGeometry *PresetGeometry `xml:"prstGeom,omitempty"`
}

type PicShapePropOption func(*PicShapeProp)

func WithTransformGroup(options ...TFGroupOption) PicShapePropOption {
	return func(p *PicShapeProp) {
		p.TransformGroup = NewTransformGroup(options...)
	}
}

func NewPicShapeProp(options ...PicShapePropOption) *PicShapeProp {
	p := &PicShapeProp{}

	for _, opt := range options {
		opt(p)
	}

	return p
}

func (p *PicShapeProp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "pic:spPr"

	if p.BwMode != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "bwMode"}, Value: *p.BwMode})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if p.TransformGroup != nil {
		if err := e.EncodeElement(p.TransformGroup, xml.StartElement{Name: xml.Name{Local: "a:xfrm"}}); err != nil {
			return err
		}
	}

	if p.PresetGeometry != nil {
		if err := e.EncodeElement(p.PresetGeometry, xml.StartElement{Name: xml.Name{Local: "a:prstGeom"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
