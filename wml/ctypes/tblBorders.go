package ctypes

import (
	"encoding/xml"
)

type TableBorders struct {
	// 1. Table Top Border
	Top *Border `xml:"top,omitempty"`

	// 2. Table Left Border
	Left *Border `xml:"left,omitempty"`

	// 3. Table Bottom Border
	Bottom *Border `xml:"bottom,omitempty"`

	// 4. Table Right Border
	Right *Border `xml:"right,omitempty"`

	// 5. Table Inside Horizontal Edges Border
	InsideH *Border `xml:"insideH,omitempty"`
	// 6. Table Inside Vertical Edges Border
	InsideV *Border `xml:"insideV,omitempty"`
}

func DefaultTableBorders() *TableBorders {
	return &TableBorders{}
}

func (t TableBorders) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:tblBorders"

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// 1. Top
	if t.Top != nil {
		if err := t.Top.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:top"},
		}); err != nil {
			return err
		}
	}

	// 2. Left
	if t.Left != nil {
		if err := t.Left.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:left"},
		}); err != nil {
			return err
		}
	}

	// 3. Bottom
	if t.Bottom != nil {
		if err := t.Bottom.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:bottom"},
		}); err != nil {
			return err
		}
	}

	// 4. Right
	if t.Right != nil {
		if err := t.Right.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:right"},
		}); err != nil {
			return err
		}
	}

	// 5. insideH
	if t.InsideH != nil {
		if err := t.InsideH.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:insideH"},
		}); err != nil {
			return err
		}
	}

	// 6. InsideV
	if t.InsideV != nil {
		if err := t.InsideV.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:insideV"},
		}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})

}
