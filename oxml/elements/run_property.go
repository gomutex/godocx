package elements

import (
	"encoding/xml"

	"github.com/gomutex/godocx/constants"
)

// RunProperty represents the properties of a run of text within a paragraph.
type RunProperty struct {
	Color   *Color
	Style   *RunStyle
	Size    *Sz
	SizeCs  *SzCs
	Shading *Shading

	// Bold             *Bold
	// BoldCs           *BoldCs
	// Italic           *Italic
	// ItalicCs         *ItalicCs
	// Strike           *Strike
	// Highlight        *Highlight
	// Underline        *Underline
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

// AddColor sets the font color for the run.
func (rp *RunProperty) AddColor(color string) *RunProperty {
	rp.Color = NewColor(color)
	return rp
}

func (rp *RunProperty) AddRunStyle(style string) *RunProperty {
	rp.Style = NewRunStyle(style)
	return rp
}

// AddSize sets the font size for the run.
func (rp *RunProperty) AddSize(size int) *RunProperty {
	rp.Size = NewSz(size)
	rp.SizeCs = NewSzCs(size)
	return rp
}

func (rp *RunProperty) AddShading(shdType ShadingType, color, fill string) *RunProperty {
	rp.Shading = NewShading().SetShadingType(shdType).SetColor(color).SetFill(fill)
	return rp
}

// // AddSpacing sets the character spacing for the run.
// func (rp *RunProperty) AddSpacing(spacing int) *RunProperty {
// 	rp.CharacterSpacing = &CharacterSpacing{Val: spacing}
// 	return rp
// }

// // AddHighlight sets the highlight color for the run.
// func (rp *RunProperty) AddHighlight(color string) *RunProperty {
// 	rp.Highlight = &Highlight{Val: color}
// 	return rp
// }

// // AddVertAlign sets the vertical alignment for the run.
// func (rp *RunProperty) AddVertAlign(alignType VertAlignType) *RunProperty {
// 	rp.VertAlign = &VertAlign{Val: alignType}
// 	return rp
// }

// // AddBold enables bold formatting for the run.
// func (rp *RunProperty) AddBold() *RunProperty {
// 	rp.Bold = &Bold{}
// 	rp.BoldCs = &BoldCs{}
// 	return rp
// }

// // AddDisableBold disables bold formatting for the run.
// func (rp *RunProperty) AddDisableBold() *RunProperty {
// 	rp.Bold = &Bold{Val: false}
// 	rp.BoldCs = &BoldCs{Val: false}
// 	return rp
// }

// // AddCaps enables capitalization formatting for the run.
// func (rp *RunProperty) AddCaps() *RunProperty {
// 	rp.Caps = &Caps{}
// 	return rp
// }

// // AddItalic enables italic formatting for the run.
// func (rp *RunProperty) AddItalic() *RunProperty {
// 	rp.Italic = &Italic{}
// 	rp.ItalicCs = &ItalicCs{}
// 	return rp
// }

// // AddStrike enables strike-through formatting for the run.
// func (rp *RunProperty) AddStrike() *RunProperty {
// 	rp.Strike = &Strike{}
// 	return rp
// }

// // AddUnderline sets the underline type for the run.
// func (rp *RunProperty) AddUnderline(lineType string) *RunProperty {
// 	rp.Underline = &Underline{Val: lineType}
// 	return rp
// }

// // AddVanish enables text visibility to be toggled in the run.
// func (rp *RunProperty) AddVanish() *RunProperty {
// 	rp.Vanish = &Vanish{}
// 	return rp
// }

// // AddSpecVanish enables text visibility to be toggled in the run.
// func (rp *RunProperty) AddSpecVanish() *RunProperty {
// 	rp.SpecVanish = &SpecVanish{}
// 	return rp
// }

// // AddFonts sets the run fonts.
// func (rp *RunProperty) AddFonts(fonts RunFonts) *RunProperty {
// 	rp.Fonts = &fonts
// 	return rp
// }

// // AddTextBorder sets the text border for the run.
// func (rp *RunProperty) AddTextBorder(border TextBorder) *RunProperty {
// 	rp.TextBorder = &border
// 	return rp
// }

// // AddDelete sets the deletion formatting for the run.
// func (rp *RunProperty) AddDelete(d Delete) *RunProperty {
// 	rp.Del = &d
// 	return rp
// }

// // AddInsert sets the insertion formatting for the run.
// func (rp *RunProperty) AddInsert(i Insert) *RunProperty {
// 	rp.Ins = &i
// 	return rp
// }

// MarshalXML marshals RunProperty to XML.
func (rp *RunProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:rPr"
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if rp.Color != nil {
		err := e.EncodeElement(rp.Color, xml.StartElement{Name: xml.Name{Local: "w:color"}})
		if err != nil {
			return err
		}
	}

	if rp.Style != nil {
		err := e.EncodeElement(rp.Style, xml.StartElement{Name: xml.Name{Local: "w:rStyle"}})
		if err != nil {
			return err
		}
	}

	if rp.Size != nil {
		err := e.EncodeElement(rp.Size, xml.StartElement{Name: xml.Name{Local: "w:sz"}})
		if err != nil {
			return err
		}
	}

	if rp.SizeCs != nil {
		err := e.EncodeElement(rp.SizeCs, xml.StartElement{Name: xml.Name{Local: "w:szCs"}})
		if err != nil {
			return err
		}
	}

	if rp.Shading != nil {
		err := e.EncodeElement(rp.Shading, xml.StartElement{Name: xml.Name{Local: "w:shd"}})
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
