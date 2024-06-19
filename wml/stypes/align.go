package stypes

import (
	"encoding/xml"
	"errors"
)

// Align Type
type Align string

// Constants for valid values
const (
	AlignLeft    Align = "left"    // Left Aligned Horizontally
	AlignCenter  Align = "center"  // Centered Horizontally
	AlignRight   Align = "right"   // Right Aligned Horizontally
	AlignInside  Align = "inside"  // Inside
	AlignOutside Align = "outside" // Outside
)

// AlignFromStr converts a string to Align type.
func AlignFromStr(value string) (Align, error) {
	switch value {
	case "left":
		return AlignLeft, nil
	case "center":
		return AlignCenter, nil
	case "right":
		return AlignRight, nil
	case "inside":
		return AlignInside, nil
	case "outside":
		return AlignOutside, nil
	default:
		return "", errors.New("Invalid Align value")
	}
}

// UnmarshalXMLAttr unmarshals XML attribute into Align.
func (x *Align) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := AlignFromStr(attr.Value)
	if err != nil {
		return err
	}
	*x = val
	return nil
}
