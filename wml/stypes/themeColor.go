package stypes

import (
	"encoding/xml"
	"errors"
)

// ThemeColor represents the possible values for Theme Color in WordprocessingML.
type ThemeColor string

// List of valid ThemeColor values.
const (
	ThemeColorDark1             ThemeColor = "dark1"
	ThemeColorLight1            ThemeColor = "light1"
	ThemeColorDark2             ThemeColor = "dark2"
	ThemeColorLight2            ThemeColor = "light2"
	ThemeColorAccent1           ThemeColor = "accent1"
	ThemeColorAccent2           ThemeColor = "accent2"
	ThemeColorAccent3           ThemeColor = "accent3"
	ThemeColorAccent4           ThemeColor = "accent4"
	ThemeColorAccent5           ThemeColor = "accent5"
	ThemeColorAccent6           ThemeColor = "accent6"
	ThemeColorHyperlink         ThemeColor = "hyperlink"
	ThemeColorFollowedHyperlink ThemeColor = "followedHyperlink"
	ThemeColorNone              ThemeColor = "none"
	ThemeColorBackground1       ThemeColor = "background1"
	ThemeColorText1             ThemeColor = "text1"
	ThemeColorBackground2       ThemeColor = "background2"
	ThemeColorText2             ThemeColor = "text2"
)

// ThemeColorFromStr converts a string value to ThemeColor.
func ThemeColorFromStr(value string) (ThemeColor, error) {
	switch value {
	case "dark1":
		return ThemeColorDark1, nil
	case "light1":
		return ThemeColorLight1, nil
	case "dark2":
		return ThemeColorDark2, nil
	case "light2":
		return ThemeColorLight2, nil
	case "accent1":
		return ThemeColorAccent1, nil
	case "accent2":
		return ThemeColorAccent2, nil
	case "accent3":
		return ThemeColorAccent3, nil
	case "accent4":
		return ThemeColorAccent4, nil
	case "accent5":
		return ThemeColorAccent5, nil
	case "accent6":
		return ThemeColorAccent6, nil
	case "hyperlink":
		return ThemeColorHyperlink, nil
	case "followedHyperlink":
		return ThemeColorFollowedHyperlink, nil
	case "none":
		return ThemeColorNone, nil
	case "background1":
		return ThemeColorBackground1, nil
	case "text1":
		return ThemeColorText1, nil
	case "background2":
		return ThemeColorBackground2, nil
	case "text2":
		return ThemeColorText2, nil
	default:
		return "", errors.New("invalid ThemeColor value")
	}
}

// UnmarshalXMLAttr unmarshals an XML attribute into a ThemeColor.
func (t *ThemeColor) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := ThemeColorFromStr(attr.Value)
	if err != nil {
		return err
	}

	*t = val

	return nil
}
