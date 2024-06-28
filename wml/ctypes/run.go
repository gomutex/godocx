package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/dml"
	"github.com/gomutex/godocx/internal"
	"github.com/gomutex/godocx/wml/runcontent"
	"github.com/gomutex/godocx/wml/stypes"
)

// A Run is part of a paragraph that has its own style. It could be
type Run struct {
	// Attributes
	RsidRPr *stypes.LongHexNum // Revision Identifier for Run Properties
	RsidR   *stypes.LongHexNum // Revision Identifier for Run
	RsidDel *stypes.LongHexNum // Revision Identifier for Run Deletion

	//1. Run Properties
	Property *RunProperty

	// 2. Choice - Run Inner content
	Children []*RunChild
}

type RunChild struct {
	//TODO: Maintain sequence and rest of the elements
	//Field Code
	InstrText *string
	Text      *runcontent.Text
	Drawing   *dml.Drawing
	Tab       *Empty `xml:"tab,omitempty"`
	Break     *runcontent.Break
}

func NewRun() *Run {
	return &Run{}
}

// Get the Run property. If it is not initialized, create one and return it.
func (r *Run) getProp() *RunProperty {
	if r.Property == nil {
		r.Property = &RunProperty{}
	}
	return r.Property
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
	r.getProp().Color = NewColor(colorCode)
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
	r.getProp().Size = NewFontSize(size * 2)
	return r
}

func (r *Run) Shading(shdType stypes.Shading, color, fill string) *Run {
	r.getProp().Shading = NewShading().SetShadingType(shdType).SetColor(color).SetFill(fill)
	return r
}

// AddHighlight sets the highlight color for the run.
func (r *Run) Highlight(color string) *Run {
	r.getProp().Highlight = NewCTString(color)
	return r
}

// AddBold enables bold formatting for the run.
func (r *Run) Bold(value bool) *Run {
	r.getProp().Bold = OnOffFromBool(value)
	return r
}

func (r *Run) Italic(value bool) *Run {
	r.getProp().Italic = OnOffFromBool(value)
	return r
}

// Specifies that the contents of this run shall be displayed with a single horizontal line through the center of the line.
func (r *Run) Strike(value bool) *Run {
	r.getProp().Strike = OnOffFromBool(value)
	return r
}

// Specifies that the contents of this run shall be displayed with two horizontal lines through each character displayed on the line
func (r *Run) DoubleStrike(value bool) *Run {
	r.getProp().DoubleStrike = OnOffFromBool(value)
	return r
}

// Display All Characters As Capital Letters
//
// Any lowercase characters in this text run shall be formatted for display only as their capital letter character equivalents
func (r *Run) Caps(value bool) *Run {
	r.getProp().Caps = OnOffFromBool(value)
	return r
}

// Specifies that all small letter characters in this text run shall be formatted for display only as their capital letter character equivalents
func (r *Run) SmallCaps(value bool) *Run {
	r.getProp().Caps = OnOffFromBool(value)
	return r
}

func (r *Run) Outline(value bool) *Run {
	r.getProp().Outline = OnOffFromBool(value)
	return r
}

func (r *Run) Shadow(value bool) *Run {
	r.getProp().Shadow = OnOffFromBool(value)
	return r
}

func (r *Run) Emboss(value bool) *Run {
	r.getProp().Emboss = OnOffFromBool(value)
	return r
}

func (r *Run) Imprint(value bool) *Run {
	r.getProp().Imprint = OnOffFromBool(value)
	return r
}

// Do Not Check Spelling or Grammar
func (r *Run) NoGrammer(value bool) *Run {
	r.getProp().NoGrammar = OnOffFromBool(value)
	return r
}

// Use Document Grid Settings For Inter-Character Spacing
func (r *Run) SnapToGrid(value bool) *Run {
	r.getProp().SnapToGrid = OnOffFromBool(value)
	return r
}

// Hidden Text
func (r *Run) HideText(value bool) *Run {
	r.getProp().Vanish = OnOffFromBool(value)
	return r
}

func (r *Run) Spacing(value int) *Run {
	r.getProp().Spacing = NewDecimalNum(value)
	return r
}

func (r *Run) Underline(value stypes.Underline) *Run {
	r.getProp().Underline = NewGenSingleStrVal(value)
	return r
}

func (r *Run) Style(value string) *Run {
	r.getProp().Style = NewRunStyle(value)
	return r
}

func (r Run) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:r"

	if r.RsidRPr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidRPr"}, Value: string(*r.RsidRPr)})
	}
	if r.RsidR != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidR"}, Value: string(*r.RsidR)})
	}
	if r.RsidDel != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidDel"}, Value: string(*r.RsidDel)})
	}

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	if r.Property != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:rPr"}}
		if err = e.EncodeElement(r.Property, propsElement); err != nil {
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
	// Decode attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "rsidRPr":
			r.RsidRPr = internal.ToPtr(stypes.LongHexNum(attr.Value))
		case "rsidR":
			r.RsidR = internal.ToPtr(stypes.LongHexNum(attr.Value))
		case "rsidDel":
			r.RsidDel = internal.ToPtr(stypes.LongHexNum(attr.Value))
		}
	}

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
				r.Property = &RunProperty{}
				if err = d.DecodeElement(r.Property, &elem); err != nil {
					return err
				}
			case "tab":
				tabElem := &Empty{}
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
