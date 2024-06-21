package stypes

import (
	"encoding/xml"
	"errors"
)

// XAlign Type
type XAlign string

// Constants for valid values
const (
	XAlignLeft    XAlign = "left"    // Left XAligned Horizontally
	XAlignCenter  XAlign = "center"  // Centered Horizontally
	XAlignRight   XAlign = "right"   // Right XAligned Horizontally
	XAlignInside  XAlign = "inside"  // Inside
	XAlignOutside XAlign = "outside" // Outside
)

// XAlignFromStr converts a string to XAlign type.
func XAlignFromStr(value string) (XAlign, error) {
	switch value {
	case "left":
		return XAlignLeft, nil
	case "center":
		return XAlignCenter, nil
	case "right":
		return XAlignRight, nil
	case "inside":
		return XAlignInside, nil
	case "outside":
		return XAlignOutside, nil
	default:
		return "", errors.New("Invalid XAlign value")
	}
}

// UnmarshalXMLAttr unmarshals XML attribute into XAlign.
func (x *XAlign) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := XAlignFromStr(attr.Value)
	if err != nil {
		return err
	}
	*x = val
	return nil
}

// YAlign Type
type YAlign string

// Constants for valid values
const (
	YAlignInline  YAlign = "inline"  // In line With Text
	YAlignTop     YAlign = "top"     // Top
	YAlignCenter  YAlign = "center"  // Centered Vertically
	YAlignBottom  YAlign = "bottom"  // Bottom
	YAlignInside  YAlign = "inside"  // Inside Anchor Extents
	YAlignOutside YAlign = "outside" // Outside Anchor Extents
)

// YAlignFromStr converts a string to YAlign type.
func YAlignFromStr(value string) (YAlign, error) {
	switch value {
	case "inline":
		return YAlignInline, nil
	case "top":
		return YAlignTop, nil
	case "center":
		return YAlignCenter, nil
	case "bottom":
		return YAlignBottom, nil
	case "inside":
		return YAlignInside, nil
	case "outside":
		return YAlignOutside, nil
	default:
		return "", errors.New("Invalid YAlign value")
	}
}

// UnmarshalXMLAttr unmarshals XML attribute into YAlign.
func (x *YAlign) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := YAlignFromStr(attr.Value)
	if err != nil {
		return err
	}
	*x = val
	return nil
}
