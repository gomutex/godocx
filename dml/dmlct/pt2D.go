package dmlct

import (
	"encoding/xml"
	"strconv"
)

// Wrapping Polygon Point2D
type Point2D struct {
	XAxis uint64 `xml:"x,attr,omitempty"`
	YAxis uint64 `xml:"y,attr,omitempty"`
}

func NewPoint2D(x, y uint64) Point2D {
	return Point2D{
		XAxis: uint64(x),
		YAxis: uint64(y),
	}
}

func (p Point2D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "x"}, Value: strconv.FormatUint(p.XAxis, 10)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "y"}, Value: strconv.FormatUint(p.YAxis, 10)})

	return e.EncodeElement("", start)
}
