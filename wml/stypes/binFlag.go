package stypes

import (
	"encoding/xml"
	"errors"
)

// This simple type specifies a set of values for any binary (on or off) property defined in a WordprocessingML document.
// A value of on, 1, or true specifies that the property shall be turned on. This is the default value for this attribute, and is implied when the parent element is present, but this attribute is omitted.
//
// A value of off, 0, or false specifies that the property shall be explicitly turned off.
type BinFlag string

const (
	BinFlagZero  BinFlag = "0"
	BinFlagOne   BinFlag = "1"
	BinFlagFalse BinFlag = "false"
	BinFlagTrue  BinFlag = "true"
	BinFlagOff   BinFlag = "off"
	BinFlagOn    BinFlag = "on"
)

func BinFlagFromStr(s string) (BinFlag, error) {
	switch s {
	case "0":
		return BinFlagZero, nil
	case "1":
		return BinFlagOne, nil
	case "false":
		return BinFlagFalse, nil
	case "true":
		return BinFlagTrue, nil
	case "off":
		return BinFlagOff, nil
	case "on":
		return BinFlagOn, nil
	default:
		return "", errors.New("invalid BinFlag string")
	}
}

func (d *BinFlag) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := BinFlagFromStr(attr.Value)
	if err != nil {
		return err
	}

	*d = val

	return nil

}
