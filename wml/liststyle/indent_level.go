package liststyle

import (
	"encoding/xml"
	"strconv"
)

// IndentLevel represents the indentation level of a numbering in a document.
type IndentLevel struct {
	Val int
}

// NewIndentLevel creates a new IndentLevel instance.
func NewIndentLevel(val int) *IndentLevel {
	return &IndentLevel{Val: val}
}

// MarshalXML implements the xml.Marshaler interface for IndentLevel.
func (i *IndentLevel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:ilvl"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: strconv.Itoa(i.Val)})
	return e.EncodeElement("", start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for IndentLevel.
func (i *IndentLevel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "val" {
			val, err := strconv.Atoi(a.Value)
			if err != nil {
				return err
			}
			i.Val = val
			break
		}
	}

	return d.Skip() // Skipping the inner content
}
