package stypes

import (
	"encoding/xml"
	"errors"
)

// Text Frame Drop Cap Location
type DropCap string

const (
	DropCapNone   DropCap = "none"   // Not Drop Cap
	DropCapInside DropCap = "drop"   // Drop Cap Inside Margin
	DropCapMargin DropCap = "margin" // Drop Cap Outside Margin
)

// DropCapFromStr converts a string to DropCap type.
func DropCapFromStr(value string) (DropCap, error) {
	switch value {
	case "none":
		return DropCapNone, nil
	case "drop":
		return DropCapInside, nil
	case "margin":
		return DropCapMargin, nil
	default:
		return "", errors.New("Invalid DropCap value")
	}
}

// UnmarshalXMLAttr unmarshals XML attribute into DropCap.
func (d *DropCap) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := DropCapFromStr(attr.Value)
	if err != nil {
		return err
	}
	*d = val
	return nil
}
