package stypes

import (
	"encoding/xml"
	"errors"
)

type FontTypeHint string

const (
	FontTypeHintDefault  FontTypeHint = "default"  // High ANSI Font
	FontTypeHintEastAsia FontTypeHint = "eastAsia" // East Asian Font
	FontTypeHintCS       FontTypeHint = "cs"       // Complex Script Font
)

func FontTypeHintFromStr(value string) (FontTypeHint, error) {
	switch value {
	case "default":
		return FontTypeHintDefault, nil
	case "eastAsia":
		return FontTypeHintEastAsia, nil
	case "cs":
		return FontTypeHintCS, nil
	default:
		return "", errors.New("invalid FontTypeHint value")
	}
}

func (d *FontTypeHint) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := FontTypeHintFromStr(attr.Value)
	if err != nil {
		return err
	}

	*d = val

	return nil
}
