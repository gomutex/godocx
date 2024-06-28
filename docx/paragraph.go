package docx

import (
	"errors"
	"fmt"

	"github.com/gomutex/godocx/common/units"
	"github.com/gomutex/godocx/dml"
	"github.com/gomutex/godocx/dml/dmlct"
	"github.com/gomutex/godocx/dml/dmlpic"
	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/runcontent"
	"github.com/gomutex/godocx/wml/stypes"
)

type Paragraph struct {
	// Reverse inheriting the Rootdoc into Paragrah to access other elements
	Root *RootDoc

	// Paragraph Complex Type
	CT ctypes.Paragraph
}

func NewParagraph(root *RootDoc) *Paragraph {
	return &Paragraph{
		Root: root,
	}
}

func NewParagraphChild() *ctypes.ParagraphChild {
	return &ctypes.ParagraphChild{}
}

func DefaultParagraphChild() *ctypes.ParagraphChild {
	return &ctypes.ParagraphChild{}
}

// AddParagraph adds a new paragraph with the specified text to the document.
// It returns the created Paragraph instance.
//
// Parameters:
//   - text: The text to be added to the paragraph.
//
// Returns:
//   - p: The created Paragraph instance.
func (rd *RootDoc) AddParagraph(text string) *Paragraph {
	p := Paragraph{}
	p.AddText(text)
	bodyElem := DocumentChild{
		Para: &p,
	}
	rd.Document.Body.Children = append(rd.Document.Body.Children, bodyElem)

	return &p
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
func (p *Paragraph) Style(value string) {
	if p.CT.Property == nil {
		p.CT.Property = ctypes.DefaultParaProperty()
	}
	p.CT.Property.Style = ctypes.NewParagraphStyle(value)
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

func (p *Paragraph) Justification(value string) error {
	if p.CT.Property == nil {
		p.CT.Property = ctypes.DefaultParaProperty()
	}

	val, err := stypes.JustificationFromStr(value)
	if err != nil {
		return err
	}

	p.CT.Property.Justification = ctypes.NewGenSingleStrVal(val)

	return nil
}

func (p Paragraph) Numbering(id int, level int) {

	if p.CT.Property == nil {
		p.CT.Property = ctypes.DefaultParaProperty()
	}

	if p.CT.Property.NumProp == nil {
		p.CT.Property.NumProp = &ctypes.NumProp{}
	}
	p.CT.Property.NumProp.NumID = ctypes.NewDecimalNum(id)
	p.CT.Property.NumProp.ILvl = ctypes.NewDecimalNum(level)
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
func (p *Paragraph) AddText(text string) *ctypes.Run {
	t := runcontent.TextFromString(text)

	runChildren := []*ctypes.RunChild{}
	runChildren = append(runChildren, &ctypes.RunChild{
		Text: t,
	})
	run := &ctypes.Run{
		Children: runChildren,
	}

	p.CT.Children = append(p.CT.Children, ctypes.ParagraphChild{Run: run})

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

func (p *Paragraph) AddDrawing(rID string, imgCount uint, width units.Inch, height units.Inch) *dml.Inline {
	eWidth := width.ToEmu()
	eHeight := height.ToEmu()

	inline := dml.NewInline(
		*dmlct.NewPostvSz2D(eWidth, eHeight),
		dml.DocProp{
			ID:   uint64(imgCount),
			Name: fmt.Sprintf("Image%d", imgCount),
		},
		*dml.NewPicGraphic(dmlpic.NewPic(rID, imgCount, eWidth, eHeight)),
	)

	runChildren := []*ctypes.RunChild{}
	drawing := &dml.Drawing{}

	drawing.Inline = append(drawing.Inline, inline)

	runChildren = append(runChildren, &ctypes.RunChild{
		Drawing: drawing,
	})

	run := &ctypes.Run{
		Children: runChildren,
	}

	p.CT.Children = append(p.CT.Children, ctypes.ParagraphChild{Run: run})

	return &inline
}

func (p *Paragraph) GetStyle(styleID string) (*ctypes.Style, error) {
	if p.CT.Property == nil || p.CT.Property.Style == nil {
		return nil, errors.New("No property for the style")
	}

	style := p.Root.GetStyleByID(p.CT.Property.Style.Val, stypes.StyleTypeParagraph)
	if style == nil {
		return nil, errors.New("No style found for the paragraph")
	}

	return style, nil
}
