package stypes

import (
	"encoding/xml"
	"errors"
)

type SectionMark string

const (
	SectionMarkNextPage       SectionMark = "nextPage"   //Next Page Section Break
	SectionMarkNextColumn     SectionMark = "nextColumn" //Column Section Break
	SectionMarkNextContinuous SectionMark = "continuous" //Continuous Section Break
	SectionMarkEvenPage       SectionMark = "evenPage"   //Even Page Section Break
	SectionMarkOddPage        SectionMark = "oddPage"    //Odd Page Section Break
)

func SectionMarkFromStr(value string) (SectionMark, error) {
	switch value {
	case "nextPage":
		return SectionMarkNextPage, nil
	case "nextColumn":
		return SectionMarkNextColumn, nil
	case "continuous":
		return SectionMarkNextContinuous, nil
	case "evenPage":
		return SectionMarkEvenPage, nil
	case "oddPage":
		return SectionMarkOddPage, nil
	default:
		return "", errors.New("Invalid Section Mark")
	}
}

func (d *SectionMark) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := SectionMarkFromStr(attr.Value)
	if err != nil {
		return err
	}

	*d = val

	return nil

}
