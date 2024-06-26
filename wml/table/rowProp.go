package table

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/gomutex/godocx/elemtypes"
	"github.com/gomutex/godocx/wml/ctypes"
)

// Table Row Properties
type RowProperty struct {
	//1. Choice - ZeroOrMore

	//Table Row Conditional Formatting
	Cnf *ctypes.Cnf

	// Associated HTML div ID
	DivId *elemtypes.SingleIntVal

	//Grid Columns Before First Cell
	GridBefore *elemtypes.SingleIntVal

	//Grid Columns After Last Cell
	GridAfter *elemtypes.SingleIntVal

	//Preferred Width Before Table Row
	WidthBefore *ctypes.TableWidth

	//Preferred Width After Table Row
	WidthAfter *ctypes.TableWidth

	//Table Row Cannot Break Across Pages
	CantSplit *elemtypes.OptBinFlagElem

	//Table Row Height
	Height *ctypes.TableRowHeight

	//Repeat Table Row on Every New Page
	Header *elemtypes.OptBinFlagElem

	//Table Row Cell Spacing
	CellSpacing *ctypes.TableWidth

	// Table Row Alignment
	JC *ctypes.Justification

	//Hidden Table Row Marker
	Hidden *elemtypes.OptBinFlagElem

	//2.Inserted Table Row
	Ins *ctypes.TrackChange

	//3. Deleted Table Row
	Del *ctypes.TrackChange

	//4.Revision Information for Table Row Properties
	Change *TRPrChange
}

// NewRowProperty creates a new RowProperty instance.
func DefaultRowProperty() *RowProperty {
	return &RowProperty{}
}

func (r *RowProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:trPr"

	err := e.EncodeToken(start)
	if err != nil {
		return fmt.Errorf("marshaling row propery: %w", err)
	}

	if err = r.MarshalChoice(e, start); err != nil {
		return err
	}

	if r.Ins != nil {
		if err := r.Ins.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:ins"}}); err != nil {
			return err
		}
	}

	if r.Del != nil {
		if err := r.Del.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:del"}}); err != nil {
			return err
		}
	}

	if r.Change != nil {
		if err := r.Change.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:trPrChange"}}); err != nil {
			return err
		}
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (r *RowProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

loop:
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "cnfStyle":
				r.Cnf = &ctypes.Cnf{}
				if err = d.DecodeElement(r.Cnf, &elem); err != nil {
					return err
				}

			case "divId":
				r.DivId = &elemtypes.SingleIntVal{}
				if err = d.DecodeElement(r.DivId, &elem); err != nil {
					return err
				}

			case "gridBefore":
				r.GridBefore = &elemtypes.SingleIntVal{}
				if err := d.DecodeElement(r.GridBefore, &elem); err != nil {
					return err
				}

			case "gridAfter":
				r.GridAfter = &elemtypes.SingleIntVal{}
				if err := d.DecodeElement(r.GridAfter, &elem); err != nil {
					return err
				}

			case "tblWBefore":
				r.WidthBefore = &ctypes.TableWidth{}
				if err := d.DecodeElement(r.WidthBefore, &elem); err != nil {
					return err
				}

			case "tblWAfter":
				r.WidthAfter = &ctypes.TableWidth{}
				if err := d.DecodeElement(r.WidthAfter, &elem); err != nil {
					return err
				}

			case "cantSplit":
				r.CantSplit = &elemtypes.OptBinFlagElem{}
				if err := d.DecodeElement(r.CantSplit, &elem); err != nil {
					return err
				}

			case "trHeight":
				r.Height = &ctypes.TableRowHeight{}
				if err := d.DecodeElement(r.Height, &elem); err != nil {
					return err
				}

			case "tblHeader":
				r.Header = &elemtypes.OptBinFlagElem{}
				if err := d.DecodeElement(r.Header, &elem); err != nil {
					return err
				}

			case "tblCellSpacing":
				r.CellSpacing = &ctypes.TableWidth{}
				if err := d.DecodeElement(r.CellSpacing, &elem); err != nil {
					return err
				}

			case "jc":
				r.JC = &ctypes.Justification{}
				if err := d.DecodeElement(r.JC, &elem); err != nil {
					return err
				}

			case "hidden":
				r.Hidden = &elemtypes.OptBinFlagElem{}
				if err := d.DecodeElement(r.Hidden, &elem); err != nil {
					return err
				}

			case "ins":
				r.Ins = &ctypes.TrackChange{}
				if err := d.DecodeElement(r.Ins, &elem); err != nil {
					return err
				}
			case "del":
				r.Del = &ctypes.TrackChange{}
				if err := d.DecodeElement(r.Del, &elem); err != nil {
					return err
				}
			case "trPrChange":
				r.Change = &TRPrChange{}
				if err := d.DecodeElement(r.Change, &elem); err != nil {
					return err
				}
			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if elem.Name.Local == start.Name.Local {
				break loop
			}
		}
	}

	return nil
}

func (r *RowProperty) MarshalChoice(e *xml.Encoder, start xml.StartElement) error {
	var err error

	if r.Cnf != nil {
		if err = r.Cnf.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:cnfStyle"},
		}); err != nil {
			return fmt.Errorf("marshaling cnfstyle failed in table row property: %w", err)
		}
	}

	if r.DivId != nil {
		if err = r.DivId.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:divId"},
		}); err != nil {
			return fmt.Errorf("marshaling divid failed in table row property: %w", err)
		}
	}

	if r.GridBefore != nil {
		if err = r.GridBefore.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:gridBefore"}}); err != nil {
			return fmt.Errorf("marshaling gridBefore failed in table row property: %w", err)
		}
	}

	if r.GridAfter != nil {
		if err = r.GridAfter.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:gridAfter"}}); err != nil {
			return fmt.Errorf("marshaling gridAfter failed in table row property: %w", err)
		}
	}

	if r.WidthBefore != nil {
		if err = r.WidthBefore.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:tblWBefore"}}); err != nil {
			return fmt.Errorf("marshaling tblWBefore failed in table row property: %w", err)
		}
	}

	if r.WidthAfter != nil {
		if err = r.WidthAfter.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:tblWAfter"}}); err != nil {
			return fmt.Errorf("marshaling tblWAfter failed in table row property: %w", err)
		}
	}

	if r.CantSplit != nil {
		if err = r.CantSplit.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:cantSplit"}}); err != nil {
			return fmt.Errorf("marshaling cantSplit failed in table row property: %w", err)
		}
	}

	if r.Height != nil {
		if err = r.Height.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:trHeight"}}); err != nil {
			return fmt.Errorf("marshaling trHeight failed in table row property: %w", err)
		}
	}

	if r.Header != nil {
		if err = r.Header.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:tblHeader"}}); err != nil {
			return fmt.Errorf("marshaling tblHeader failed in table row property: %w", err)
		}

	}

	if r.CellSpacing != nil {
		if err = r.CellSpacing.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:tblCellSpacing"}}); err != nil {
			return fmt.Errorf("marshaling tblCellSpacing failed in table row property: %w", err)
		}
	}

	if r.JC != nil {
		if err = r.JC.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:jc"}}); err != nil {
			return fmt.Errorf("marshaling jc failed in table row property: %w", err)
		}
	}

	if r.Hidden != nil {
		if err = r.Hidden.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:hidden"}}); err != nil {
			return fmt.Errorf("marshaling hidden failed in table row property: %w", err)
		}
	}

	return nil
}

type TRPrChange struct {
	ID     int         `xml:"id,attr"`
	Author string      `xml:"author,attr"`
	Date   *string     `xml:"date,attr,omitempty"`
	Prop   RowProperty `xml:"tcPr"`
}

func (t TRPrChange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:TRPrChange"

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
