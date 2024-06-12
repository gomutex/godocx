package wml

import (
	"encoding/xml"
	"strconv"
)

type TableBorderElement struct {
	BorderType BorderType
	Position   TableBorderPosition
	Size       uint64
	Space      uint64
	Color      string
}

type TableBorders struct {
	Top     *TableBorderElement
	Bottom  *TableBorderElement
	Left    *TableBorderElement
	Right   *TableBorderElement
	InsideH *TableBorderElement
	InsideV *TableBorderElement
}

// DefaulTableBorderElement creates a new TableBorderElement instance with default values.
// It sets the default border type to single, size to 2, space to 0, color to "000000",
// and position to TableBorderPositionBottom.
func DefaulTableBorderElement() *TableBorderElement {
	return &TableBorderElement{
		BorderType: BorderTypeSingle,
		Size:       2,
		Space:      0,
		Color:      "000000",
		Position:   TableBorderPositionBottom,
	}
}

// NewTableBorderElement returns a new table border with the given properties.
func NewTableBorderElement(borderType BorderType, size uint64, color string, position TableBorderPosition, space uint64) *TableBorderElement {
	return &TableBorderElement{
		BorderType: borderType,
		Size:       size,
		Color:      color,
		Position:   position,
		Space:      space,
	}
}

func DefaultTableBorders() *TableBorders {
	return &TableBorders{}
}

// MarshalXML implements the xml.Marshaler interface for TableBorderElement.
func (t *TableBorderElement) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = string(t.Position)

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(t.BorderType)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:sz"}, Value: strconv.FormatInt(int64(t.Size), 10)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:space"}, Value: strconv.FormatInt(int64(t.Space), 10)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:color"}, Value: t.Color})

	return e.EncodeElement("", start)
}

func (t *TableBorderElement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {

	switch start.Name.Local {
	case "top":
		t.Position = TableBorderPositionTop
	case "bottom":
		t.Position = TableBorderPositionBottom
	case "right":
		t.Position = TableBorderPositionRight
	case "left":
		t.Position = TableBorderPositionLeft
	case "insideH":
		t.Position = TableBorderPositionInsideH
	case "insideV":
		t.Position = TableBorderPositionInsideV
	}

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "sz":
			valueStr := attr.Value
			if valueStr != "" {
				value, err := strconv.ParseUint(valueStr, 10, 0)
				if err != nil {
					return err
				}
				t.Size = value
			}
		case "space":
			valueStr := attr.Value
			if valueStr != "" {
				value, err := strconv.ParseUint(valueStr, 10, 0)
				if err != nil {
					return err
				}
				t.Space = value
			}
		case "val":
			t.BorderType = BorderType(attr.Value)
		case "color":
			t.Color = attr.Value
		}
	}

	return d.Skip() // Skipping the inner content
}

func (t *TableBorders) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "tblBorders"

	if err := e.EncodeToken(start); err != nil {
		return err
	}
	// Marshal each border individually
	if t.Top != nil {
		if err := e.Encode(t.Top); err != nil {
			return err
		}
	}
	if t.Left != nil {
		if err := e.Encode(t.Left); err != nil {
			return err
		}
	}
	if t.Bottom != nil {
		if err := e.Encode(t.Bottom); err != nil {
			return err
		}
	}
	if t.Right != nil {
		if err := e.Encode(t.Right); err != nil {
			return err
		}
	}
	if t.InsideH != nil {
		if err := e.Encode(t.InsideH); err != nil {
			return err
		}
	}
	if t.InsideV != nil {
		if err := e.Encode(t.InsideV); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})

}

func (t *TableBorders) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	// Loop through each element within tblBorders
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "top":

				t.Top = &TableBorderElement{}
				if err := d.DecodeElement(t.Top, &elem); err != nil {
					return err
				}
			case "left":
				t.Left = &TableBorderElement{}
				if err := d.DecodeElement(t.Left, &elem); err != nil {
					return err
				}
			case "bottom":
				t.Bottom = &TableBorderElement{}
				if err := d.DecodeElement(t.Bottom, &elem); err != nil {
					return err
				}
			case "right":
				t.Right = &TableBorderElement{}
				if err := d.DecodeElement(t.Right, &elem); err != nil {
					return err
				}
			case "insideH":
				t.InsideH = &TableBorderElement{}
				if err := d.DecodeElement(t.InsideH, &elem); err != nil {
					return err
				}
			case "insideV":
				t.InsideV = &TableBorderElement{}
				if err := d.DecodeElement(t.InsideV, &elem); err != nil {
					return err
				}
			}

		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
}
