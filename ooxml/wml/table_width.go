package wml

import (
	"encoding/xml"
	"strconv"
)

type WidthType string

const (
	WidthTypeDxa         WidthType = "dxa"
	WidthTypeAuto        WidthType = "Auto"
	WidthTypePct         WidthType = "pct"
	WidthTypeNil         WidthType = "nil"
	WidthTypeUnsupported WidthType = "unsupported"
)

// TableWidth represents the width of a table in a document.
type TableWidth struct {
	Width     uint64
	WidthType WidthType
}

func (t *TableWidth) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:tblW"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:w"}, Value: strconv.FormatUint(uint64(t.Width), 10)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:type"}, Value: string(t.WidthType)})

	return e.EncodeElement("", start)
}

func (t *TableWidth) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "w":
			valueStr := attr.Value
			if valueStr != "" {
				value, err := strconv.ParseUint(valueStr, 10, 0)
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
