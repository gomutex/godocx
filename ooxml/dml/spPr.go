package dml

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
	bwMode         *string
	TransformGroup *TransformGroup
	PresetGeometry *PresetGeometry
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

	if p.bwMode != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "bwMode"}, Value: *p.bwMode})
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

func (p *PicShapeProp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	for _, a := range start.Attr {
		switch a.Name.Local {
		case "bwMode":
			p.bwMode = &a.Value
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
			case "xfrm":
				tfg := &TransformGroup{}
				if err := d.DecodeElement(tfg, &elem); err != nil {
					return err
				}
				p.TransformGroup = tfg
			case "prstGeom":
				pg := &PresetGeometry{}
				if err := d.DecodeElement(pg, &elem); err != nil {
					return err
				}
				p.PresetGeometry = pg
			}
		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
}
