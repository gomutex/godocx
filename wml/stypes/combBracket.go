package stypes

import (
	"encoding/xml"
	"errors"
)

// CombineBrackets Type
type CombineBrackets string

const (
	CombineBracketsNone   CombineBrackets = "none"   // No Enclosing Brackets
	CombineBracketsRound  CombineBrackets = "round"  // Round Brackets
	CombineBracketsSquare CombineBrackets = "square" // Square Brackets
	CombineBracketsAngle  CombineBrackets = "angle"  // Angle Brackets
	CombineBracketsCurly  CombineBrackets = "curly"  // Curly Brackets
)

func CombineBracketsFromStr(value string) (CombineBrackets, error) {
	switch value {
	case "none":
		return CombineBracketsNone, nil
	case "round":
		return CombineBracketsRound, nil
	case "square":
		return CombineBracketsSquare, nil
	case "angle":
		return CombineBracketsAngle, nil
	case "curly":
		return CombineBracketsCurly, nil
	default:
		return "", errors.New("invalid CombineBrackets value")
	}
}

func (c *CombineBrackets) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := CombineBracketsFromStr(attr.Value)
	if err != nil {
		return err
	}

	*c = val

	return nil
}
