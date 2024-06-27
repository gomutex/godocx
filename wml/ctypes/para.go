package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/runcontent"
)

type Paragraph struct {
	id string

	// 1. Paragraph Properties
	Property *ParagraphProp `xml:"pPr,omitempty"`

	// 2. Choices (Slice of Child elements)
	Children []ParagraphChild
}

type ParagraphChild struct {
	Link *Hyperlink // w:hyperlink
	Run  *Run       // i.e w:r
}

type Hyperlink struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main hyperlink,omitempty"`
	ID      string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr"`
	// Run     Run
	Children []ParagraphChild
}

func (p Paragraph) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
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
				r := NewRun()
				if err = d.DecodeElement(r, &elem); err != nil {
					return err
				}

				p.Children = append(p.Children, ParagraphChild{Run: r})
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

func (p *Paragraph) AddText(text string) *Run {
	t := runcontent.TextFromString(text)

	runChildren := []*RunChild{}
	runChildren = append(runChildren, &RunChild{
		Text: t,
	})
	run := &Run{
		Children: runChildren,
	}

	p.Children = append(p.Children, ParagraphChild{Run: run})

	return run
}

func AddParagraph(text string) *Paragraph {
	p := Paragraph{}
	p.AddText(text)
	return &p
}
