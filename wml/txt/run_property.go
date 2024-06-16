package txt

import (
	"encoding/xml"

	"github.com/gomutex/godocx/common/constants"
	"github.com/gomutex/godocx/elemtypes"
	"github.com/gomutex/godocx/wml/formatting"
)

// RunProperty represents the properties of a run of text within a paragraph.
type RunProperty struct {
	Color     *formatting.Color
	Style     *RunStyle
	Size      *Sz
	SizeCs    *SzCs
	Shading   *formatting.Shading
	Highlight *formatting.Highlight
	Bold      *elemtypes.NullBoolElem
	Italic    *formatting.Italic

	Strike    *formatting.Strike
	Underline *formatting.Underline
	// Vanish           *Vanish
	// SpecVanish       *SpecVanish
	// VertAlign        *VertAlign
	// CharacterSpacing *CharacterSpacing
	// Fonts            *RunFonts
	// TextBorder       *TextBorder
	// Del              *Delete
	// Ins              *Insert
}

// NewRunProperty creates a new RunProperty with default values.
func NewRunProperty() RunProperty {
	return RunProperty{}
}

// MarshalXML marshals RunProperty to XML.
func (rp *RunProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:rPr"
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if rp.Color != nil {
		err = e.EncodeElement(rp.Color, xml.StartElement{Name: xml.Name{Local: "w:color"}})
		if err != nil {
			return err
		}
	}

	if rp.Style != nil {
		err = e.EncodeElement(rp.Style, xml.StartElement{Name: xml.Name{Local: "w:rStyle"}})
		if err != nil {
			return err
		}
	}

	if rp.Size != nil {
		err = e.EncodeElement(rp.Size, xml.StartElement{Name: xml.Name{Local: "w:sz"}})
		if err != nil {
			return err
		}
	}

	if rp.SizeCs != nil {
		err = e.EncodeElement(rp.SizeCs, xml.StartElement{Name: xml.Name{Local: "w:szCs"}})
		if err != nil {
			return err
		}
	}

	if rp.Shading != nil {
		err = e.EncodeElement(rp.Shading, xml.StartElement{Name: xml.Name{Local: "w:shd"}})
		if err != nil {
			return err
		}
	}

	if rp.Highlight != nil {
		err = e.EncodeElement(rp.Highlight, xml.StartElement{Name: xml.Name{Local: "w:highlight"}})
		if err != nil {
			return err
		}
	}

	if rp.Bold != nil {
		// err = e.EncodeElement(rp.Bold, xml.StartElement{Name: xml.Name{Local: "w:b"}})
		err = rp.Bold.MarshalXML(e, xml.StartElement{})
		if err != nil {
			return err
		}
	}

	if rp.Italic != nil {
		err = e.EncodeElement(rp.Italic, xml.StartElement{Name: xml.Name{Local: "w:i"}})
		if err != nil {
			return err
		}
	}

	if rp.Strike != nil {
		err = e.EncodeElement(rp.Strike, xml.StartElement{Name: xml.Name{Local: "w:strike"}})
		if err != nil {
			return err
		}
	}

	if rp.Underline != nil {
		err = e.EncodeElement(rp.Underline, xml.StartElement{Name: xml.Name{Local: "w:u"}})
		if err != nil {
			return err
		}
	}

	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}

	return nil
}

// UnmarshalXML unmarshals XML to RunProperty.
func (rp *RunProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name {
			case xml.Name{Space: constants.WMLNamespace, Local: "color"}, xml.Name{Space: constants.AltWMLNamespace, Local: "color"}:
				if err = d.DecodeElement(&rp.Color, &t); err != nil {
					return err
				}
			case xml.Name{Space: constants.WMLNamespace, Local: "rStyle"}, xml.Name{Space: constants.AltWMLNamespace, Local: "rStyle"}:
				if err = d.DecodeElement(&rp.Style, &t); err != nil {
					return err
				}
			case xml.Name{Space: constants.WMLNamespace, Local: "sz"}, xml.Name{Space: constants.AltWMLNamespace, Local: "sz"}:
				if err = d.DecodeElement(&rp.Size, &t); err != nil {
					return err
				}
			case xml.Name{Space: constants.WMLNamespace, Local: "szCs"}, xml.Name{Space: constants.AltWMLNamespace, Local: "szCs"}:
				if err = d.DecodeElement(&rp.SizeCs, &t); err != nil {
					return err
				}
			case xml.Name{Space: constants.WMLNamespace, Local: "shd"}, xml.Name{Space: constants.AltWMLNamespace, Local: "shd"}:
				if err = d.DecodeElement(&rp.Shading, &t); err != nil {
					return err
				}
			case xml.Name{Space: constants.WMLNamespace, Local: "highlight"}, xml.Name{Space: constants.AltWMLNamespace, Local: "highlight"}:
				if err = d.DecodeElement(&rp.Highlight, &t); err != nil {
					return err
				}
			case xml.Name{Space: constants.WMLNamespace, Local: "b"}, xml.Name{Space: constants.AltWMLNamespace, Local: "b"}:
				rp.Bold = formatting.DefaultBold()
				if err = rp.Bold.UnmarshalXML(d, t); err != nil {
					return err
				}
				// if err = d.DecodeElement(rp.Bold, &t); err != nil {
				// 	return err
				// }
			case xml.Name{Space: constants.WMLNamespace, Local: "i"}, xml.Name{Space: constants.AltWMLNamespace, Local: "i"}:
				if err = d.DecodeElement(&rp.Italic, &t); err != nil {
					return err
				}
			case xml.Name{Space: constants.WMLNamespace, Local: "strike"}, xml.Name{Space: constants.AltWMLNamespace, Local: "strike"}:
				if err = d.DecodeElement(&rp.Strike, &t); err != nil {
					return err
				}
			case xml.Name{Space: constants.WMLNamespace, Local: "u"}, xml.Name{Space: constants.AltWMLNamespace, Local: "u"}:
				if err = d.DecodeElement(&rp.Underline, &t); err != nil {
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
