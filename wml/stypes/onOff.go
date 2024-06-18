package stypes

import (
	"encoding/xml"
	"errors"
)

// On/Off Value
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
