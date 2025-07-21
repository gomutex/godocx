package ctypes

import (
	"encoding/xml"
)

type Pict struct {
	Shape *Shape `xml:"shape,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface.
func (b Pict) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:pict"

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if b.Shape != nil {
		if err = b.Shape.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

type Shape struct {
	Type  string `xml:"type,attr,omitempty"`
	Style string `xml:"style,attr,omitempty"`

	ImageData *ImageData `xml:"imagedata,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface.
func (b Shape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "v:shape"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "type"}, Value: b.Type},
		{Name: xml.Name{Local: "style"}, Value: b.Style},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if b.ImageData != nil {
		if err := b.ImageData.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

type ImageData struct {
	RId   string `xml:"id,attr,omitempty"`
	Title string `xml:"title,attr,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface.
func (b ImageData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "v:imagedata"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "r:id"}, Value: b.RId},
		{Name: xml.Name{Local: "o:title"}, Value: b.Title},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
