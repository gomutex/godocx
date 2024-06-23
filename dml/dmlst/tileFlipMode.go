package dmlst

import (
	"encoding/xml"
	"errors"
)

// TileFlipMode represents tile flip mode values based on the schema.
type TileFlipMode string

// Constants representing valid TileFlipMode values as per the schema.
const (
	TileFlipModeNone       TileFlipMode = "none" // Tile Flip Mode Enum (None)
	TileFlipModeHorizontal TileFlipMode = "x"    // Tile Flip Mode Enum (Horizontal)
	TileFlipModeVertical   TileFlipMode = "y"    // Tile Flip Mode Enum (Vertical)
	TileFlipModeBoth       TileFlipMode = "xy"   // Tile Flip Mode Enum (Horizontal and Vertical)
)

// TileFlipModeFromStr converts a string to TileFlipMode type.
// Returns an error if the string does not match any valid TileFlipMode value.
func TileFlipModeFromStr(value string) (TileFlipMode, error) {
	switch value {
	case "none":
		return TileFlipModeNone, nil
	case "x":
		return TileFlipModeHorizontal, nil
	case "y":
		return TileFlipModeVertical, nil
	case "xy":
		return TileFlipModeBoth, nil
	default:
		return "", errors.New("Invalid TileFlipMode value")
	}
}

// UnmarshalXMLAttr unmarshals XML attribute into TileFlipMode.
// Implements the xml.UnmarshalerAttr interface.
func (t *TileFlipMode) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := TileFlipModeFromStr(attr.Value)
	if err != nil {
		return err
	}
	*t = val
	return nil
}
