package stypes

import (
	"encoding/xml"
	"errors"
)

type TextAlign string

const (
	TextAlignTop      TextAlign = "top"      // Align Text at Top
	TextAlignCenter   TextAlign = "center"   // Align Text at Center
	TextAlignBaseline TextAlign = "baseline" // Align Text at Baseline
	TextAlignBottom   TextAlign = "bottom"   // Align Text at Bottom
	TextAlignAuto     TextAlign = "auto"     // Automatically Determine Alignment
)

func TextAlignFromStr(value string) (TextAlign, error) {
	switch value {
	case "top":
		return TextAlignTop, nil
	case "center":
		return TextAlignCenter, nil
	case "baseline":
		return TextAlignBaseline, nil
	case "bottom":
		return TextAlignBottom, nil
	case "auto":
		return TextAlignAuto, nil
	default:
		return "", errors.New("Invalid Text Alignment")
	}
}

func (a *TextAlign) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := TextAlignFromStr(attr.Value)
	if err != nil {
		return err
	}

	*a = val

	return nil
}
