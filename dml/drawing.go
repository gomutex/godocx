package dml

import (
	"encoding/xml"

	"github.com/gomutex/godocx/common/constants"
)

type DrawingPositionType string

const (
	DrawingPositionAnchor DrawingPositionType = "wp:anchor"
	DrawingPositionInline DrawingPositionType = "wp:inline"
)

type Drawing struct {
	Inline []Inline  `xml:"inline,omitempty"`
	Anchor []*Anchor `xml:"anchor,omitempty"`
}

func (dr *Drawing) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
loop:
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name {
			case xml.Name{Space: constants.WMLDrawingNS, Local: "anchor"}:
				ar := NewAnchor()
				if err = d.DecodeElement(ar, &elem); err != nil {
					return err
				}

				dr.Anchor = append(dr.Anchor, ar)
			case xml.Name{Space: constants.WMLDrawingNS, Local: "inline"}:
				il := Inline{}
				if err = d.DecodeElement(&il, &elem); err != nil {
					return err
				}

				dr.Inline = append(dr.Inline, il)
			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break loop
		}
	}

	return nil
}

func (dr Drawing) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:drawing"

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for _, data := range dr.Anchor {
		if err = data.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	for _, data := range dr.Inline {
		if err = data.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
