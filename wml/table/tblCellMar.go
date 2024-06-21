package table

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/stypes"
)

// Table Cell Margin Defaults
type TableCellMargins struct {
	// 1. Table Cell Top Margin Default
	Top *ctypes.TableWidth `xml:"top,omitempty"`

	// 2. Table Cell Left Margin Default
	Left *ctypes.TableWidth `xml:"left,omitempty"`

	// 3. Table Cell Bottom Margin Default
	Bottom *ctypes.TableWidth `xml:"bottom,omitempty"`

	// 4. Table Cell Right Margin Default
	Right *ctypes.TableWidth `xml:"right,omitempty"`
}

func DefaultTableCellMargins() TableCellMargins {
	return TableCellMargins{}
}

func (tcm TableCellMargins) Margin(top, left, bottom, right int) TableCellMargins {
	tcm.Top = ctypes.NewTableWidth(top, stypes.TableWidthDxa)
	tcm.Left = ctypes.NewTableWidth(left, stypes.TableWidthDxa)
	tcm.Bottom = ctypes.NewTableWidth(bottom, stypes.TableWidthDxa)
	tcm.Right = ctypes.NewTableWidth(right, stypes.TableWidthDxa)
	return tcm
}

func (tcm TableCellMargins) MarginTop(v int, t stypes.TableWidth) TableCellMargins {
	tcm.Top = ctypes.NewTableWidth(v, t)
	return tcm
}

func (tcm TableCellMargins) MarginRight(v int, t stypes.TableWidth) TableCellMargins {
	tcm.Right = ctypes.NewTableWidth(v, t)
	return tcm
}

func (tcm TableCellMargins) MarginLeft(v int, t stypes.TableWidth) TableCellMargins {
	tcm.Left = ctypes.NewTableWidth(v, t)
	return tcm
}

func (tcm TableCellMargins) MarginBottom(v int, t stypes.TableWidth) TableCellMargins {
	tcm.Bottom = ctypes.NewTableWidth(v, t)
	return tcm
}

func (tcm *TableCellMargins) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	elem := xml.StartElement{Name: xml.Name{Local: "w:tblCellMar"}}

	if err = e.EncodeToken(elem); err != nil {
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

	return e.EncodeToken(elem.End())
}
