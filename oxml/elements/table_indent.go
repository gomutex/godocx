package elements

import (
	"encoding/xml"
	"strconv"
)

// TableWidth represents the width of a table in a document.
type TableIndent struct {
	Width     int
	WidthType WidthType
}

func NewTableIndent(width int, widthType WidthType) *TableIndent {
	return &TableIndent{
		Width:     width,
		WidthType: widthType,
	}
}

func (t *TableIndent) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:tblInd"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:w"}, Value: strconv.FormatInt(int64(t.Width), 10)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:type"}, Value: string(t.WidthType)})

	return e.EncodeElement("", start)
}

func (t *TableIndent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "w":
			valueStr := attr.Value
			if valueStr != "" {
				value, err := strconv.Atoi(valueStr)
				if err != nil {
					return err
				}
				t.Width = value
			}
		case "type":
			t.WidthType = WidthType(attr.Value)
		}
	}

	return d.Skip() // Skipping the inner content
}
