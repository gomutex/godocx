package dmlst

import (
	"encoding/xml"
	"errors"
)

type RelFromH string

const (
	RelFromHCharacter     RelFromH = "character"
	RelFromHColumn        RelFromH = "column"
	RelFromHInsideMargin  RelFromH = "insideMargin"
	RelFromHLeftMargin    RelFromH = "leftMargin"
	RelFromHMargin        RelFromH = "margin"
	RelFromHOutsizeMargin RelFromH = "outsizeMargin"
	RelFromHPage          RelFromH = "page"
	RelFromHRightMargin   RelFromH = "rightMargin"
)

type RelFromV string

const (
	RelFromVBottomMargin  RelFromV = "bottomMargin"
	RelFromVInsideMargin  RelFromV = "insideMargin"
	RelFromVLine          RelFromV = "line"
	RelFromVMargin        RelFromV = "margin"
	RelFromVOutsizeMargin RelFromV = "outsizeMargin"
	RelFromVPage          RelFromV = "page"
	RelFromVParagraph     RelFromV = "paragraph"
	RelFromVTopMargin     RelFromV = "topMargin"
)

// RelFromHFromStr converts a string to RelFromH type.
func RelFromHFromStr(value string) (RelFromH, error) {
	switch value {
	case "character":
		return RelFromHCharacter, nil
	case "column":
		return RelFromHColumn, nil
	case "insideMargin":
		return RelFromHInsideMargin, nil
	case "leftMargin":
		return RelFromHLeftMargin, nil
	case "margin":
		return RelFromHMargin, nil
	case "outsizeMargin":
		return RelFromHOutsizeMargin, nil
	case "page":
		return RelFromHPage, nil
	case "rightMargin":
		return RelFromHRightMargin, nil
	default:
		return "", errors.New("Invalid RelFromH value")
	}
}

// UnmarshalXMLAttr unmarshals XML attribute into RelFromH.
func (r *RelFromH) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := RelFromHFromStr(attr.Value)
	if err != nil {
		return err
	}
	*r = val
	return nil
}

// RelFromVFromStr converts a string to RelFromV type.
func RelFromVFromStr(value string) (RelFromV, error) {
	switch value {
	case "bottomMargin":
		return RelFromVBottomMargin, nil
	case "insideMargin":
		return RelFromVInsideMargin, nil
	case "line":
		return RelFromVLine, nil
	case "margin":
		return RelFromVMargin, nil
	case "outsizeMargin":
		return RelFromVOutsizeMargin, nil
	case "page":
		return RelFromVPage, nil
	case "paragraph":
		return RelFromVParagraph, nil
	case "topMargin":
		return RelFromVTopMargin, nil
	default:
		return "", errors.New("Invalid RelFromV value")
	}
}

// UnmarshalXMLAttr unmarshals XML attribute into RelFromV.
func (r *RelFromV) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := RelFromVFromStr(attr.Value)
	if err != nil {
		return err
	}
	*r = val
	return nil
}
