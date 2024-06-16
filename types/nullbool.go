package types

import (
	"encoding/xml"
	"strconv"
)

//
// Based on https://pkg.go.dev/database/sql#NullBool

// NullBool represents a bool that may be null.
// NullBool implements the [Scanner] interface so
// it can be used as a scan destination, similar to [NullString].
type NullBool struct {
	Bool  bool
	Valid bool // Valid is true if Bool is not NULL
}

func NewNullBool(value bool) NullBool {
	return NullBool{
		Bool:  value,
		Valid: true,
	}
}

// NullBoolFromStr creates a new NullBool instance from a string.
// The function accepts a string parameter 'value' which can be either "true", "1", "false", or "0".
// If the input string matches "true" or "1", the function sets the Bool field of the returned NullBool instance to true.
// If the input string matches "false" or "0", the function sets the Bool field of the returned NullBool instance to false.
//
// Example usage:
//
//	nBool := NullBoolFromStr("true")
//	fmt.Println(nBool.Bool) // Output: true
//	fmt.Println(nBool.Valid) // Output: true
//
//	nBool = NullBoolFromStr("0")
//	fmt.Println(nBool.Bool) // Output: false
//	fmt.Println(nBool.Valid) // Output: true
func NullBoolFromStr(value string) NullBool {
	nBool := NullBool{Valid: true}
	if value == "true" || value == "1" {
		nBool.Bool = true
	}

	return nBool
}

// ToIntFlag returns the integer representation of the NullBool.
// If the Bool field is true, it returns 1.
// If the Bool field is false, it also returns 0.
// Note: This method does not consider the Valid field. You ensure to check Valid field before calling this if you want optional field
func (n NullBool) ToIntFlag() int {
	if n.Bool {
		return 1
	}

	return 0
}

// ToIntFlag returns the string representation of the NullBool.
// If the Bool field is true, it returns "1".
// If the Bool field is false, it also returns "0".
// Note: This method does not consider the Valid field. You ensure to check Valid field before calling this if you want optional field
func (n NullBool) ToStringFlag() string {
	if n.Bool {
		return "1"
	}
	return "0"
}

func (n *NullBool) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := strconv.ParseBool(attr.Value)
	if err != nil {
		return err
	}

	n.Bool = val
	n.Valid = true
	return nil

}
