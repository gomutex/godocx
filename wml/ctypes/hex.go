package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// LongHexNum represents a long hexadecimal number value.
type LongHexNum struct {
	Val stypes.LongHexNum `xml:"val,attr"`
}

// NewLongHexNum validates and creates a new LongHexNum instance.
func NewLongHexNum(value string) (*LongHexNum, error) {
	longHexNum, err := stypes.LongHexNumFromStr(value)
	if err != nil {
		return nil, err
	}
	return &LongHexNum{
		Val: longHexNum,
	}, nil
}

// MarshalXML implements the xml.Marshaler interface.
func (s LongHexNum) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Validate the LongHexNum value before marshaling
	if _, err := stypes.LongHexNumFromStr(string(s.Val)); err != nil {
		return err
	}

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(s.Val)})
	return e.EncodeElement("", start)
}
