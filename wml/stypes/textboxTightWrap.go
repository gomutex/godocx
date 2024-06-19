package stypes

import (
	"encoding/xml"
	"errors"
)

type TextboxTightWrap string

const (
	TextboxTightWrapNone             TextboxTightWrap = "none"             // Do Not Tight Wrap
	TextboxTightWrapAllLines         TextboxTightWrap = "allLines"         // Tight Wrap All Lines
	TextboxTightWrapFirstAndLastLine TextboxTightWrap = "firstAndLastLine" // Tight Wrap First and Last Lines
	TextboxTightWrapFirstLineOnly    TextboxTightWrap = "firstLineOnly"    // Tight Wrap First Line
	TextboxTightWrapLastLineOnly     TextboxTightWrap = "lastLineOnly"     // Tight Wrap Last Line
)

func TextboxTightWrapFromStr(value string) (TextboxTightWrap, error) {
	switch value {
	case "none":
		return TextboxTightWrapNone, nil
	case "allLines":
		return TextboxTightWrapAllLines, nil
	case "firstAndLastLine":
		return TextboxTightWrapFirstAndLastLine, nil
	case "firstLineOnly":
		return TextboxTightWrapFirstLineOnly, nil
	case "lastLineOnly":
		return TextboxTightWrapLastLineOnly, nil
	default:
		return "", errors.New("Invalid Textbox Tight Wrap value")
	}
}

func (t *TextboxTightWrap) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := TextboxTightWrapFromStr(attr.Value)
	if err != nil {
		return err
	}

	*t = val

	return nil
}
