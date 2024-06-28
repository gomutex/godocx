package dml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/gomutex/godocx/common/constants"
	"github.com/gomutex/godocx/dml/dmlct"
	"github.com/gomutex/godocx/dml/dmlst"
)

// This element specifies that the DrawingML object located at this position in the document is an inline object. Within a WordprocessingML document, drawing objects can exist in two states:
//
//• Inline - The drawing object is in line with the text, and affects the line height and layout of its line (like a character glyph of similar size).

type Inline struct {
	/// Specifies the minimum distance which shall be maintained between the top edge of this drawing object and any subsequent text within the document when this graphical object is displayed within the document's contents.,
	/// The distance shall be measured in EMUs (English Mektric Units).,
	//
	// NOTE!: As per http://www.datypic.com/sc/ooxml/e-wp_inline.html, Dist* attributes is optional
	// But MS Word requires them to be there

	//Distance From Text on Top Edge
	DistT uint `xml:"distT,attr,omitempty"`

	//Distance From Text on Bottom Edge
	DistB uint `xml:"distB,attr,omitempty"`

	//Distance From Text on Left Edge
	DistL uint `xml:"distL,attr,omitempty"`

	//Distance From Text on Right Edge
	DistR uint `xml:"distR,attr,omitempty"`

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

func NewInline(extent dmlct.PSize2D, docProp DocProp, graphic Graphic) Inline {
	return Inline{
		Extent:  extent,
		DocProp: docProp,
		Graphic: graphic,
		CNvGraphicFramePr: &NonVisualGraphicFrameProp{
			GraphicFrameLocks: &GraphicFrameLocks{
				NoChangeAspect: dmlst.NewOptBool(true),
			},
		},
	}
}

func (i Inline) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:inline"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "xmlns:a"}, Value: constants.DrawingMLMainNS},
		{Name: xml.Name{Local: "xmlns:pic"}, Value: constants.DrawingMLPicNS},
	}

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distT"}, Value: strconv.FormatUint(uint64(i.DistT), 10)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distB"}, Value: strconv.FormatUint(uint64(i.DistB), 10)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distL"}, Value: strconv.FormatUint(uint64(i.DistL), 10)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distR"}, Value: strconv.FormatUint(uint64(i.DistR), 10)})

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
