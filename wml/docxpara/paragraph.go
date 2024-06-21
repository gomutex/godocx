package docxpara

import (
	"encoding/xml"

	"github.com/gomutex/godocx/common/units"
	"github.com/gomutex/godocx/dml"
	"github.com/gomutex/godocx/dml/dmlct"
	"github.com/gomutex/godocx/dml/dmlpic"
	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/docxrun"
	"github.com/gomutex/godocx/wml/runcontent"
)

type ParagraphChild struct {
	Link *Hyperlink   // w:hyperlink
	Run  *docxrun.Run // i.e w:r
}

type Paragraph struct {
	id       string
	Property *ParagraphProp

	Children []*ParagraphChild
}

func NewParagraph() *Paragraph {
	return &Paragraph{
		Property: DefaultParaProperty(),
	}
}

func DefaultParagraph() *Paragraph {
	return &Paragraph{
		Property: DefaultParaProperty(),
	}
}

func NewParagraphChild() *ParagraphChild {
	return &ParagraphChild{}
}

func DefaultParagraphChild() *ParagraphChild {
	return &ParagraphChild{}
}

// Style sets the paragraph style.
//
// Parameters:
//   - value: A string representing the style value. It can be any valid style defined in the WordprocessingML specification.
//
// Returns:
//   - *Paragraph: A pointer to the modified Paragraph instance with the updated style.
//
// Example:
//
//	p1 := docx.AddParagraph("Example para")
//	paragraph.Style("List Number")
func (p *Paragraph) Style(value string) *Paragraph {
	if p.Property == nil {
		p.Property = DefaultParaProperty()
	}
	p.Property.Style = NewParagraphStyle(value)
	return p
}

// Justification sets the paragraph justification type.
//
// Parameters:
//   - value: A string representing the justification value. It can be one of the following:
//     - "left" for left justification
//     - "center" for center justification
//     - "right" for right justification
//     - "both" for justification with equal spacing on both sides
//     - "distribute": Paragraph characters are distributed to fill the entire width of paragraph
//
// Returns:
//   - *Paragraph: A pointer to the modified Paragraph instance with the updated justification.

func (p *Paragraph) Justification(value string) (*Paragraph, error) {
	if p.Property == nil {
		p.Property = DefaultParaProperty()
	}

	jc, err := ctypes.NewJustification(value)
	if err != nil {
		return nil, err
	}

	p.Property.Justification = jc
	return p, nil
}

func (p *Paragraph) Numbering(id int, level int) {

	if p.Property == nil {
		p.Property = DefaultParaProperty()
	}

	if p.Property.NumProp == nil {
		p.Property.NumProp = &NumProp{}
	}
	p.Property.NumProp.NumID = ctypes.NewDecimalNum(id)
	p.Property.NumProp.ILvl = ctypes.NewDecimalNum(level)
}

// Appends a new text to the Paragraph.
// Example:
//
//	paragraph := AddParagraph()
//	modifiedParagraph := paragraph.AddText("Hello, World!")
//
// Parameters:
//   - text: A string representing the text to be added to the Paragraph.
//
// Returns:
//   - *Run: The newly created Run instance added to the Paragraph.
func (p *Paragraph) AddText(text string) *docxrun.Run {
	t := runcontent.TextFromString(text)

	runChildren := []*docxrun.RunChild{}
	runChildren = append(runChildren, &docxrun.RunChild{
		Text: t,
	})
	run := &docxrun.Run{
		Children:    runChildren,
		RunProperty: &docxrun.RunProperty{},
	}

	p.Children = append(p.Children, &ParagraphChild{Run: run})

	return run
}

// func (p *Paragraph) AddLink(text string, link string) *Hyperlink {
// 	rId := p.rootRef.addLinkRelation(link)

// 	runChildren := []*RunChild{}
// 	runChildren = append(runChildren, &RunChild{
// 		InstrText: &text,
// 	})
// 	run := &Run{
// 		Children: runChildren,
// 		RunProperty: &RunProperty{
// 			RunStyle: &RunStyle{
// 				Val: constants.HyperLinkStyle,
// 			},
// 		},
// 	}

// 	paraChild := &ParagraphChild{
// 		Run: run,
// 	}

// 	hyperLink := &Hyperlink{
// 		ID: rId,
// 	}
// 	hyperLink.Children = append(hyperLink.Children, paraChild)

// 	p.Children = append(p.Children, &ParagraphChild{Link: hyperLink})

// 	return hyperLink
// }

func (p *Paragraph) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:p"

	// Opening <w:p> element
	if err = e.EncodeToken(start); err != nil {
		return err
	}

	if p.Property != nil {
		if err = e.EncodeElement(p.Property, start); err != nil {
			return err
		}
	}

	for _, cElem := range p.Children {
		if cElem.Run != nil {
			if err = e.EncodeElement(cElem.Run, start); err != nil {
				return err
			}
		}

		if cElem.Link != nil {
			if err = e.EncodeElement(cElem.Link, start); err != nil {
				return err
			}
		}
	}

	// Closing </w:p> element
	if err = e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}

	return nil
}

func (p *Paragraph) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
loop:
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "r":
				r := docxrun.NewRun()
				if err = d.DecodeElement(r, &elem); err != nil {
					return err
				}

				p.Children = append(p.Children, &ParagraphChild{Run: r})
			case "pPr":
				p.Property = &ParagraphProp{}
				if err = d.DecodeElement(p.Property, &elem); err != nil {
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

func AddParagraph(text string) *Paragraph {
	p := &Paragraph{
		Children: []*ParagraphChild{},
	}
	p.AddText(text)

	return p
}

func (p *Paragraph) AddDrawing(rID string, imgCount uint, width units.Inch, height units.Inch) *dml.Inline {
	eWidth := width.ToEmu()
	eHeight := height.ToEmu()

	inline := dml.Inline{
		Extent:  *dmlct.NewPostvSz2D(eWidth, eHeight),
		Graphic: *dml.NewPicGraphic(dmlpic.NewPic(rID, imgCount, eWidth, eHeight)),
	}

	runChildren := []*docxrun.RunChild{}
	drawing := &dml.Drawing{}

	drawing.Inline = append(drawing.Inline, &inline)

	runChildren = append(runChildren, &docxrun.RunChild{
		Drawing: drawing,
	})

	run := &docxrun.Run{
		Children:    runChildren,
		RunProperty: &docxrun.RunProperty{},
	}

	p.Children = append(p.Children, &ParagraphChild{Run: run})

	return &inline
}

type Hyperlink struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main hyperlink,omitempty"`
	ID      string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr"`
	// Run     Run
	Children []*ParagraphChild
}
