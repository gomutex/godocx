package wml

import (
	"encoding/xml"
	"strconv"
)

// NumberingId represents the ID of a numbering in a document.
type NumberingID struct {
	Val int
}

// NewNumberingId creates a new NumberingId instance.
func NewNumberingID(val int) *NumberingID {
	return &NumberingID{Val: val}
}

// MarshalXML implements the xml.Marshaler interface for NumberingId.
func (n *NumberingID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:numId"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: strconv.Itoa(n.Val)})
	return e.EncodeElement("", start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for NumberingId.
func (n *NumberingID) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "val" {
			val, err := strconv.Atoi(a.Value)
			if err != nil {
				return err
			}
			n.Val = val
			break
		}
	}

	return d.Skip() // Skipping the inner content
}
