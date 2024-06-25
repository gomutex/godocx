package stypes

import (
	"encoding/xml"
	"fmt"
)

// Vertical Alignment Type
type VerticalJc string

const (
	VerticalJcTop    VerticalJc = "top"
	VerticalJcCenter VerticalJc = "center"
	VerticalJcBoth   VerticalJc = "both"
	VerticalJcBottom VerticalJc = "bottom"
)

// MarshalXMLAttr marshals the VerticalJc type as an XML attribute.
func (v VerticalJc) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: string(v)}, nil
}

// UnmarshalXMLAttr unmarshals an XML attribute into a VerticalJc type.
func (v *VerticalJc) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "top", "center", "both", "bottom":
		*v = VerticalJc(attr.Value)
	default:
		return fmt.Errorf("unexpected value for VerticalJc: %s", attr.Value)
	}
	return nil
}
