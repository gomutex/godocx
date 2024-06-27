package stypes

import (
	"encoding/hex"
	"encoding/xml"
	"errors"
	"strings"
)

type LongHexNum string

// LongHexNumberFromStr validates and converts a string to a LongHexNumber type
func LongHexNumFromStr(s string) (LongHexNum, error) {
	// Allow an empty value as valid
	if s == "" {
		return LongHexNum(s), nil
	}

	// // Ensure the string is exactly 4 characters long
	// if len(s) != 4 {
	// 	return "", errors.New("invalid LongHexNumber length")
	// }

	// Ensure the string contains valid hexadecimal characters
	if _, err := hex.DecodeString(s); err != nil {
		return "", errors.New("invalid LongHexNumber format")
	}

	return LongHexNum(strings.ToUpper(s)), nil
}

func (d *LongHexNum) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := LongHexNumFromStr(attr.Value)
	if err != nil {
		return err
	}

	*d = val

	return nil

}
