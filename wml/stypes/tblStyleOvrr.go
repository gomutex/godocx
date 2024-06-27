package stypes

import (
	"encoding/xml"
	"errors"
)

type TblStyleOverrideType string

const (
	TblStyleOverrideWholeTable TblStyleOverrideType = "wholeTable"
	TblStyleOverrideFirstRow   TblStyleOverrideType = "firstRow"
	TblStyleOverrideLastRow    TblStyleOverrideType = "lastRow"
	TblStyleOverrideFirstCol   TblStyleOverrideType = "firstCol"
	TblStyleOverrideLastCol    TblStyleOverrideType = "lastCol"
	TblStyleOverrideBand1Vert  TblStyleOverrideType = "band1Vert"
	TblStyleOverrideBand2Vert  TblStyleOverrideType = "band2Vert"
	TblStyleOverrideBand1Horz  TblStyleOverrideType = "band1Horz"
	TblStyleOverrideBand2Horz  TblStyleOverrideType = "band2Horz"
	TblStyleOverrideNeCell     TblStyleOverrideType = "neCell"
	TblStyleOverrideNwCell     TblStyleOverrideType = "nwCell"
	TblStyleOverrideSeCell     TblStyleOverrideType = "seCell"
	TblStyleOverrideSwCell     TblStyleOverrideType = "swCell"
)

func TblStyleOverrideTypeFromStr(value string) (TblStyleOverrideType, error) {
	switch value {
	case "wholeTable":
		return TblStyleOverrideWholeTable, nil
	case "firstRow":
		return TblStyleOverrideFirstRow, nil
	case "lastRow":
		return TblStyleOverrideLastRow, nil
	case "firstCol":
		return TblStyleOverrideFirstCol, nil
	case "lastCol":
		return TblStyleOverrideLastCol, nil
	case "band1Vert":
		return TblStyleOverrideBand1Vert, nil
	case "band2Vert":
		return TblStyleOverrideBand2Vert, nil
	case "band1Horz":
		return TblStyleOverrideBand1Horz, nil
	case "band2Horz":
		return TblStyleOverrideBand2Horz, nil
	case "neCell":
		return TblStyleOverrideNeCell, nil
	case "nwCell":
		return TblStyleOverrideNwCell, nil
	case "seCell":
		return TblStyleOverrideSeCell, nil
	case "swCell":
		return TblStyleOverrideSwCell, nil
	default:
		return "", errors.New("invalid TblStyleOverrideType value")
	}
}

func (t *TblStyleOverrideType) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := TblStyleOverrideTypeFromStr(attr.Value)
	if err != nil {
		return err
	}

	*t = val

	return nil
}
