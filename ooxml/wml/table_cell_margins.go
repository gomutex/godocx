package wml

import (
	"encoding/xml"
	"strconv"

	"github.com/gomutex/godocx/common/constants"
)

type CellMargin struct {
	Val       uint64
	WidthType WidthType
}

func (cm *CellMargin) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:w"}, Value: strconv.FormatUint(uint64(cm.Val), 10)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:type"}, Value: string(cm.WidthType)})

	return e.EncodeElement("", start)
}

func (cm *CellMargin) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "w":
			valueStr := attr.Value
			if valueStr != "" {
				value, err := strconv.ParseUint(valueStr, 10, 0)
				if err != nil {
					return err
				}
				cm.Val = value
			}
		case "type":
			cm.WidthType = WidthType(attr.Value)
		}
	}

	return d.Skip() // Skipping the inner content
}

func NewCellMargin(val uint64, widthType WidthType) *CellMargin {
	return &CellMargin{
		Val:       val,
		WidthType: widthType,
	}
}

func DefaultCellMargin() *CellMargin {
	return &CellMargin{}
}

type TableCellMargins struct {
	Top    *CellMargin
	Left   *CellMargin
	Bottom *CellMargin
	Right  *CellMargin
}

func DefaultTableCellMargins() TableCellMargins {
	return TableCellMargins{}
}

func (tcm TableCellMargins) Margin(top, right, bottom, left uint64) TableCellMargins {
	tcm.Top = NewCellMargin(top, WidthTypeDxa)
	tcm.Left = NewCellMargin(left, WidthTypeDxa)
	tcm.Bottom = NewCellMargin(bottom, WidthTypeDxa)
	tcm.Right = NewCellMargin(right, WidthTypeDxa)
	return tcm
}

func (tcm TableCellMargins) MarginTop(v uint64, t WidthType) TableCellMargins {
	tcm.Top = NewCellMargin(v, t)
	return tcm
}

func (tcm TableCellMargins) MarginRight(v uint64, t WidthType) TableCellMargins {
	tcm.Right = NewCellMargin(v, t)
	return tcm
}

func (tcm TableCellMargins) MarginLeft(v uint64, t WidthType) TableCellMargins {
	tcm.Left = NewCellMargin(v, t)
	return tcm
}

func (tcm TableCellMargins) MarginBottom(v uint64, t WidthType) TableCellMargins {
	tcm.Bottom = NewCellMargin(v, t)
	return tcm
}

func (tcm *TableCellMargins) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	elem := xml.StartElement{Name: xml.Name{Local: "w:tblCellMar"}}

	// Opening <w:tblCellMar> element
	if err = e.EncodeToken(elem); err != nil {
		return err
	}

	// Encoding <w:Top> element
	if tcm.Top != nil {
		if err = e.EncodeElement(tcm.Top, xml.StartElement{Name: xml.Name{Local: "w:top"}}); err != nil {
			return err
		}
	}

	// Encoding <w:Left> element
	if tcm.Left != nil {
		if err = e.EncodeElement(tcm.Left, xml.StartElement{Name: xml.Name{Local: "w:left"}}); err != nil {
			return err
		}
	}

	// Encoding <w:Bottom> element
	if tcm.Bottom != nil {
		if err = e.EncodeElement(tcm.Bottom, xml.StartElement{Name: xml.Name{Local: "w:bottom"}}); err != nil {
			return err
		}
	}

	// Encoding <w:Right> element
	if tcm.Right != nil {
		if err = e.EncodeElement(tcm.Right, xml.StartElement{Name: xml.Name{Local: "w:right"}}); err != nil {
			return err
		}
	}

	// Closing </w:tblCellMar> element
	return e.EncodeToken(elem.End())
}

func (tcm *TableCellMargins) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name {
			case xml.Name{Space: constants.WMLNamespace, Local: "top"}, xml.Name{Space: constants.AltWMLNamespace, Local: "top"}:
				if err = d.DecodeElement(&tcm.Top, &t); err != nil {
					return err
				}
			case xml.Name{Space: constants.WMLNamespace, Local: "left"}, xml.Name{Space: constants.AltWMLNamespace, Local: "left"}:
				if err = d.DecodeElement(&tcm.Left, &t); err != nil {
					return err
				}
			case xml.Name{Space: constants.WMLNamespace, Local: "bottom"}, xml.Name{Space: constants.AltWMLNamespace, Local: "bottom"}:
				if err = d.DecodeElement(&tcm.Bottom, &t); err != nil {
					return err
				}
			case xml.Name{Space: constants.WMLNamespace, Local: "right"}, xml.Name{Space: constants.AltWMLNamespace, Local: "right"}:
				if err = d.DecodeElement(&tcm.Right, &t); err != nil {
					return err
				}
			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			// Check if it's the end of the w:pPr element
			if t == start.End() {
				return nil
			}
		}
	}
}
