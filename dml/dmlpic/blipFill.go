package dmlpic

import "encoding/xml"

type BlipFill struct {
	// 1. Blip
	Blip *Blip `xml:"blip,omitempty"`

	//2.Source Rectangle
	// TODO: Implement

	// 3. Choice of a:EG_FillModeProperties
	// TODO: Add the tile and group them into EG_FillModeProperties
	Stretch *Stretch `xml:"stretch,omitempty"`
}

// NewBlipFill creates a new BlipFill with the given relationship ID (rID)
// The rID is used to reference the image in the presentation.
func NewBlipFill(rID string) BlipFill {
	return BlipFill{
		Blip: &Blip{
			EmbedID: rID,
		},
	}
}

func (b *BlipFill) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "pic:blipFill"

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	// 1. Blip
	if b.Blip != nil {
		if err := e.EncodeElement(b.Blip, xml.StartElement{Name: xml.Name{Local: "a:blip"}}); err != nil {
			return err
		}
	}

	//2. TODO: Implement srcRect

	// 3.
	if b.Stretch != nil {
		if err := e.EncodeElement(b.Stretch, xml.StartElement{Name: xml.Name{Local: "a:stretch"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

type Stretch struct {
	FillRect *FillRect `xml:"fillRect,omitempty"`
}

func (s *Stretch) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:stretch"

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if s.FillRect != nil {
		if err := e.EncodeElement(s.FillRect, xml.StartElement{Name: xml.Name{Local: "a:fillRect"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
