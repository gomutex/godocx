package dml

import (
	"encoding/xml"

	"github.com/gomutex/godocx/common/constants"
	"github.com/gomutex/godocx/dml/dmlpic"
)

type Graphic struct {
	Data *GraphicData `xml:"graphicData,omitempty"`
}

func NewGraphic(data *GraphicData) *Graphic {
	return &Graphic{Data: data}
}

func DefaultGraphic() *Graphic {
	return &Graphic{}
}

type GraphicData struct {
	URI string      `xml:"uri,attr,omitempty"`
	Pic *dmlpic.Pic `xml:"pic,omitempty"`
}

func NewPicGraphic(pic *dmlpic.Pic) *Graphic {
	return &Graphic{
		Data: &GraphicData{
			URI: constants.DrawingMLPicNS,
			Pic: pic,
		},
	}
}

func (g Graphic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:graphic"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "xmlns:a"}, Value: constants.DrawingMLMainNS},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if g.Data != nil {
		if err = g.Data.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (gd GraphicData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:graphicData"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "uri"}, Value: constants.DrawingMLPicNS},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if gd.Pic != nil {
		if err := e.EncodeElement(gd.Pic, xml.StartElement{Name: xml.Name{Local: "pic:pic"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
