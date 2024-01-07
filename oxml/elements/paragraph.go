package elements

import (
	"encoding/xml"

	"github.com/gomutex/godocx/constants"
)

type ParagraphChild struct {
	Link *Hyperlink // w:hyperlink
	Run  *Run       // i.e w:r
}

type Paragraph struct {
	id       string
	property *ParagraphProperty

	Children []*ParagraphChild
}

func NewParagraph() *Paragraph {
	return &Paragraph{}
}

func NewParagraphChild() *ParagraphChild {
	return &ParagraphChild{}
}

func (p *Paragraph) GetType() string {
	return "p"
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
func (p *Paragraph) AddText(text string) *Run {
	t := TextFromString(text)

	runChildren := []*RunChild{}
	runChildren = append(runChildren, &RunChild{
		Text: t,
	})
	run := &Run{
		Children:    runChildren,
		RunProperty: &RunProperty{},
	}

	p.Children = append(p.Children, &ParagraphChild{Run: run})

	return run
}

// func (para *Paragraph) AddLink(text string, link string) *Hyperlink {
// 	rId := para.rootRef.addLinkRelation(link)

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

// 	para.Children = append(para.Children, &ParagraphChild{Link: hyperLink})

// 	return hyperLink
// }

func (para *Paragraph) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:p"

	// Opening <w:p> element
	if err = e.EncodeToken(start); err != nil {
		return err
	}

	if para.property != nil {
		if err = e.EncodeElement(para.property, start); err != nil {
			return err
		}
	}

	for _, cElem := range para.Children {
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
			switch elem.Name {
			case xml.Name{Space: constants.WMLNamespace, Local: "r"}, xml.Name{Space: constants.AltWMLNamespace, Local: "r"}:
				r := NewRun()
				if err := d.DecodeElement(r, &elem); err != nil {
					return err
				}

				p.Children = append(p.Children, &ParagraphChild{Run: r})
			case xml.Name{Space: constants.WMLNamespace, Local: "pPr"}, xml.Name{Space: constants.AltWMLNamespace, Local: "pPr"}:

				p.property = &ParagraphProperty{}
				if err := d.DecodeElement(p.property, &elem); err != nil {
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
