package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// Justification represents the justification of a paragraph.
type Justification struct {
	Val stypes.Justification `xml:"val,attr"`
}

// NewJustification creates a new Justification.
func NewJustification(val string) (*Justification, error) {
	jc, err := stypes.JustificationFromStr(val)
	if err != nil {
		return nil, err
	}
	return &Justification{Val: jc}, nil
}

// DefaultJustification creates the default Justification with the value "centerGroup".
func DefaultJustification() *Justification {
	return &Justification{Val: stypes.JustificationCenter}
}

func (j *Justification) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(j.Val)})
	return e.EncodeElement("", start)
}
