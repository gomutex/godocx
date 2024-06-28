package stypes

import (
	"encoding/xml"
	"errors"
)

// Underline represents different styles for underline.
type Underline string

const (
	UnderlineNone            Underline = "none"
	UnderlineSingle          Underline = "single"
	UnderlineWords           Underline = "words"
	UnderlineDouble          Underline = "double"
	UnderlineDotted          Underline = "dotted"
	UnderlineThick           Underline = "thick"
	UnderlineDash            Underline = "dash"
	UnderlineDotDash         Underline = "dotDash"
	UnderlineDotDotDash      Underline = "dotDotDash"
	UnderlineWavy            Underline = "wavy"
	UnderlineDottedHeavy     Underline = "dottedHeavy"
	UnderlineDashHeavy       Underline = "dashHeavy"
	UnderlineDotDashHeavy    Underline = "dotDashHeavy"
	UnderlineDotDotDashHeavy Underline = "dotDotDashHeavy"
	UnderlineWavyHeavy       Underline = "wavyHeavy"
	UnderlineDashLong        Underline = "dashLong"
	UnderlineWavyDouble      Underline = "wavyDouble"
	UnderlineDashLongHeavy   Underline = "dashLongHeavy"
)

func UnderlineFromStr(value string) (Underline, error) {
	switch value {
	case "none":
		return UnderlineNone, nil
	case "single":
		return UnderlineSingle, nil
	case "words":
		return UnderlineWords, nil
	case "double":
		return UnderlineDouble, nil
	case "dotted":
		return UnderlineDotted, nil
	case "thick":
		return UnderlineThick, nil
	case "dash":
		return UnderlineDash, nil
	case "dotDash":
		return UnderlineDotDash, nil
	case "dotDotDash":
		return UnderlineDotDotDash, nil
	case "wavy":
		return UnderlineWavy, nil
	case "dottedHeavy":
		return UnderlineDottedHeavy, nil
	case "dashHeavy":
		return UnderlineDashHeavy, nil
	case "dotDashHeavy":
		return UnderlineDotDashHeavy, nil
	case "dotDotDashHeavy":
		return UnderlineDotDotDashHeavy, nil
	case "wavyHeavy":
		return UnderlineWavyHeavy, nil
	case "dashLong":
		return UnderlineDashLong, nil
	case "wavyDouble":
		return UnderlineWavyDouble, nil
	case "dashLongHeavy":
		return UnderlineDashLongHeavy, nil
	default:
		return "", errors.New("invalid Underline value")
	}
}

func (u *Underline) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := UnderlineFromStr(attr.Value)
	if err != nil {
		return err
	}

	*u = val

	return nil
}
