package stypes

import (
	"encoding/xml"
	"errors"
)

// CustTabStop represents the custom tab stop type.
type CustTabStop string

const (
	CustTabStopClear   CustTabStop = "clear"
	CustTabStopLeft    CustTabStop = "left"
	CustTabStopCenter  CustTabStop = "center"
	CustTabStopRight   CustTabStop = "right"
	CustTabStopDecimal CustTabStop = "decimal"
	CustTabStopBar     CustTabStop = "bar"
	CustTabStopNum     CustTabStop = "num"
	CustTabStopInvalid CustTabStop = ""
)

// Function to convert string to CustTabStop.
func CustTabStopFromStr(val string) (CustTabStop, error) {
	switch val {
	case "clear":
		return CustTabStopClear, nil
	case "left":
		return CustTabStopLeft, nil
	case "center":
		return CustTabStopCenter, nil
	case "right":
		return CustTabStopRight, nil
	case "decimal":
		return CustTabStopDecimal, nil
	case "bar":
		return CustTabStopBar, nil
	case "num":
		return CustTabStopNum, nil
	default:
		return CustTabStopInvalid, errors.New("Invalid CustTabStop value")
	}
}

// Method to unmarshal XML attribute into CustTabStop.
func (t *CustTabStop) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := CustTabStopFromStr(attr.Value)
	if err != nil {
		return err
	}
	*t = val
	return nil
}
