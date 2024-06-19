package stypes

import (
	"encoding/xml"
	"errors"
)

// Wrap Type
type Wrap string

// Constants for valid values
const (
	WrapAuto      Wrap = "auto"      // Default Text Wrapping Around Frame
	WrapNotBeside Wrap = "notBeside" // No Text Wrapping Beside Frame
	WrapAround    Wrap = "around"    // Allow Text Wrapping Around Frame
	WrapTight     Wrap = "tight"     // Tight Text Wrapping Around Frame
	WrapThrough   Wrap = "through"   // Through Text Wrapping Around Frame
	WrapNone      Wrap = "none"      // No Text Wrapping Around Frame
)

// WrapFromStr converts a string to Wrap type.
func WrapFromStr(value string) (Wrap, error) {
	switch value {
	case "auto":
		return WrapAuto, nil
	case "notBeside":
		return WrapNotBeside, nil
	case "around":
		return WrapAround, nil
	case "tight":
		return WrapTight, nil
	case "through":
		return WrapThrough, nil
	case "none":
		return WrapNone, nil
	default:
		return "", errors.New("Invalid Wrap value")
	}
}

// UnmarshalXMLAttr unmarshals XML attribute into Wrap.
func (w *Wrap) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := WrapFromStr(attr.Value)
	if err != nil {
		return err
	}
	*w = val
	return nil
}
