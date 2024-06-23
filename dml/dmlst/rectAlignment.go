package dmlst

import (
	"encoding/xml"
	"errors"
)

type RectAlignment string

const (
	RectAlignmentTopLeft     RectAlignment = "tl"  // Rectangle Alignment Enum (Top Left)
	RectAlignmentTop         RectAlignment = "t"   // Rectangle Alignment Enum (Top)
	RectAlignmentTopRight    RectAlignment = "tr"  // Rectangle Alignment Enum (Top Right)
	RectAlignmentLeft        RectAlignment = "l"   // Rectangle Alignment Enum (Left)
	RectAlignmentCenter      RectAlignment = "ctr" // Rectangle Alignment Enum (Center)
	RectAlignmentRight       RectAlignment = "r"   // Rectangle Alignment Enum (Right)
	RectAlignmentBottomLeft  RectAlignment = "bl"  // Rectangle Alignment Enum (Bottom Left)
	RectAlignmentBottom      RectAlignment = "b"   // Rectangle Alignment Enum (Bottom)
	RectAlignmentBottomRight RectAlignment = "br"  // Rectangle Alignment Enum (Bottom Right)
)

// RectAlignmentFromStr converts a string to RectAlignment type.
func RectAlignmentFromStr(value string) (RectAlignment, error) {
	switch value {
	case "tl":
		return RectAlignmentTopLeft, nil
	case "t":
		return RectAlignmentTop, nil
	case "tr":
		return RectAlignmentTopRight, nil
	case "l":
		return RectAlignmentLeft, nil
	case "ctr":
		return RectAlignmentCenter, nil
	case "r":
		return RectAlignmentRight, nil
	case "bl":
		return RectAlignmentBottomLeft, nil
	case "b":
		return RectAlignmentBottom, nil
	case "br":
		return RectAlignmentBottomRight, nil
	default:
		return "", errors.New("Invalid RectAlignment value")
	}
}

// UnmarshalXMLAttr unmarshals XML attribute into RectAlignment.
func (r *RectAlignment) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := RectAlignmentFromStr(attr.Value)
	if err != nil {
		return err
	}
	*r = val
	return nil
}
