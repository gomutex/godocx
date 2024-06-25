package stypes

import (
	"encoding/xml"
	"errors"
)

type MergeCell string

const (
	MergeCellContinue MergeCell = "continue" // Continue Merged Region
	MergeCellRestart  MergeCell = "restart"  // Start/Restart Merged Region
)

func MergeCellFromStr(value string) (MergeCell, error) {
	switch value {
	case "continue":
		return MergeCellContinue, nil
	case "restart":
		return MergeCellRestart, nil
	default:
		return "", errors.New("invalid MergeCell value")
	}
}

func (m *MergeCell) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := MergeCellFromStr(attr.Value)
	if err != nil {
		return err
	}

	*m = val

	return nil
}
