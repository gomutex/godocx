package table

import (
	"encoding/xml"
	"errors"
	"strconv"

	"github.com/gomutex/godocx/elemtypes"
	"github.com/gomutex/godocx/wml/ctypes"
)

type CellProperty struct {

	// 1. Table Cell Conditional Formatting
	CnfStyle *ctypes.Cnf `xml:"cnfStyle,omitempty"`

	// 2. Preferred Table Cell Width
	Width *ctypes.TableWidth `xml:"tcW,omitempty"`

	// 3.Grid Columns Spanned by Current Table Cell
	GridSpan *elemtypes.SingleIntVal `xml:"gridSpan,omitempty"`

	// 4.Horizontally Merged Cell
	HMerge *ctypes.Merge `xml:"hMerge,omitempty"`

	// 5.Vertically Merged Cell
	VMerge *ctypes.Merge `xml:"vMerge,omitempty"`

	// 6.Table Cell Borders
	Borders *CellBorders `xml:"tcBorders,omitempty"`

	//7.Table Cell Shading
	Shading *ctypes.Shading `xml:"shd,omitempty"`

	//8.Don't Wrap Cell Content
	NoWrap *elemtypes.OptBinFlagElem `xml:"noWrap,omitempty"`

	//9.Single Table Cell Margins
	Margins *CellMargins `xml:"tcMar,omitempty"`

	//10.Table Cell Text Flow Direction
	TextDirection *ctypes.TextDirection `xml:"textDirection,omitempty"`

	//11.Fit Text Within Cell
	FitText *elemtypes.OptBinFlagElem `xml:"tcFitText,omitempty"`

	//12. Table Cell Vertical Alignment
	VertAlign *VertAlign `xml:"vAlign,omitempty"`

	//13.Ignore End Of Cell Marker In Row Height Calculation
	HideMark *elemtypes.OptBinFlagElem `xml:"hideMark,omitempty"`

	//14. Choice - ZeroOrOne
	// At max only one of these element should exist
	CellInsertion *ctypes.TrackChange `xml:"cellIns,omitempty"`
	CellDeletion  *ctypes.TrackChange `xml:"cellDel,omitempty"`
	//Vertically Merged/Split Table Cells
	CellMerge *CellMerge `xml:"cellMerge,omitempty"`

	//15.Revision Information for Table Cell Properties
	PrChange *TCPrChange `xml:"tcPrChange,omitempty"`
}

func (t *CellProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:tcPr"

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	//1. Table Cell Conditional Formatting
	if t.CnfStyle != nil {
		if err = t.CnfStyle.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:cnfStyle"}}); err != nil {
			return err
		}
	}

	//2. Preferred Table Cell Width
	if t.Width != nil {
		if err = t.Width.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:tcW"}}); err != nil {
			return err
		}
	}

	//3. Grid Columns Spanned by Current Table Cell
	if t.GridSpan != nil {
		if err = t.GridSpan.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:gridSpan"}}); err != nil {
			return err
		}
	}

	//4. Horizontally Merged Cell
	if t.HMerge != nil {
		if err = t.HMerge.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:hMerge"}}); err != nil {
			return err
		}
	}

	//5. Vertically Merged Cell
	if t.VMerge != nil {
		if err = t.VMerge.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:vMerge"}}); err != nil {
			return err
		}
	}

	//6.Table Cell Borders
	if t.Borders != nil {
		if err = t.Borders.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:tcBorders"}}); err != nil {
			return err
		}
	}

	//7.Table Cell Shading
	if t.Shading != nil {
		if err = t.Shading.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:shd"}}); err != nil {
			return err
		}
	}

	//8.Don't Wrap Cell Content
	if t.NoWrap != nil {
		if err = t.NoWrap.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:noWrap"}}); err != nil {
			return err
		}
	}

	//9.Single Table Cell Margins
	if t.Margins != nil {
		if err = t.Margins.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:tcMar"}}); err != nil {
			return err
		}
	}

	//10.Table Cell Text Flow Direction
	if t.TextDirection != nil {
		if err = t.TextDirection.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:textDirection"}}); err != nil {
			return err
		}
	}

	//11.Table Cell Text Flow Direction
	if t.FitText != nil {
		if err = t.FitText.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:tcFitText"}}); err != nil {
			return err
		}
	}

	//12.Table Cell Vertical Alignment
	if t.VertAlign != nil {
		if err = t.VertAlign.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:vAlign"}}); err != nil {
			return err
		}
	}

	//13.Ignore End Of Cell Marker In Row Height Calculation
	if t.HideMark != nil {
		if err = t.HideMark.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:hideMark"}}); err != nil {
			return err
		}
	}

	var nMarkupElems uint8
	//14. Choice: Cell Markup Elements
	if t.CellInsertion != nil {
		if err = t.CellInsertion.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:cellIns"},
		}); err != nil {
			return err
		}
		nMarkupElems += 1
	}

	if t.CellDeletion != nil {
		if err = t.CellDeletion.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:cellDel"},
		}); err != nil {
			return err
		}
		nMarkupElems += 1
	}

	if t.CellMerge != nil {
		if err = t.CellMerge.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:cellMerge"},
		}); err != nil {
			return err
		}
		nMarkupElems += 1
	}

	if nMarkupElems > 1 {
		return errors.New("more than 1 element found in EG_CellMarkupElements when marshaling table's cell property")
	}

	//15. Revision Information for Table Cell Properties
	if t.PrChange != nil {
		if err = t.PrChange.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

type TCPrChange struct {
	ID     int          `xml:"id,attr"`
	Author string       `xml:"author,attr"`
	Date   *string      `xml:"date,attr,omitempty"`
	Prop   CellProperty `xml:"tcPr"`
}

func (t TCPrChange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:tcPrChange"

	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "w:id"}, Value: strconv.Itoa(t.ID)},
		{Name: xml.Name{Local: "w:author"}, Value: t.Author},
	}

	if t.Date != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:date"}, Value: *t.Date})
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if err := t.Prop.MarshalXML(e, xml.StartElement{}); err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
