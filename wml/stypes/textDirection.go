package stypes

import (
	"encoding/xml"
	"errors"
)

type TextDirection string

const (
	TextDirectionLrTb  TextDirection = "lrTb"  // Left to Right, Top to Bottom
	TextDirectionTbRl  TextDirection = "tbRl"  // Top to Bottom, Right to Left
	TextDirectionBtLr  TextDirection = "btLr"  // Bottom to Top, Left to Right
	TextDirectionLrTbV TextDirection = "lrTbV" // Left to Right, Top to Bottom Rotated
	TextDirectionTbRlV TextDirection = "tbRlV" // Top to Bottom, Right to Left Rotated
	TextDirectionTbLrV TextDirection = "tbLrV" // Top to Bottom, Left to Right Rotated
)

func TextDirectionFromStr(value string) (TextDirection, error) {
	switch value {
	case "lrTb":
		return TextDirectionLrTb, nil
	case "tbRl":
		return TextDirectionTbRl, nil
	case "btLr":
		return TextDirectionBtLr, nil
	case "lrTbV":
		return TextDirectionLrTbV, nil
	case "tbRlV":
		return TextDirectionTbRlV, nil
	case "tbLrV":
		return TextDirectionTbLrV, nil
	default:
		return "", errors.New("Invalid Text Direction")
	}
}

func (d *TextDirection) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := TextDirectionFromStr(attr.Value)
	if err != nil {
		return err
	}

	*d = val

	return nil
}
