package elements

import (
	"encoding/xml"

	"github.com/gomutex/godocx/constants"
)

// A Run is part of a paragraph that has its own style. It could be
type Run struct {
	RunProperty *RunProperty
	Children    []*RunChild
}

type RunChild struct {
	InstrText *string
	Text      *Text
	Drawing   *Drawing
}

type Hyperlink struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main hyperlink,omitempty"`
	ID      string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr"`
	// Run     Run
	Children []*ParagraphChild
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

	r.RunProperty.Color = NewColor(colorCode)

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
func (r *Run) Size(size uint) *Run {
	r.RunProperty.Size = NewSz(size * 2)
	return r
}

func (r *Run) Shading(shdType ShadingType, color, fill string) *Run {
	r.RunProperty.Shading = NewShading().SetShadingType(shdType).SetColor(color).SetFill(fill)
	return r
}

// AddHighlight sets the highlight color for the run.
func (r *Run) Highlight(color ColorIndex) *Run {
	r.RunProperty.Highlight = NewHighlight(color)
	return r
}

// AddBold enables bold formatting for the run.
func (r *Run) Bold(value bool) *Run {
	r.RunProperty.Bold = NewBold(value)
	return r
}

func (r *Run) Italic(value bool) *Run {
	r.RunProperty.Italic = NewItalic(value)
	return r
}

func (r *Run) Strike(value bool) *Run {
	r.RunProperty.Strike = NewStrike(value)
	return r
}

func (r *Run) Underline(value UnderlineStyle) *Run {
	r.RunProperty.Underline = NewUnderline(value)
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
			err := data.Text.MarshalXML(e, start)
			if err != nil {
				return err
			}
		}

		if data.InstrText != nil {
			cElem := xml.StartElement{Name: xml.Name{Local: "w:instrText"}}
			if err = e.EncodeElement(data.InstrText, cElem); err != nil {
				return err
			}
		}

		if data.Drawing != nil {
			err := data.Drawing.MarshalXML(e, start)
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
			switch elem.Name {
			case xml.Name{Space: constants.WMLNamespace, Local: "t"}, xml.Name{Space: constants.AltWMLNamespace, Local: "t"}:
				txt := NewText()
				if err := d.DecodeElement(txt, &elem); err != nil {
					return err
				}

				r.Children = append(r.Children, &RunChild{Text: txt})
			case xml.Name{Space: constants.WMLNamespace, Local: "rPr"}, xml.Name{Space: constants.AltWMLNamespace, Local: "rPr"}:
				r.RunProperty = &RunProperty{}
				if err := d.DecodeElement(r.RunProperty, &elem); err != nil {
					return err
				}
			case xml.Name{Space: constants.WMLNamespace, Local: "drawing"}, xml.Name{Space: constants.AltWMLNamespace, Local: "drawing"}:
				drawingElem := &Drawing{}
				if err := d.DecodeElement(drawingElem, &elem); err != nil {
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
