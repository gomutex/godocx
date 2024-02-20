package elements

import (
	"encoding/xml"

	"github.com/gomutex/godocx/constants"
)

type Graphic struct {
	Data *GraphicData
}

func DefaultGraphic() *Graphic {
	return &Graphic{}
}

type GraphicData struct {
	URI string
	Pic *Pic
}

func (g *Graphic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:graphic"
	start.Attr = []xml.Attr{}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if g.Data != nil {
		if err = g.Data.MarshalXML(e, start); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (g *Graphic) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "graphicData":
				gd := &GraphicData{}
				if err := decoder.DecodeElement(gd, &elem); err != nil {
					return err
				}
				g.Data = gd
			}
		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
}

func (gd *GraphicData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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

func (gd *GraphicData) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "uri" {
			gd.URI = a.Value
			break
		}
	}

	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "pic":
				pic := &Pic{}
				if err := decoder.DecodeElement(pic, &elem); err != nil {
					return err
				}
				gd.Pic = pic
			}
		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
}
