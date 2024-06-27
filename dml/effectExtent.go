package dml

import (
	"encoding/xml"
	"strconv"
)

type EffectExtent struct {
	LeftEdge   int64 `xml:"l,attr,omitempty"`
	TopEdge    int64 `xml:"t,attr,omitempty"`
	RightEdge  int64 `xml:"r,attr,omitempty"`
	BottomEdge int64 `xml:"b,attr,omitempty"`
}

func NewEffectExtent(left, top, right, bottom int64) *EffectExtent {
	return &EffectExtent{
		LeftEdge:   left,
		TopEdge:    top,
		RightEdge:  right,
		BottomEdge: bottom,
	}
}

func (x EffectExtent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:effectExtent"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "l"}, Value: strconv.FormatInt(x.LeftEdge, 10)},
		{Name: xml.Name{Local: "t"}, Value: strconv.FormatInt(x.TopEdge, 10)},
		{Name: xml.Name{Local: "r"}, Value: strconv.FormatInt(x.RightEdge, 10)},
		{Name: xml.Name{Local: "b"}, Value: strconv.FormatInt(x.BottomEdge, 10)},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

/*
func (x *EffectExtent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "cx" {
			cx, err := strconv.ParseUint(a.Value, 10, 32)
			if err != nil {
				return nil
			}
			x.Width = cx
		} else if a.Name.Local == "cy" {
			cy, err := strconv.ParseUint(a.Value, 10, 32)
			if err != nil {
				return nil
			}
			x.Height = cy
		}
	}

	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {

			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
} */
