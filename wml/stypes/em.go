package stypes

import (
	"encoding/xml"
	"errors"
)

// Em Type
type Em string

const (
	EmNone     Em = "none"     // No Emphasis Mark
	EmDot      Em = "dot"      // Dot Emphasis Mark Above Characters
	EmComma    Em = "comma"    // Comma Emphasis Mark Above Characters
	EmCircle   Em = "circle"   // Circle Emphasis Mark Above Characters
	EmUnderDot Em = "underDot" // Dot Emphasis Mark Below Characters
)

func EmFromStr(value string) (Em, error) {
	switch value {
	case "none":
		return EmNone, nil
	case "dot":
		return EmDot, nil
	case "comma":
		return EmComma, nil
	case "circle":
		return EmCircle, nil
	case "underDot":
		return EmUnderDot, nil
	default:
		return "", errors.New("Invalid Em value")
	}
}

func (e *Em) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := EmFromStr(attr.Value)
	if err != nil {
		return err
	}

	*e = val

	return nil
}
