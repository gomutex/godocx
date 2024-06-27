package dmlst

import (
	"encoding/xml"
	"strconv"
)

//
// Based on https://pkg.go.dev/database/sql#OptBool

// OptBool represents a bool that may be null.
// OptBool implements the [Scanner] interface so
// it can be used as a scan destination, similar to [NullString].
type OptBool struct {
	Bool  bool
	Valid bool // Valid is true if Bool is not NULL
}

func NewOptBool(value bool) OptBool {
	return OptBool{
		Bool:  value,
		Valid: true,
	}
}

// OptBoolFromStr creates a new OptBool instance from a string.
// The function accepts a string parameter 'value' which can be either "true", "1", "false", or "0".
// If the input string matches "true" or "1", the function sets the Bool field of the returned OptBool instance to true.
// If the input string matches "false" or "0", the function sets the Bool field of the returned OptBool instance to false.
//
// Example usage:
//
//	nBool := OptBoolFromStr("true")
//	fmt.Println(nBool.Bool) // Output: true
//	fmt.Println(nBool.Valid) // Output: true
//
//	nBool = OptBoolFromStr("0")
//	fmt.Println(nBool.Bool) // Output: false
//	fmt.Println(nBool.Valid) // Output: true
func OptBoolFromStr(value string) OptBool {
	nBool := OptBool{Valid: true}
	if value == "true" || value == "1" {
		nBool.Bool = true
	}

	return nBool
}

// ToIntFlag returns the integer representation of the OptBool.
// If the Bool field is true, it returns 1.
// If the Bool field is false, it also returns 0.
// Note: This method does not consider the Valid field. You ensure to check Valid field before calling this if you want optional field
func (n OptBool) ToIntFlag() int {
	if n.Bool {
		return 1
	}

	return 0
}

// ToIntFlag returns the string representation of the OptBool.
// If the Bool field is true, it returns "1".
// If the Bool field is false, it also returns "0".
// Note: This method does not consider the Valid field. You ensure to check Valid field before calling this if you want optional field
func (n OptBool) ToStringFlag() string {
	if n.Bool {
		return "1"
	}
	return "0"
}

func (n *OptBool) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := strconv.ParseBool(attr.Value)
	if err != nil {
		return err
	}

	n.Bool = val
	n.Valid = true
	return nil

}
