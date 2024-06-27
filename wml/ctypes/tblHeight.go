package ctypes

import (
	"encoding/xml"
	"errors"
	"strconv"

	"github.com/gomutex/godocx/wml/stypes"
)

// TableRowHeight represents the height of a table row in a document.
type TableRowHeight struct {
	Val   *int               `xml:"val,attr,omitempty"`
	HRule *stypes.HeightRule `xml:"hRule,attr,omitempty"`
}

// NewTableRowHeight creates a new TableRowHeight instance.
func NewTableRowHeight(val int, hRule stypes.HeightRule) *TableRowHeight {
	return &TableRowHeight{
		Val:   &val,
		HRule: &hRule,
	}
}

// MarshalXML marshals TableRowHeight to XML.
func (h TableRowHeight) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if h.Val != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: strconv.Itoa(*h.Val)})
	}
	if h.HRule != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:hRule"}, Value: string(*h.HRule)})
	}
	return e.EncodeElement("", start)
}

// UnmarshalXML unmarshals XML into TableRowHeight.
func (h *TableRowHeight) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "val":
			val, err := strconv.Atoi(attr.Value)
			if err != nil {
				return err
			}
			h.Val = &val
		case "hRule":
			hrule, err := stypes.HeightRuleFromStr(attr.Value)
			if err != nil {
				return err
			}
			h.HRule = &hrule
		}
	}
	return d.Skip() // skip any inner elements
}

// HeightRuleFromStr converts a string to HeightRule type.
func HeightRuleFromStr(value string) (stypes.HeightRule, error) {
	switch value {
	case "auto":
		return stypes.HeightRuleAuto, nil
	case "exact":
		return stypes.HeightRuleExact, nil
	case "atLeast":
		return stypes.HeightRuleAtLeast, nil
	default:
		return "", errors.New("Invalid HeightRule value")
	}
}
