package dml

import (
	"encoding/xml"

	"github.com/gomutex/godocx/common/constants"
)

type Inline struct {
	/// Specifies the minimum distance which shall be maintained between the top edge of this drawing object and any subsequent text within the document when this graphical object is displayed within the document's contents.,
	/// The distance shall be measured in EMUs (English Mektric Units).,
	DistTAttr int
	DistBAttr int
	DistLAttr int
	DistRAttr int

	// Child elements:
	Extent            *Extent
	DocProp           *DocProp
	Graphic           *Graphic
	cNvGraphicFramePr *NonVisualGraphicFrameProp
}

func NewInline() *Inline {
	return &Inline{}
}

func (i *Inline) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:inline"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "xmlns:a"}, Value: constants.DrawingMLMainNS},
		{Name: xml.Name{Local: "xmlns:pic"}, Value: constants.DrawingMLPicNS},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if i.Extent != nil {
		if err = e.EncodeElement(i.Extent, xml.StartElement{Name: xml.Name{Local: "wp:extent"}}); err != nil {
			return err
		}
	}

	if i.DocProp != nil {
		if err = e.EncodeElement(i.DocProp, xml.StartElement{Name: xml.Name{Local: "wp:docPr"}}); err != nil {
			return err
		}
	}

	if i.cNvGraphicFramePr != nil {
		if err = e.EncodeElement(i.cNvGraphicFramePr, xml.StartElement{Name: xml.Name{Local: "wp:cNvGraphicFramePr"}}); err != nil {
			return err
		}
	}

	if i.Graphic != nil {
		if err = i.Graphic.MarshalXML(e, start); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (i *Inline) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "extent":
				i.Extent = &Extent{}
				if err := decoder.DecodeElement(i.Extent, &elem); err != nil {
					return err
				}
			case "docPr":
				i.DocProp = &DocProp{}
				if err := decoder.DecodeElement(i.DocProp, &elem); err != nil {
					return err
				}
			case "cNvGraphicFramePr":
				i.cNvGraphicFramePr = &NonVisualGraphicFrameProp{}
				if err := decoder.DecodeElement(i.cNvGraphicFramePr, &elem); err != nil {
					return err
				}
			case "graphic":
				i.Graphic = DefaultGraphic()
				if err := decoder.DecodeElement(i.Graphic, &elem); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
}
