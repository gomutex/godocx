package dmlst

import (
	"encoding/xml"
	"errors"
)

// WrapText type
type WrapText string

// Constants for valid values
const (
	WrapTextBothSides WrapText = "bothSides" // Both Sides
	WrapTextLeft      WrapText = "left"      // Left Side Only
	WrapTextRight     WrapText = "right"     // Right Side Only
	WrapTextLargest   WrapText = "largest"   // Largest Side Only
)

// WrapTextFromStr converts a string to WrapText type.
func WrapTextFromStr(value string) (WrapText, error) {
	switch value {
	case "bothSides":
		return WrapTextBothSides, nil
	case "left":
		return WrapTextLeft, nil
	case "right":
		return WrapTextRight, nil
	case "largest":
		return WrapTextLargest, nil
	default:
		return "", errors.New("Invalid WrapText value")
	}
}

// UnmarshalXMLAttr unmarshals XML attribute into WrapText.
func (w *WrapText) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := WrapTextFromStr(attr.Value)
	if err != nil {
		return err
	}
	*w = val
	return nil
}
