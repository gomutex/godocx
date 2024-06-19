package stypes

import (
	"encoding/xml"
	"errors"
)

// VerticalAlignRun Type
type VerticalAlignRun string

const (
	VerticalAlignRunBaseline    VerticalAlignRun = "baseline"    // Regular Vertical Positioning
	VerticalAlignRunSuperscript VerticalAlignRun = "superscript" // Superscript
	VerticalAlignRunSubscript   VerticalAlignRun = "subscript"   // Subscript
)

func VerticalAlignRunFromStr(value string) (VerticalAlignRun, error) {
	switch value {
	case "baseline":
		return VerticalAlignRunBaseline, nil
	case "superscript":
		return VerticalAlignRunSuperscript, nil
	case "subscript":
		return VerticalAlignRunSubscript, nil
	default:
		return "", errors.New("Invalid VerticalAlignRun Type")
	}
}

func (v *VerticalAlignRun) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := VerticalAlignRunFromStr(attr.Value)
	if err != nil {
		return err
	}

	*v = val

	return nil
}
