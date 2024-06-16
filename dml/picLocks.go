package dml

import (
	"encoding/xml"

	"github.com/gomutex/godocx/types"
)

type PicLocks struct {
	DisallowShadowGrouping types.NullBool `xml:"noGrp,attr,omitempty"`
	NoSelect               types.NullBool `xml:"noSelect,attr,omitempty"`
	NoRot                  types.NullBool `xml:"noRot,attr,omitempty"`
	NoChangeAspect         types.NullBool `xml:"noChangeAspect,attr,omitempty"`
	NoMove                 types.NullBool `xml:"noMove,attr,omitempty"`
	NoResize               types.NullBool `xml:"noResize,attr,omitempty"`
	NoEditPoints           types.NullBool `xml:"noEditPoints,attr,omitempty"`
	NoAdjustHandles        types.NullBool `xml:"noAdjustHandles,attr,omitempty"`
	NoChangeArrowheads     types.NullBool `xml:"noChangeArrowheads,attr,omitempty"`
	NoChangeShapeType      types.NullBool `xml:"noChangeShapeType,attr,omitempty"`
	NoCrop                 types.NullBool `xml:"noCrop,attr,omitempty"`
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
