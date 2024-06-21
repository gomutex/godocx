package dml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/gomutex/godocx/dml/dmlct"
)

type Inline struct {
	/// Specifies the minimum distance which shall be maintained between the top edge of this drawing object and any subsequent text within the document when this graphical object is displayed within the document's contents.,
	/// The distance shall be measured in EMUs (English Mektric Units).,
	//Distance From Text on Top Edge
	DistT *uint `xml:"distT,attr,omitempty"`

	//Distance From Text on Bottom Edge
	DistB *uint `xml:"distB,attr,omitempty"`

	//Distance From Text on Left Edge
	DistL *uint `xml:"distL,attr,omitempty"`

	//Distance From Text on Right Edge
	DistR *uint `xml:"distR,attr,omitempty"`

	// Child elements:
	// 1. Drawing Object Size
	Extent dmlct.PSize2D `xml:"extent,omitempty"`

	// 2. Inline Wrapping Extent
	EffectExtent *EffectExtent `xml:"effectExtent,omitempty"`

	// 3. Drawing Object Non-Visual Properties
	DocProp DocProp `xml:"docPr,omitempty"`

	//4.Common DrawingML Non-Visual Properties
	CNvGraphicFramePr *NonVisualGraphicFrameProp `xml:"cNvGraphicFramePr,omitempty"`

	//5.Graphic Object
	Graphic Graphic `xml:"graphic,omitempty"`
}

func NewInline() *Inline {
	return &Inline{}
}

func (i *Inline) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:inline"
	start.Attr = []xml.Attr{
		// {Name: xml.Name{Local: "xmlns:a"}, Value: constants.DrawingMLMainNS},
		// {Name: xml.Name{Local: "xmlns:pic"}, Value: constants.DrawingMLPicNS},
	}

	if i.DistT != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distT"}, Value: strconv.FormatUint(uint64(*i.DistT), 10)})
	}

	if i.DistB != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distB"}, Value: strconv.FormatUint(uint64(*i.DistB), 10)})
	}

	if i.DistL != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distL"}, Value: strconv.FormatUint(uint64(*i.DistL), 10)})
	}

	if i.DistR != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distR"}, Value: strconv.FormatUint(uint64(*i.DistR), 10)})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	// 1.Extent
	if err := i.Extent.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "wp:extent"}}); err != nil {
		return fmt.Errorf("marshalling Extent: %w", err)
	}

	// 2. EffectExtent
	if i.EffectExtent != nil {

		if err := i.EffectExtent.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "wp:effectExtent"}}); err != nil {
			return fmt.Errorf("EffectExtent: %v", err)
		}
	}

	// 3. docPr
	if err = i.DocProp.MarshalXML(e, xml.StartElement{}); err != nil {
		return fmt.Errorf("marshalling DocProp: %w", err)
	}

	// 4. cNvGraphicFramePr
	if i.CNvGraphicFramePr != nil {
		if err = i.CNvGraphicFramePr.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("marshalling cNvGraphicFramePr: %w", err)
		}
	}

	// 5. graphic
	if err = i.Graphic.MarshalXML(e, xml.StartElement{}); err != nil {
		return fmt.Errorf("marshalling Graphic: %w", err)
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
