package docx

import (
	"errors"
	"fmt"

	"github.com/gomutex/godocx/wml/ctypes"
)

// Return a heading paragraph newly added to the end of the document.
// The heading paragraph will contain text and have its paragraph style determined by level.
// If level is 0, the style is set to Title.
// The style is set to Heading {level}.
// if level is outside the range 0-9, error will be returned
func (rd *RootDoc) AddHeading(text string, level uint) (*Paragraph, error) {
	if level < 0 || level > 9 {
		return nil, errors.New("Heading level not supported")
	}

	p := newParagraph(rd)
	p.ct.Property = ctypes.DefaultParaProperty()

	style := "Title"
	if level != 0 {
		style = fmt.Sprintf("Heading%d", level)
	}

	p.ct.Property.Style = ctypes.NewParagraphStyle(style)

	bodyElem := DocumentChild{
		Para: p,
	}
	rd.Document.Body.Children = append(rd.Document.Body.Children, bodyElem)

	p.AddText(text)
	return p, nil
}
