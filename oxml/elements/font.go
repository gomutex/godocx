package elements

import (
	"encoding/xml"

	"github.com/gomutex/godocx/constants"
	"github.com/gomutex/godocx/oxml/types"
)

type Font struct {
	Name    string
	Charset Charset
	Family  Family
	Pitch   Pitch
}

func NewFont(name, charset, family string, pitch types.FontPitchType) *Font {
	return &Font{
		Name:    name,
		Charset: Charset{Value: charset},
		Family:  Family{Value: family},
		Pitch:   Pitch{Value: pitch},
	}
}

func (f *Font) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:font"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:name"}, Value: f.Name})

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if err := e.EncodeElement(f.Family, xml.StartElement{Name: xml.Name{Local: "w:family"}}); err != nil {
		return err
	}

	return e.EncodeToken(start.End())

}

// UnmarshalXML unmarshals XML to Font.
func (f *Font) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		for _, attr := range start.Attr {
			switch attr.Name.Local {
			case "name":
				f.Name = attr.Value
			}
		}

		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name {
			case xml.Name{Space: constants.WMLNamespace, Local: "charset"}, xml.Name{Space: constants.AltWMLNamespace, Local: "charset"}:
				if err = d.DecodeElement(&f.Charset, &t); err != nil {
					return err
				}
			case xml.Name{Space: constants.WMLNamespace, Local: "family"}, xml.Name{Space: constants.AltWMLNamespace, Local: "family"}:
				if err = d.DecodeElement(&f.Family, &t); err != nil {
					return err
				}
			case xml.Name{Space: constants.WMLNamespace, Local: "pitch"}, xml.Name{Space: constants.AltWMLNamespace, Local: "pitch"}:
				if err = d.DecodeElement(&f.Pitch, &t); err != nil {
					return err
				}
			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if t == start.End() {
				return nil
			}
		}
	}
}

// Family represents the family of a text or element.
type Family struct {
	Value string
}

// NewFamily creates a new Family instance with the specified family value.
func NewFamily(value string) *Family {
	return &Family{Value: value}
}

// MarshalXML implements the xml.Marshaler interface for the Family type.
//
// It encodes the Family instance into XML using the "w:family" element with a "w:val" attribute.
func (c *Family) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:family"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: c.Value})
	return e.EncodeElement("", start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for the Family type.
//
// It decodes the XML representation of Family, extracting the value from the "val" attribute.
// The inner content of the XML element is skipped.
func (c *Family) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var attr string
	for _, a := range start.Attr {
		if a.Name.Local == "val" {
			attr = a.Value
			break
		}
	}

	c.Value = attr
	return d.Skip() // Skipping the inner content
}

// Charset represents the charset of a text or element.
type Charset struct {
	Value string
}

// NewCharset creates a new Charset instance with the specified charset value.
func NewCharset(value string) *Charset {
	return &Charset{Value: value}
}

// MarshalXML implements the xml.Marshaler interface for the Charset type.
//
// It encodes the Charset instance into XML using the "w:charset" element with a "w:val" attribute.
func (c *Charset) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:charset"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: c.Value})
	return e.EncodeElement("", start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for the Charset type.
//
// It decodes the XML representation of Charset, extracting the value from the "val" attribute.
// The inner content of the XML element is skipped.
func (c *Charset) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var attr string
	for _, a := range start.Attr {
		if a.Name.Local == "val" {
			attr = a.Value
			break
		}
	}

	c.Value = attr
	return d.Skip() // Skipping the inner content
}

// Pitch represents the pitch of a text or element.
type Pitch struct {
	Value types.FontPitchType
}

// NewPitch creates a new Pitch instance with the specified pitch value.
func NewPitch(value types.FontPitchType) *Pitch {
	return &Pitch{Value: value}
}

// MarshalXML implements the xml.Marshaler interface for the Pitch type.
//
// It encodes the Pitch instance into XML using the "w:pitch" element with a "w:val" attribute.
func (c *Pitch) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:pitch"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(c.Value)})
	return e.EncodeElement("", start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for the Pitch type.
//
// It decodes the XML representation of Pitch, extracting the value from the "val" attribute.
// The inner content of the XML element is skipped.
func (c *Pitch) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var attr string
	for _, a := range start.Attr {
		if a.Name.Local == "val" {
			attr = a.Value
			break
		}
	}

	c.Value = types.FontPitchType(attr)
	return d.Skip() // Skipping the inner content
}
