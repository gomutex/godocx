package dml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type Inline struct {
	/// Specifies the minimum distance which shall be maintained between the top edge of this drawing object and any subsequent text within the document when this graphical object is displayed within the document's contents.,
	/// The distance shall be measured in EMUs (English Mektric Units).,
	DistTAttr *uint `xml:"distT,attr,omitempty"`
	DistBAttr *uint `xml:"distB,attr,omitempty"`
	DistLAttr *uint `xml:"distL,attr,omitempty"`
	DistRAttr *uint `xml:"distR,attr,omitempty"`

	// Child elements:
	Extent            *Extent                    `xml:"extent,omitempty"`
	DocProp           *DocProp                   `xml:"docPr,omitempty"`
	Graphic           *Graphic                   `xml:"graphic,omitempty"`
	CNvGraphicFramePr *NonVisualGraphicFrameProp `xml:"cNvGraphicFramePr,omitempty"`
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

	if i.DistTAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distT"}, Value: strconv.FormatUint(uint64(*i.DistTAttr), 10)})
	}

	if i.DistBAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distB"}, Value: strconv.FormatUint(uint64(*i.DistBAttr), 10)})
	}

	if i.DistLAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distL"}, Value: strconv.FormatUint(uint64(*i.DistLAttr), 10)})
	}

	if i.DistRAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distR"}, Value: strconv.FormatUint(uint64(*i.DistRAttr), 10)})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if i.Extent != nil {
		if err := i.Extent.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "wp:extent"}}); err != nil {
			return fmt.Errorf("marshalling Extent: %w", err)
		}
	}

	if i.DocProp != nil {
		if err = i.DocProp.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("marshalling DocProp: %w", err)
		}
	}

	if i.CNvGraphicFramePr != nil {
		if err = i.CNvGraphicFramePr.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("marshalling cNvGraphicFramePr: %w", err)
		}
	}

	if i.Graphic != nil {
		if err = i.Graphic.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("marshalling Graphic: %w", err)
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
