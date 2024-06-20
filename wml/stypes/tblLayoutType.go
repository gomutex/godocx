package stypes

import (
	"encoding/xml"
	"errors"
)

// Table Layout Type
type TableLayout string

const (
	TableLayoutFixed   TableLayout = "fixed"
	TableLayoutAutoFit TableLayout = "autofit"
	TableLayoutInvalid TableLayout = ""
)

func TableLayoutFromStr(val string) (TableLayout, error) {
	switch val {
	case "fixed":
		return TableLayoutFixed, nil
	case "autofit":
		return TableLayoutAutoFit, nil
	default:
		return TableLayoutInvalid, errors.New("Invalid Table Layout Type")
	}
}

func (t *TableLayout) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := TableLayoutFromStr(attr.Value)
	if err != nil {
		return err
	}

	*t = val

	return nil
}
