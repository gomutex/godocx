package dmlprops

import (
	"encoding/xml"

	"github.com/gomutex/godocx/dml/dmlst"
)

// Picture Locks
type PicLocks struct {
	DisallowShadowGrouping dmlst.OptBool `xml:"noGrp,attr,omitempty"`
	NoSelect               dmlst.OptBool `xml:"noSelect,attr,omitempty"`
	NoRot                  dmlst.OptBool `xml:"noRot,attr,omitempty"`
	NoChangeAspect         dmlst.OptBool `xml:"noChangeAspect,attr,omitempty"`
	NoMove                 dmlst.OptBool `xml:"noMove,attr,omitempty"`
	NoResize               dmlst.OptBool `xml:"noResize,attr,omitempty"`
	NoEditPoints           dmlst.OptBool `xml:"noEditPoints,attr,omitempty"`
	NoAdjustHandles        dmlst.OptBool `xml:"noAdjustHandles,attr,omitempty"`
	NoChangeArrowheads     dmlst.OptBool `xml:"noChangeArrowheads,attr,omitempty"`
	NoChangeShapeType      dmlst.OptBool `xml:"noChangeShapeType,attr,omitempty"`
	NoCrop                 dmlst.OptBool `xml:"noCrop,attr,omitempty"`
}

func DefaultPicLocks() *PicLocks {
	return &PicLocks{
		NoChangeAspect:     dmlst.NewOptBool(true),
		NoChangeArrowheads: dmlst.NewOptBool(true),
	}
}
func (p PicLocks) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
