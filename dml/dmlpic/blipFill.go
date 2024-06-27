package dmlpic

import (
	"encoding/xml"
	"fmt"

	"github.com/gomutex/godocx/dml/dmlct"
	"github.com/gomutex/godocx/dml/shapes"
)

type BlipFill struct {
	// 1. Blip
	Blip *Blip `xml:"blip,omitempty"`

	//2.Source Rectangle
	SrcRect *dmlct.RelativeRect `xml:"srcRect,omitempty"`

	// 3. Choice of a:EG_FillModeProperties
	FillModeProps FillModeProps `xml:",any"`

	//Attributes:
	DPI          *uint32 `xml:"dpi,attr,omitempty"`          //DPI Setting
	RotWithShape *bool   `xml:"rotWithShape,attr,omitempty"` //Rotate With Shape
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

func (b BlipFill) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "pic:blipFill"

	if b.DPI != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dpi"}, Value: fmt.Sprintf("%d", *b.DPI)})
	}

	if b.RotWithShape != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rotWithShape"}, Value: fmt.Sprintf("%t", *b.RotWithShape)})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	// 1. Blip
	if b.Blip != nil {
		if err := b.Blip.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "a:blip"}}); err != nil {
			return err
		}
	}

	// 2. SrcRect
	if b.SrcRect != nil {
		if err = b.SrcRect.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "a:SrcRect"}}); err != nil {
			return err
		}
	}

	// 3. Choice: FillModProperties
	if err = b.FillModeProps.MarshalXML(e, xml.StartElement{}); err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

type FillModeProps struct {
	Stretch *shapes.Stretch `xml:"stretch,omitempty"`
	Tile    *shapes.Tile    `xml:"tile,omitempty"`
}

func (f FillModeProps) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	if f.Stretch != nil {
		return f.Stretch.MarshalXML(e, xml.StartElement{})
	}

	if f.Tile != nil {
		return f.Tile.MarshalXML(e, xml.StartElement{})
	}

	return nil
}
