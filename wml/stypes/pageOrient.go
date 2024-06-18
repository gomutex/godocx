package stypes

import (
	"encoding/xml"
	"errors"
)

type PageOrient string

const (
	PageOrientPortrait  PageOrient = "portrait"
	PageOrientLandscape PageOrient = "landscape"
)

func PageOrientFromStr(value string) (PageOrient, error) {
	switch value {
	case "portrait":
		return PageOrientPortrait, nil
	case "landscape":
		return PageOrientLandscape, nil
	default:
		return "", errors.New("Invalid Orient Input")
	}
}

func (d *PageOrient) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := PageOrientFromStr(attr.Value)
	if err != nil {
		return err
	}

	*d = val

	return nil

}
