package ctypes

import (
	"encoding/xml"
	"fmt"
)

// Numbering Definition Instance Reference
type NumProp struct {
	//Numbering Level Reference
	ILvl *DecimalNum `xml:"ilvl,omitempty"`

	//Numbering Definition Instance Reference
	NumID *DecimalNum `xml:"numId,omitempty"`

	//Previous Paragraph Numbering Properties
	NumChange *TrackChangeNum `xml:"numberingChange,omitempty"`

	//Inserted Numbering Properties
	Ins *TrackChange `xml:"ins,omitempty"`
}

// NewNumberingProperty creates a new NumberingProperty instance.
func NewNumberingProperty() *NumProp {
	return &NumProp{}
}

func (n NumProp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:numPr"

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if n.ILvl != nil {
		if err = n.ILvl.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:ilvl"},
		}); err != nil {
			return fmt.Errorf("ILvl: %w", err)
		}
	}

	if n.NumID != nil {
		if err = n.NumID.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:numId"},
		}); err != nil {
			return fmt.Errorf("NumID: %w", err)
		}
	}

	if n.NumChange != nil {
		if err = n.NumChange.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:numberingChange"},
		}); err != nil {
			return fmt.Errorf("NumChange: %w", err)
		}
	}

	if n.Ins != nil {
		if err = n.Ins.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:ins"},
		}); err != nil {
			return fmt.Errorf("NumID: %w", err)
		}
	}

	return e.EncodeToken(start.End())
}
