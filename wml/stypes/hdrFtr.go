package stypes

import (
	"encoding/xml"
	"errors"
)

// Header or Footer Type
type HdrFtrType string

const (
	HdrFtrEven    HdrFtrType = "even"    //Even Numbered Pages Only
	HdrFtrDefault HdrFtrType = "default" //Default Header or Footer
	HdrFtrFirst   HdrFtrType = "first"   //First Page Only
)

func HdrFtrFromStr(value string) (HdrFtrType, error) {
	switch value {
	case "default":
		return HdrFtrDefault, nil
	case "even":
		return HdrFtrEven, nil
	case "first":
		return HdrFtrFirst, nil
	default:
		return "", errors.New("Invalid Header or Footer Type")

	}
}

func (d *HdrFtrType) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := HdrFtrFromStr(attr.Value)
	if err != nil {
		return err
	}

	*d = val

	return nil

}
