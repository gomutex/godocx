package stypes

import (
	"encoding/xml"
	"errors"
)

type StyleType string

const (
	StyleTypeParagraph StyleType = "paragraph"
	StyleTypeCharacter StyleType = "character"
	StyleTypeTable     StyleType = "table"
	StyleTypeNumbering StyleType = "numbering"
)

func StyleTypeFromStr(value string) (StyleType, error) {
	switch value {
	case "paragraph":
		return StyleTypeParagraph, nil
	case "character":
		return StyleTypeCharacter, nil
	case "table":
		return StyleTypeTable, nil
	case "numbering":
		return StyleTypeNumbering, nil
	default:
		return "", errors.New("invalid StyleType value")
	}
}

func (s *StyleType) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := StyleTypeFromStr(attr.Value)
	if err != nil {
		return err
	}

	*s = val

	return nil
}
