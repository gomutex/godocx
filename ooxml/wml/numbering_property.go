package wml

import (
	"encoding/xml"
)

// NumberingProperty represents the properties of a numbering in a document.
type NumberingProperty struct {
	ID    *NumberingID
	Level *IndentLevel
}

// NewNumberingProperty creates a new NumberingProperty instance.
func NewNumberingProperty() *NumberingProperty {
	return &NumberingProperty{}
}

// AddNumber adds the numbering ID and indent level to the NumberingProperty.
func (n *NumberingProperty) AddNumber(id *NumberingID, level *IndentLevel) *NumberingProperty {
	n.ID = id
	n.Level = level
	return n
}

// MarshalXML implements the xml.Marshaler interface for NumberingProperty.
func (n *NumberingProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:numPr"
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if n.ID != nil {
		if err := e.EncodeElement(n.ID, xml.StartElement{Name: xml.Name{Local: "w:numId"}}); err != nil {
			return err
		}
	}

	if n.Level != nil {
		if err := e.EncodeElement(n.Level, xml.StartElement{Name: xml.Name{Local: "w:ilvl"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(start.End())
}

// UnmarshalXML implements the xml.Unmarshaler interface for NumberingProperty.
func (n *NumberingProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		t, err := d.Token()
		if err != nil {
			return err
		}

		switch tok := t.(type) {
		case xml.StartElement:
			switch tok.Name.Local {
			case "numId":
				n.ID = &NumberingID{}
				if err := d.DecodeElement(n.ID, &tok); err != nil {
					return err
				}
			case "ilvl":
				n.Level = &IndentLevel{}
				if err := d.DecodeElement(n.Level, &tok); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if tok == start.End() {
				return nil
			}
		}
	}
}
