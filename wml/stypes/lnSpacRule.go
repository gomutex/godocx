package stypes

import (
	"encoding/xml"
	"errors"
)

type LineSpacingRule string

const (
	LineSpacingRuleAuto    LineSpacingRule = "auto"    // Automatically Determined Line Height
	LineSpacingRuleExact   LineSpacingRule = "exact"   // Exact Line Height
	LineSpacingRuleAtLeast LineSpacingRule = "atLeast" // Minimum Line Height
)

func LineSpacingRuleFromStr(value string) (LineSpacingRule, error) {
	switch value {
	case "auto":
		return LineSpacingRuleAuto, nil
	case "exact":
		return LineSpacingRuleExact, nil
	case "atLeast":
		return LineSpacingRuleAtLeast, nil
	default:
		return "", errors.New("invalid LineSpacingRule value")
	}
}

func (d *LineSpacingRule) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := LineSpacingRuleFromStr(attr.Value)
	if err != nil {
		return err
	}

	*d = val

	return nil
}
