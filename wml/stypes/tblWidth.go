package stypes

import (
	"encoding/xml"
	"errors"
)

type TableWidth string

const (
	TableWidthDxa         TableWidth = "dxa"  //Width in Twentieths of a Point
	TableWidthAuto        TableWidth = "auto" //Automatically Determined Width
	TableWidthPct         TableWidth = "pct"  //Width in Fiftieths of a Percent
	TableWidthNil         TableWidth = "nil"  //No Width
	TableWidthUnsupported TableWidth = ""
)

func TableWidthFromStr(value string) (TableWidth, error) {
	switch value {
	case "dxa":
		return TableWidthDxa, nil
	case "auto":
		return TableWidthAuto, nil
	case "pct":
		return TableWidthPct, nil
	case "nil":
		return TableWidthNil, nil
	default:
		return "", errors.New("Invalid TableWidth value")
	}
}

func (to *TableWidth) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := TableWidthFromStr(attr.Value)
	if err != nil {
		return err
	}
	*to = val
	return nil
}
