package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// Optional Bool Element: Helper element that only has one attribute which is optional
type OnOff struct {
	Val *stypes.OnOff `xml:"val,attr,omitempty"`
}

func OnOffFromBool(value bool) *OnOff {
	o := stypes.OnOffFalse
	if value {
		o = stypes.OnOffTrue
	}

	return &OnOff{
		Val: &o,
	}
}

func OnOffFromStr(value string) (*OnOff, error) {
	o, err := stypes.OnOffFromStr(value)
	if err != nil {
		return nil, err
	}
	return &OnOff{
		Val: &o,
	}, nil
}

// Disable sets the value to false and valexists true
func (n *OnOff) Disable() {
	o := stypes.OnOffFalse
	n.Val = &o
}

// MarshalXML implements the xml.Marshaler interface for the Bold type.
// It encodes the instance into XML using the "w:XMLName" element with a "w:val" attribute.
func (n OnOff) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if n.Val != nil { // Add val attribute only if the val exists
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(*n.Val)})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
