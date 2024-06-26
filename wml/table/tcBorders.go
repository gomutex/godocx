package table

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/ctypes"
)

// Table Cell Borders
type CellBorders struct {
	// 1. Table Cell Top Border
	Top *ctypes.Border `xml:"top,omitempty"`

	// 2. Table Cell Left Border
	Left *ctypes.Border `xml:"left,omitempty"`

	// 3. Table Cell Bottom Border
	Bottom *ctypes.Border `xml:"bottom,omitempty"`

	// 4. Table Cell Right Border
	Right *ctypes.Border `xml:"right,omitempty"`

	// 5. Table Cell Inside Horizontal Edges Border
	InsideH *ctypes.Border `xml:"insideH,omitempty"`

	// 6. Table Cell Inside Vertical Edges Border
	InsideV *ctypes.Border `xml:"insideV,omitempty"`

	// 7. Table Cell Top Left to Bottom Right Diagonal Border
	TL2BR *ctypes.Border `xml:"tl2br,omitempty"`

	// 8. Table Cell Top Right to Bottom Left Diagonal Border
	TR2BL *ctypes.Border `xml:"tr2bl,omitempty"`
}

func DefaultCellBorders() *CellBorders {
	return &CellBorders{}
}

func (t *CellBorders) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:tcBorders"

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

	// 7. Table Cell Top Left to Bottom Right Diagonal Border
	if t.TL2BR != nil {
		if err := t.TL2BR.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tl2br"},
		}); err != nil {
			return err
		}
	}

	// 8. Table Cell Top Left to Bottom Right Diagonal Border
	if t.TR2BL != nil {
		if err := t.TR2BL.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tr2bl"},
		}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})

}
