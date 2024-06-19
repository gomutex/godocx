package stypes

import (
	"encoding/xml"
	"errors"
)

// Anchor Type
type Anchor string

// Constants for valid values
const (
	AnchorText   Anchor = "text"   // Relative to Text Extents
	AnchorMargin Anchor = "margin" // Relative To Margin
	AnchorPage   Anchor = "page"   // Relative to Page
)

// AnchorFromStr converts a string to Anchor type.
func AnchorFromStr(value string) (Anchor, error) {
	switch value {
	case "text":
		return AnchorText, nil
	case "margin":
		return AnchorMargin, nil
	case "page":
		return AnchorPage, nil
	default:
		return "", errors.New("Invalid Anchor value")
	}
}

// UnmarshalXMLAttr unmarshals XML attribute into Anchor.
func (h *Anchor) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := AnchorFromStr(attr.Value)
	if err != nil {
		return err
	}
	*h = val
	return nil
}
