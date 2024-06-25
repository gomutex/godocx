package table

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/stypes"
)

// Table Cell Margin Defaults
type CellMargins struct {
	// 1. Table Cell Top Margin Default
	Top *ctypes.TableWidth `xml:"top,omitempty"`

	// 2. Table Cell Left Margin Default
	Left *ctypes.TableWidth `xml:"left,omitempty"`

	// 3. Table Cell Bottom Margin Default
	Bottom *ctypes.TableWidth `xml:"bottom,omitempty"`

	// 4. Table Cell Right Margin Default
	Right *ctypes.TableWidth `xml:"right,omitempty"`
}

func DefaultCellMargins() CellMargins {
	return CellMargins{}
}

func (tcm CellMargins) Margin(top, left, bottom, right int) CellMargins {
	tcm.Top = ctypes.NewTableWidth(top, stypes.TableWidthDxa)
	tcm.Left = ctypes.NewTableWidth(left, stypes.TableWidthDxa)
	tcm.Bottom = ctypes.NewTableWidth(bottom, stypes.TableWidthDxa)
	tcm.Right = ctypes.NewTableWidth(right, stypes.TableWidthDxa)
	return tcm
}

func (tcm CellMargins) MarginTop(v int, t stypes.TableWidth) CellMargins {
	tcm.Top = ctypes.NewTableWidth(v, t)
	return tcm
}

func (tcm CellMargins) MarginRight(v int, t stypes.TableWidth) CellMargins {
	tcm.Right = ctypes.NewTableWidth(v, t)
	return tcm
}

func (tcm CellMargins) MarginLeft(v int, t stypes.TableWidth) CellMargins {
	tcm.Left = ctypes.NewTableWidth(v, t)
	return tcm
}

func (tcm CellMargins) MarginBottom(v int, t stypes.TableWidth) CellMargins {
	tcm.Bottom = ctypes.NewTableWidth(v, t)
	return tcm
}

func (tcm *CellMargins) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {

	if err = e.EncodeToken(start); err != nil {
		return err
	}

	// 1. Top
	if tcm.Top != nil {
		if err = tcm.Top.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:top"}}); err != nil {
			return err
		}
	}

	// 2. Left
	if tcm.Left != nil {
		if err = tcm.Left.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:left"}}); err != nil {
			return err
		}
	}

	// 3. Bottom
	if tcm.Bottom != nil {
		if err = tcm.Bottom.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:bottom"}}); err != nil {
			return err
		}
	}

	// 4. Right
	if tcm.Right != nil {
		if err = tcm.Right.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:right"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
