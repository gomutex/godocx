package table

import (
	"encoding/xml"

	"github.com/gomutex/godocx/elemtypes"
	"github.com/gomutex/godocx/wml/ctypes"
)

// This element specifies the set of table-wide properties applied to the current table. These properties affect the appearance of all rows and cells within the parent table, but may be overridden by individual table-level exception, row, and cell level properties as defined by each property.
type TableProperty struct {
	// 1. Referenced Table Style
	Style *TableStyle `xml:"tblStyle,omitempty"`

	// 2. Floating Table Positioning
	FloatPos *FloatPos `xml:"tblpPr,omitempty"`

	// 3.Floating Table Allows Other Tables to Overlap
	Overlap *Overlap `xml:"tblOverlap,omitempty"`

	// 4. Visually Right to Left Table
	BidiVisual *elemtypes.OptBinFlagElem `xml:"bidiVisual,omitempty"`

	// 5. Number of Rows in Row Band
	RowCountInRowBand *ctypes.DecimalNum `xml:"tblStyleRowBandSize,omitempty"`

	// 6. Number of Columns in Column Band
	RowCountInColBand *ctypes.DecimalNum `xml:"tblStyleColBandSize,omitempty"`

	// 7. Preferred Table Width
	Width *ctypes.TableWidth `xml:"tblW,omitempty"`

	// 8.Table Alignment
	Justification *ctypes.Justification `xml:"jc,omitempty"`

	// 9.Table Cell Spacing Default
	CellSpacing *ctypes.TableWidth `xml:"blCellSpacing,omitempty"`

	// 10. Table Indent from Leading Margin
	Indent *ctypes.TableWidth `xml:"tblInd,omitempty"`

	// 11. Table Indent from Leading Margin
	Borders *TableBorders `xml:"tblBorders,omitempty"`

	// 12. Table Shading
	Shading *ctypes.Shading `xml:"shd,omitempty"`

	// 13. Table Layout
	Layout *TableLayout `xml:"tblLayout,omitempty"`

	// 14. Table Cell Margin Defaults
	CellMargin *TableCellMargins `xml:"tblCellMar,omitempty"`

	// 15. Table Style Conditional Formatting Settings
	TableLook *elemtypes.SingleStrVal `xml:"tblLook,omitempty"`
}

func DefaultTableProperty() *TableProperty {
	return &TableProperty{}
}

func (t *TableProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:tblPr"

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	// 1. tblStyle
	if t.Style != nil {
		if err = t.Style.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblStyle"},
		}); err != nil {
			return err
		}
	}

	// 2. tblpPr
	if t.FloatPos != nil {
		if err = t.FloatPos.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblpPr"},
		}); err != nil {
			return err
		}
	}

	// 3. tblOverlap
	if t.Overlap != nil {
		if err = t.Overlap.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblOverlap"},
		}); err != nil {
			return err
		}
	}

	// 4. tblOverlap
	if t.BidiVisual != nil {
		if err = t.BidiVisual.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:bidiVisual"},
		}); err != nil {
			return err
		}
	}

	// 5. tblStyleRowBandSize
	if t.RowCountInRowBand != nil {
		if err = t.RowCountInRowBand.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblStyleRowBandSize"},
		}); err != nil {
			return err
		}
	}

	// 6. tblStyleColBandSize
	if t.RowCountInColBand != nil {
		if err = t.RowCountInColBand.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblStyleColBandSize"},
		}); err != nil {
			return err
		}
	}

	// 7. tblStyleColBandSize
	if t.Width != nil {
		if err = t.Width.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblW"},
		}); err != nil {
			return err
		}
	}

	// 8. jc
	if t.Justification != nil {
		if err = t.Justification.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:jc"},
		}); err != nil {
			return err
		}
	}

	// 9. blCellSpacing
	if t.CellSpacing != nil {
		if err = t.CellSpacing.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:blCellSpacing"},
		}); err != nil {
			return err
		}
	}

	// 10. tblInd
	if t.Indent != nil {
		if err = t.Indent.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblInd"},
		}); err != nil {
			return err
		}
	}

	// 11. tblBorders
	if t.Borders != nil {
		if err = t.Borders.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblBorders"},
		}); err != nil {
			return err
		}
	}

	// 12. shd
	if t.Shading != nil {
		if err = t.Shading.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:shd"},
		}); err != nil {
			return err
		}
	}

	// 13. tblLayout
	if t.Layout != nil {
		if err = t.Layout.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblLayout"},
		}); err != nil {
			return err
		}
	}

	// 14. CellMargin
	if t.CellMargin != nil {
		if err = t.CellMargin.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblCellMar"},
		}); err != nil {
			return err
		}
	}

	// 15. TableLook
	if t.TableLook != nil {
		if err = t.TableLook.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:tblLook"},
		}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
