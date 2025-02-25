package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/internal"
	"github.com/gomutex/godocx/wml/stypes"
)

type Paragraph struct {
	id string

	// Attributes
	RsidRPr      *stypes.LongHexNum // Revision Identifier for Paragraph Glyph Formatting
	RsidR        *stypes.LongHexNum // Revision Identifier for Paragraph
	RsidDel      *stypes.LongHexNum // Revision Identifier for Paragraph Deletion
	RsidP        *stypes.LongHexNum // Revision Identifier for Paragraph Properties
	RsidRDefault *stypes.LongHexNum // Default Revision Identifier for Runs

	// 1. Paragraph Properties
	Property *ParagraphProp

	// 2. Choices (Slice of Child elements)
	Children []ParagraphChild
}

type ParagraphChild struct {
	Link *Hyperlink // w:hyperlink
	Run  *Run       // i.e w:r
}

type Hyperlink struct {
	XMLName  xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main hyperlink,omitempty"`
	ID       string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr"`
	Run      *Run
	Children []ParagraphChild
}

func (p Paragraph) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:p"

	if p.RsidRPr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidRPr"}, Value: string(*p.RsidRPr)})
	}

	if p.RsidR != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidR"}, Value: string(*p.RsidR)})
	}
	if p.RsidDel != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidDel"}, Value: string(*p.RsidDel)})
	}
	if p.RsidP != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidP"}, Value: string(*p.RsidP)})
	}
	if p.RsidRDefault != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidRDefault"}, Value: string(*p.RsidRDefault)})
	}

	if err = e.EncodeToken(start); err != nil {
		return err
	}

	if p.Property != nil {
		if err = p.Property.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:r"},
		}); err != nil {
			return err
		}
	}

	for _, cElem := range p.Children {
		if cElem.Run != nil {
			if err = cElem.Run.MarshalXML(e, xml.StartElement{
				Name: xml.Name{Local: "w:r"},
			}); err != nil {
				return err
			}
		}

		if cElem.Link != nil {
			if err = e.EncodeElement(cElem.Link, xml.StartElement{
				Name: xml.Name{Local: "w:hyperlink"},
			}); err != nil {
				return err
			}
		}
	}

	// Closing </w:p> element
	return e.EncodeToken(start.End())
}

func (p *Paragraph) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	// Decode attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "rsidRPr":
			p.RsidRPr = internal.ToPtr(stypes.LongHexNum(attr.Value))
		case "rsidR":
			p.RsidR = internal.ToPtr(stypes.LongHexNum(attr.Value))
		case "rsidDel":
			p.RsidDel = internal.ToPtr(stypes.LongHexNum(attr.Value))
		case "rsidP":
			p.RsidP = internal.ToPtr(stypes.LongHexNum(attr.Value))
		case "rsidRDefault":
			p.RsidRDefault = internal.ToPtr(stypes.LongHexNum(attr.Value))
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
	t := TextFromString(text)

	runChildren := []RunChild{}
	runChildren = append(runChildren, RunChild{
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
