package stypes

import (
	"encoding/xml"
	"errors"
)

type TblOverlap string

const (
	TblOverlapNever   TblOverlap = "never"
	TblOverlapOverlap TblOverlap = "overlap"
)

func TblOverlapFromStr(value string) (TblOverlap, error) {
	switch value {
	case "never":
		return TblOverlapNever, nil
	case "overlap":
		return TblOverlapOverlap, nil
	default:
		return "", errors.New("Invalid TblOverlap value")
	}
}

func (to *TblOverlap) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := TblOverlapFromStr(attr.Value)
	if err != nil {
		return err
	}
	*to = val
	return nil
}
