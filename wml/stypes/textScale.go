package stypes

import (
	"encoding/xml"
	"errors"
	"strconv"
)

type TextScale uint16

func TextScaleFromUint16(u uint16) (TextScale, error) {
	if u > 600 {
		return 0, errors.New("Invalid Text Scale")
	}

	return TextScale(u), nil
}

func TextScaleFromStr(s string) (TextScale, error) {
	u, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return 0, err
	}
	return TextScaleFromUint16(uint16(u))
}

func (t *TextScale) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := TextScaleFromStr(attr.Value)
	if err != nil {
		return err
	}

	*t = val

	return nil

}

func (t *TextScale) ToStr() string {
	return strconv.FormatUint(uint64(*t), 10)
}
