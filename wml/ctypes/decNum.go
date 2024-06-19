package ctypes

import (
	"encoding/xml"
	"strconv"
)

type DecimalNum struct {
	Val int `xml:"val,attr"`
}

func NewDecimalNum(value int) *DecimalNum {
	return &DecimalNum{
		Val: value,
	}
}

func (s *DecimalNum) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: strconv.Itoa(s.Val)})
	err := e.EncodeElement("", start)

	return err
}
