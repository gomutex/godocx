package ctypes

import (
	"encoding/xml"
	"fmt"

	"github.com/gomutex/godocx/wml/stypes"
)

// Style Conditional Table Formatting Properties
type TableStyleProp struct {
	// Sequence:

	//1.Table Style Conditional Formatting Paragraph Properties
	ParaProp *ParagraphProp `xml:"pPr,omitempty"`

	//2.Table Style Conditional Formatting Run Properties
	RunProp *RunProperty `xml:"rPr,omitempty"`

	//3.Table Style Conditional Formatting Table Properties
	TableProp *TableProp `xml:"tblPr,omitempty"`

	//4.Table Style Conditional Formatting Table Row Properties
	RowProp *RowProperty `xml:"trPr,omitempty"`

	//5.Table Style Conditional Formatting Table Cell Properties
	CellProp *CellProperty `xml:"tcPr,omitempty"`

	//Attributes:

	//Table Style Conditional Formatting Type
	Type stypes.TblStyleOverrideType `xml:"type,attr"`
}

func (t *TableStyleProp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:tblStylePr"

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:type"}, Value: string(t.Type)})

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	//1.Table Style Conditional Formatting Paragraph Properties
	if t.ParaProp != nil {
		if err := t.ParaProp.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:pPr"}}); err != nil {
			return fmt.Errorf("TableStyleProp ParaProp: %w", err)
		}
	}

	//2.Table Style Conditional Formatting Run Properties
	if t.RunProp != nil {
		if err := t.RunProp.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:rPr"}}); err != nil {
			return fmt.Errorf("TableStyleProp RunProp: %w", err)
		}
	}

	//3.Table Style Conditional Formatting Table Properties
	if t.TableProp != nil {
		if err := t.TableProp.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:tblPr"}}); err != nil {
			return fmt.Errorf("TableStyleProp TableProp: %w", err)
		}
	}

	//4.Table Style Conditional Formatting Table Row Properties
	if t.RowProp != nil {
		if err := t.RowProp.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:trPr"}}); err != nil {
			return fmt.Errorf("TableStyleProp RowProp: %w", err)
		}
	}

	//5.Table Style Conditional Formatting Table Cell Properties
	if t.CellProp != nil {
		if err := t.CellProp.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:tcPr"}}); err != nil {
			return fmt.Errorf("TableStyleProp CellProp: %w", err)
		}
	}

	return e.EncodeToken(start.End())
}
