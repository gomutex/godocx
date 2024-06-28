package stypes

import (
	"encoding/xml"
	"errors"
)

// This simple type specifies a set of values for any binary (on or off) property defined in a WordprocessingML document.
// A value of on, 1, or true specifies that the property shall be turned on. This is the default value for this attribute, and is implied when the parent element is present, but this attribute is omitted.
//
// A value of off, 0, or false specifies that the property shall be explicitly turned off.
type OnOff string

const (
	OnOffZero  OnOff = "0"
	OnOffOne   OnOff = "1"
	OnOffFalse OnOff = "false"
	OnOffTrue  OnOff = "true"
	OnOffOff   OnOff = "off"
	OnOffOn    OnOff = "on"
)

func OnOffFromStr(s string) (OnOff, error) {
	switch s {
	case "0":
		return OnOffZero, nil
	case "1":
		return OnOffOne, nil
	case "false":
		return OnOffFalse, nil
	case "true":
		return OnOffTrue, nil
	case "off":
		return OnOffOff, nil
	case "on":
		return OnOffOn, nil
	default:
		return "", errors.New("invalid OnOff string")
	}
}

func (d *OnOff) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := OnOffFromStr(attr.Value)
	if err != nil {
		return err
	}

	*d = val

	return nil

}
