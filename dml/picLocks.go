package dml

import (
	"encoding/xml"

	"github.com/gomutex/godocx/types"
)

type PicLocks struct {
	DisallowShadowGrouping types.NullBool // noGrp
	NoSelect               types.NullBool // noSelect
	NoRot                  types.NullBool // noRot
	NoChangeAspect         types.NullBool // noChangeAspect
	NoMove                 types.NullBool // noMove
	NoResize               types.NullBool // noResize
	NoEditPoints           types.NullBool // noEditPoints
	NoAdjustHandles        types.NullBool // noAdjustHandles
	NoChangeArrowheads     types.NullBool // noChangeArrowheads
	NoChangeShapeType      types.NullBool // noChangeShapeType
	NoCrop                 types.NullBool // noCrop
}

func (p *PicLocks) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:picLocks"
	start.Attr = []xml.Attr{}

	if p.DisallowShadowGrouping.Valid {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noGrp"}, Value: p.DisallowShadowGrouping.ToStringFlag()})
	}

	if p.NoSelect.Valid {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noSelect"}, Value: p.NoSelect.ToStringFlag()})
	}
	if p.NoRot.Valid {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noRot"}, Value: p.NoRot.ToStringFlag()})
	}
	if p.NoChangeAspect.Valid {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noChangeAspect"}, Value: p.NoChangeAspect.ToStringFlag()})
	}
	if p.NoMove.Valid {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noMove"}, Value: p.NoMove.ToStringFlag()})
	}
	if p.NoResize.Valid {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noResize"}, Value: p.NoResize.ToStringFlag()})
	}
	if p.NoEditPoints.Valid {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noEditPoints"}, Value: p.NoEditPoints.ToStringFlag()})
	}
	if p.NoAdjustHandles.Valid {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noAdjustHandles"}, Value: p.NoAdjustHandles.ToStringFlag()})
	}
	if p.NoChangeArrowheads.Valid {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noChangeArrowheads"}, Value: p.NoChangeArrowheads.ToStringFlag()})
	}
	if p.NoChangeShapeType.Valid {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noChangeShapeType"}, Value: p.NoChangeShapeType.ToStringFlag()})
	}
	if p.NoCrop.Valid {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noCrop"}, Value: p.NoCrop.ToStringFlag()})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (p *PicLocks) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		switch a.Name.Local {
		case "noGrp":
			p.DisallowShadowGrouping = types.NullBoolFromStr(a.Value)
		case "noSelect":
			p.NoSelect = types.NullBoolFromStr(a.Value)
		case "noRot":
			p.NoRot = types.NullBoolFromStr(a.Value)
		case "noChangeAspect":
			p.NoChangeAspect = types.NullBoolFromStr(a.Value)
		case "noMove":
			p.NoMove = types.NullBoolFromStr(a.Value)
		case "noResize":
			p.NoResize = types.NullBoolFromStr(a.Value)
		case "noEditPoints":
			p.NoEditPoints = types.NullBoolFromStr(a.Value)
		case "noAdjustHandles":
			p.NoAdjustHandles = types.NullBoolFromStr(a.Value)
		case "noChangeArrowheads":
			p.NoChangeArrowheads = types.NullBoolFromStr(a.Value)
		case "noChangeShapeType":
			p.NoChangeShapeType = types.NullBoolFromStr(a.Value)
		case "noCrop":
			p.NoCrop = types.NullBoolFromStr(a.Value)
		}
	}

	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			if err := d.Skip(); err != nil {
				return err
			}
		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
}
