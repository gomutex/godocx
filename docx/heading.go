package docx

import (
	"errors"
	"fmt"

	"github.com/gomutex/godocx/wml/ctypes"
)

// AddHeading adds a new heading paragraph with the specified text and level
// to the end of the document.
//
// Parameters:
//   - text: The text to be added to the heading.
//   - level: The heading level (0-9). If level is 0, the style is
//     set to Title. Otherwise, the style is set to Heading {level}.
//     If level is outside the range 0-9, an error will be returned.
//
// Returns:
//   - *Paragraph: The newly created Paragraph instance.
//   - error: An error if the level is outside the supported range.
func (rd *RootDoc) AddHeading(text string, level uint) (*Paragraph, error) {
	p, err := rd.AddEmptyHeading(level)
	if err != nil {
		return nil, err
	}

	p.AddText(text)
	return p, nil
}

// AddEmptyHeading adds an empty heading paragraph to the end of the document.
//
// Parameters:
//   - level: The heading level (0-9). If level is 0, the style is
//     set to Title. Otherwise, the style is set to Heading {level}.
//     If level is outside the range 0-9, an error will be returned.
//
// Returns:
//   - *Paragraph: The newly created Paragraph instance.
//   - error: An error if the level is outside the supported range.
func (rd *RootDoc) AddEmptyHeading(level uint) (*Paragraph, error) {
	if level < 0 || level > 9 {
		return nil, errors.New("heading level not supported")
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

	return p, nil
}
