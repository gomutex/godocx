package oxml

import (
	"encoding/xml"

	"github.com/gomutex/godocx/constants"
	"github.com/gomutex/godocx/oxml/elements"
)

// A Run is part of a paragraph that has its own style. It could be
type Run struct {
	RunProperties *RunProperties
	Children      []*RunChild
}

type RunChild struct {
	InstrText *string
	Text      *elements.Text
}

type Hyperlink struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main hyperlink,omitempty"`
	ID      string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr"`
	// Run     Run
	Children []*ParagraphChild
}

type RunProperties struct {
	XMLName  xml.Name  `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rPr,omitempty"`
	CTColor  *CTColor  `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main color,omitempty"`
	Size     *CTSize   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main sz,omitempty"`
	RunStyle *RunStyle `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rStyle,omitempty"`
	Style    *CTStyle  `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pStyle,omitempty"`
}

type RunStyle struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rStyle,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

type CTStyle struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main pStyle,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

type CTColor struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main color"`
	Val     string   `xml:"w:val,attr"`
}

type CTSize struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main sz"`
	Val     int      `xml:"w:val,attr"`
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

	r.RunProperties.CTColor = &CTColor{
		Val: colorCode,
	}

	return r
}

// Sets the size of the Run.
//
// This method takes an integer parameter representing the desired font size.
// It updates the size property of the Run instance with the specified size,
// Example:
//
//	run := NewRun()
//	modifiedRun := run.Size(12)
//
// Parameters:
//   - size: An integer representing the font size.
//
// Returns:
//   - *Run: The modified Run instance with the updated size.
func (r *Run) Size(size int) *Run {
	r.RunProperties.Size = &CTSize{
		Val: size * 2,
	}
	return r
}

func (r *Run) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:r"

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	if r.RunProperties != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:rPr"}}
		if err = e.EncodeElement(r.RunProperties, propsElement); err != nil {
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
				txt := elements.NewText()
				if err := d.DecodeElement(txt, &elem); err != nil {
					return err
				}

				r.Children = append(r.Children, &RunChild{Text: txt})
			case xml.Name{Space: constants.WMLNamespace, Local: "rPr"}, xml.Name{Space: constants.AltWMLNamespace, Local: "rPr"}:
				r.RunProperties = &RunProperties{}
				if err := d.DecodeElement(r.RunProperties, &elem); err != nil {
					return err
				}

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
