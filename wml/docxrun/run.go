package docxrun

import (
	"encoding/xml"

	"github.com/gomutex/godocx/dml"
	"github.com/gomutex/godocx/elemtypes"
	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/formatting"
	"github.com/gomutex/godocx/wml/runcontent"
	"github.com/gomutex/godocx/wml/stypes"
)

// A Run is part of a paragraph that has its own style. It could be
type Run struct {
	//1. Run Properties
	RunProperty *RunProperty

	// 2. Choice - Run Inner content
	Children []*RunChild
}

type RunChild struct {
	//TODO: Maintain sequenec and rest of the elements
	InstrText *string
	Text      *runcontent.Text
	Drawing   *dml.Drawing
	Tab       *ctypes.Empty `xml:"tab,omitempty"`
	Break     *runcontent.Break
}

func NewRun() *Run {
	return &Run{}
}

// Sets the color of the Run.
//
// Example:
//
//	run := NewRun()
//	modifiedRun := run.Color("FF0000")
//
// Parameters:
//   - colorCode: A string representing the color code (e.g., "FF0000" for red).
//
// Returns:
//   - *Run: The modified Run instance with the updated color.
func (r *Run) Color(colorCode string) *Run {
	r.RunProperty.Color = formatting.NewColor(colorCode)
	return r
}

// Sets the size of the Run.

// This method takes an integer parameter representing the desired font size.
// It updates the size property of the Run instance with the specified size,
// Example:

// 	run := NewRun()
// 	modifiedRun := run.Size(12)

// Parameters:
//   - size: An integer representing the font size.

// Returns:
//   - *Run: The modified Run instance with the updated size.
func (r *Run) Size(size uint64) *Run {
	r.RunProperty.Size = NewFontSize(size * 2)
	return r
}

func (r *Run) Shading(shdType stypes.Shading, color, fill string) *Run {
	r.RunProperty.Shading = ctypes.NewShading().SetShadingType(shdType).SetColor(color).SetFill(fill)
	return r
}

// AddHighlight sets the highlight color for the run.
func (r *Run) Highlight(color string) *Run {
	r.RunProperty.Highlight = formatting.NewHighlight(color)
	return r
}

// AddBold enables bold formatting for the run.
func (r *Run) Bold(value bool) *Run {
	r.RunProperty.Bold = elemtypes.NewOptBoolElem(value)
	return r
}

func (r *Run) Italic(value bool) *Run {
	r.RunProperty.Italic = elemtypes.NewOptBoolElem(value)
	return r
}

// Specifies that the contents of this run shall be displayed with a single horizontal line through the center of the line.
func (r *Run) Strike(value bool) *Run {
	r.RunProperty.Strike = elemtypes.NewOptBoolElem(value)
	return r
}

// Specifies that the contents of this run shall be displayed with two horizontal lines through each character displayed on the line
func (r *Run) DoubleStrike(value bool) *Run {
	r.RunProperty.DoubleStrike = elemtypes.NewOptBoolElem(value)
	return r
}

// Display All Characters As Capital Letters
//
// Any lowercase characters in this text run shall be formatted for display only as their capital letter character equivalents
func (r *Run) Caps(value bool) *Run {
	r.RunProperty.Caps = elemtypes.NewOptBoolElem(value)
	return r
}

// Specifies that all small letter characters in this text run shall be formatted for display only as their capital letter character equivalents
func (r *Run) SmallCaps(value bool) *Run {
	r.RunProperty.Caps = elemtypes.NewOptBoolElem(value)
	return r
}

func (r *Run) Outline(value bool) *Run {
	r.RunProperty.Outline = elemtypes.NewOptBoolElem(value)
	return r
}

func (r *Run) Shadow(value bool) *Run {
	r.RunProperty.Shadow = elemtypes.NewOptBoolElem(value)
	return r
}

func (r *Run) Emboss(value bool) *Run {
	r.RunProperty.Emboss = elemtypes.NewOptBoolElem(value)
	return r
}

func (r *Run) Imprint(value bool) *Run {
	r.RunProperty.Imprint = elemtypes.NewOptBoolElem(value)
	return r
}

// Do Not Check Spelling or Grammar
func (r *Run) NoGrammer(value bool) *Run {
	r.RunProperty.NoGrammar = elemtypes.NewOptBoolElem(value)
	return r
}

// Use Document Grid Settings For Inter-Character Spacing
func (r *Run) SnapToGrid(value bool) *Run {
	r.RunProperty.SnapToGrid = elemtypes.NewOptBoolElem(value)
	return r
}

// Hidden Text
func (r *Run) HideText(value bool) *Run {
	r.RunProperty.Vanish = elemtypes.NewOptBoolElem(value)
	return r
}

func (r *Run) Spacing(value int) *Run {
	r.RunProperty.Spacing = ctypes.NewDecimalNum(value)
	return r
}

func (r *Run) Underline(value formatting.UnderlineStyle) *Run {
	r.RunProperty.Underline = formatting.NewUnderline(value)
	return r
}

func (r *Run) Style(value string) *Run {
	r.RunProperty.Style = NewRunStyle(value)
	return r
}

func (r *Run) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:r"

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	if r.RunProperty != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:rPr"}}
		if err = e.EncodeElement(r.RunProperty, propsElement); err != nil {
			return err
		}
	}

	for _, data := range r.Children {
		if data.Text != nil {
			err = data.Text.MarshalXML(e, xml.StartElement{})
			if err != nil {
				return err
			}
		}

		if data.Drawing != nil {
			err := data.Drawing.MarshalXML(e, xml.StartElement{})
			if err != nil {
				return err
			}
		}

		if data.Tab != nil {
			err := data.Tab.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:tab"}})
			if err != nil {
				return err
			}
		}

		if data.Break != nil {
			err := data.Break.MarshalXML(e, xml.StartElement{})
			if err != nil {
				return err
			}
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (r *Run) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
loop:
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "t":
				txt := runcontent.NewText()
				if err = d.DecodeElement(txt, &elem); err != nil {
					return err
				}

				r.Children = append(r.Children, &RunChild{Text: txt})
			case "rPr":
				r.RunProperty = &RunProperty{}
				if err = d.DecodeElement(r.RunProperty, &elem); err != nil {
					return err
				}
			case "tab":
				tabElem := &ctypes.Empty{}
				if err = d.DecodeElement(tabElem, &elem); err != nil {
					return err
				}

				r.Children = append(r.Children, &RunChild{
					Tab: tabElem,
				})
			case "br":
				br := &runcontent.Break{}
				if err = br.UnmarshalXML(d, elem); err != nil {
					return err
				}

				r.Children = append(r.Children, &RunChild{
					Break: br,
				})
			case "drawing":
				drawingElem := &dml.Drawing{}
				if err = d.DecodeElement(drawingElem, &elem); err != nil {
					return err
				}

				r.Children = append(r.Children, &RunChild{
					Drawing: drawingElem,
				})
			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break loop
		}
	}

	return nil
}
