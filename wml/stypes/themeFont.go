package stypes

import (
	"encoding/xml"
	"errors"
)

type ThemeFont string

const (
	ThemeFontMajorEastAsia ThemeFont = "majorEastAsia" // Major East Asian Theme Font
	ThemeFontMajorBidi     ThemeFont = "majorBidi"     // Major Complex Script Theme Font
	ThemeFontMajorAscii    ThemeFont = "majorAscii"    // Major ASCII Theme Font
	ThemeFontMajorHAnsi    ThemeFont = "majorHAnsi"    // Major High ANSI Theme Font
	ThemeFontMinorEastAsia ThemeFont = "minorEastAsia" // Minor East Asian Theme Font
	ThemeFontMinorBidi     ThemeFont = "minorBidi"     // Minor Complex Script Theme Font
	ThemeFontMinorAscii    ThemeFont = "minorAscii"    // Minor ASCII Theme Font
	ThemeFontMinorHAnsi    ThemeFont = "minorHAnsi"    // Minor High ANSI Theme Font
)

func ThemeFontFromStr(value string) (ThemeFont, error) {
	switch value {
	case "majorEastAsia":
		return ThemeFontMajorEastAsia, nil
	case "majorBidi":
		return ThemeFontMajorBidi, nil
	case "majorAscii":
		return ThemeFontMajorAscii, nil
	case "majorHAnsi":
		return ThemeFontMajorHAnsi, nil
	case "minorEastAsia":
		return ThemeFontMinorEastAsia, nil
	case "minorBidi":
		return ThemeFontMinorBidi, nil
	case "minorAscii":
		return ThemeFontMinorAscii, nil
	case "minorHAnsi":
		return ThemeFontMinorHAnsi, nil
	default:
		return "", errors.New("invalid ThemeFont value")
	}
}

func (d *ThemeFont) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := ThemeFontFromStr(attr.Value)
	if err != nil {
		return err
	}

	*d = val

	return nil
}
