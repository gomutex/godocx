package stypes

import (
	"encoding/xml"
	"errors"
)

// Document Grid Types
type DocGridType string

const (
	DocGridDefault       DocGridType = "default"       //No Document Grid
	DocGridLines         DocGridType = "lines"         //Line Grid Only
	DocGridLinesAndChars DocGridType = "linesAndChars" //Line and Character Grid
	DocGridSnapToChars   DocGridType = "snapToChars"   //Character Grid Only
)

func DocGridTypeFromStr(value string) (DocGridType, error) {
	switch value {
	case "default":
		return DocGridDefault, nil
	case "lines":
		return DocGridLines, nil
	case "linesAndChars":
		return DocGridLinesAndChars, nil
	case "snapToChars":
		return DocGridSnapToChars, nil
	default:
		return "", errors.New("Invalid Docgrid Type")

	}
}

func (d *DocGridType) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := DocGridTypeFromStr(attr.Value)
	if err != nil {
		return err
	}

	*d = val

	return nil

}
