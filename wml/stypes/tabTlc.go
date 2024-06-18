package stypes

import (
	"encoding/xml"
	"errors"
)

// Custom Tab Stop Leader Character
type CustLeadChar string

const (
	CustLeadCharNone       CustLeadChar = "none"
	CustLeadCharDot        CustLeadChar = "dot"
	CustLeadCharHyphen     CustLeadChar = "hyphen"
	CustLeadCharUnderScore CustLeadChar = "underscore"
	CustLeadCharHeavy      CustLeadChar = "heavy"
	CustLeadCharMiddleDot  CustLeadChar = "middleDot"
	CustLeadCharInvalid    CustLeadChar = ""
)

func CustLeadCharFromStr(val string) (CustLeadChar, error) {
	switch val {
	case "none":
		return CustLeadCharNone, nil
	case "dot":
		return CustLeadCharDot, nil
	case "hyphen":
		return CustLeadCharHyphen, nil
	case "underscore":
		return CustLeadCharUnderScore, nil
	case "heavy":
		return CustLeadCharHeavy, nil
	case "middleDot":
		return CustLeadCharMiddleDot, nil
	default:
		return CustLeadCharInvalid, errors.New("Invalid Lead Char")
	}
}

func (d *CustLeadChar) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := CustLeadCharFromStr(attr.Value)
	if err != nil {
		return err
	}

	*d = val

	return nil
}
