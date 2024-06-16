package sections

import (
	"encoding/xml"
	"errors"
	"strconv"
)

// Page Size : w:pgSz
type PageSize struct {
	Width  *uint64      // w:w
	Height *uint64      // w:h
	Orient *Orientation //Page Orientation
	Code   *int         //Printer Paper Code : w:code
}

type Orientation string

const (
	OrientPortrait  Orientation = "portrait"
	OrientLandscape Orientation = "landscape"
)

func OrientFromStr(value string) (Orientation, error) {
	switch value {
	case "portrait":
		return OrientPortrait, nil
	case "landscape":
		return OrientLandscape, nil
	default:
		return "", errors.New("Invalid Orient Input")
	}
}

func (p *PageSize) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:pgSz"

	if p.Width != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:w"}, Value: strconv.FormatUint(*p.Width, 10)})
	}

	if p.Height != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:h"}, Value: strconv.FormatUint(*p.Height, 10)})
	}

	if p.Orient != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:orient"}, Value: string(*p.Orient)})
	}

	if p.Code != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:code"}, Value: strconv.Itoa(*p.Code)})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (p *PageSize) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		switch a.Name.Local {
		case "w":
			width, err := strconv.ParseUint(a.Value, 10, 64)
			if err != nil {
				return nil
			}
			p.Width = &width
		case "h":
			height, err := strconv.ParseUint(a.Value, 10, 64)
			if err != nil {
				return nil
			}
			p.Height = &height
		case "orient":
			orient, err := OrientFromStr(a.Value)
			if err != nil {
				return nil
			}
			p.Orient = &orient
		case "code":
			c, err := strconv.Atoi(a.Value)
			if err != nil {
				return nil
			}
			code := int(c)
			p.Code = &code
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
}
