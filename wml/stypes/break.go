package stypes

import (
	"encoding/xml"
	"errors"
)

type BreakType string

const (
	BreakTypePage         BreakType = "page"         // Page Break
	BreakTypeColumn       BreakType = "column"       // Column Break
	BreakTypeTextWrapping BreakType = "textWrapping" // Line Break
	BreakTypeInvalid      BreakType = ""
)

type BreakClear string

const (
	BreakClearNone    BreakClear = "none"  // Restart On Next Line
	BreakClearLeft    BreakClear = "left"  // Restart In Next Text Region When In Leftmost Position
	BreakClearRight   BreakClear = "right" // Restart In Next Text Region When In Rightmost Position
	BreakClearAll     BreakClear = "all"   // Restart On Next Full Line
	BreakClearInvalid BreakClear = ""
)

func BreakTypeFromStr(value string) (BreakType, error) {
	switch value {
	case "page":
		return BreakTypePage, nil
	case "column":
		return BreakTypeColumn, nil
	case "textWrapping":
		return BreakTypeTextWrapping, nil
	default:
		return BreakTypeInvalid, errors.New("invalid BreakType value")
	}
}

func (d *BreakType) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := BreakTypeFromStr(attr.Value)
	if err != nil {
		return err
	}

	*d = val

	return nil
}
func BreakClearFromStr(value string) (BreakClear, error) {
	switch value {
	case "none":
		return BreakClearNone, nil
	case "left":
		return BreakClearLeft, nil
	case "right":
		return BreakClearRight, nil
	case "all":
		return BreakClearAll, nil
	default:
		return BreakClearInvalid, errors.New("invalid BreakClear value")
	}
}

func (d *BreakClear) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := BreakClearFromStr(attr.Value)
	if err != nil {
		return err
	}

	*d = val

	return nil
}
