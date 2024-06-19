package stypes

import (
	"encoding/xml"
	"errors"
)

// HeightRule Type
type HeightRule string

// Constants for valid values
const (
	HeightRuleAuto    HeightRule = "auto"    // Determine Height Based On Contents
	HeightRuleExact   HeightRule = "exact"   // Exact Height
	HeightRuleAtLeast HeightRule = "atLeast" // Minimum Height
)

// HeightRuleFromStr converts a string to HeightRule type.
func HeightRuleFromStr(value string) (HeightRule, error) {
	switch value {
	case "auto":
		return HeightRuleAuto, nil
	case "exact":
		return HeightRuleExact, nil
	case "atLeast":
		return HeightRuleAtLeast, nil
	default:
		return "", errors.New("Invalid HeightRule value")
	}
}

// UnmarshalXMLAttr unmarshals XML attribute into HeightRule.
func (h *HeightRule) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := HeightRuleFromStr(attr.Value)
	if err != nil {
		return err
	}
	*h = val
	return nil
}
