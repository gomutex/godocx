package doc

import (
	"encoding/xml"
	"sync"

	"github.com/gomutex/godocx/wml/docxpara"
)

// RootDoc represents the root document of an Office Open XML (OOXML) document.
// It contains information about the document path, file map, the document structure,
// and relationships with other parts of the document.
type RootDoc struct {
	Path        string        // Path represents the path of the document.
	FileMap     sync.Map      // FileMap is a synchronized map for managing files related to the document.
	Document    *Document     // Document is the main document structure.
	RootRels    Relationships // RootRels represents relationships at the root level.
	ContentType ContentTypes

	rID        int // rId is used to generate unique relationship IDs.
	ImageCount uint
}

// NewRootDoc creates a new instance of the RootDoc structure.
func NewRootDoc() *RootDoc {
	return &RootDoc{}
}

// LoadDocXml decodes the provided XML data and returns a Document instance.
// It is used to load the main document structure from the document file.
//
// Parameters:
//   - fileName: The name of the document file.
//   - fileBytes: The XML data representing the main document structure.
//
// Returns:
//   - doc: The Document instance containing the decoded main document structure.
//   - err: An error, if any occurred during the decoding process.
func LoadDocXml(fileName string, fileBytes []byte) (*Document, error) {
	doc := Document{}
	err := xml.Unmarshal(fileBytes, &doc)
	if err != nil {
		return nil, err
	}

	doc.relativePath = fileName
	return &doc, nil
}

// AddParagraph adds a new paragraph with the specified text to the document.
// It returns the created Paragraph instance.
//
// Parameters:
//   - text: The text to be added to the paragraph.
//
// Returns:
//   - p: The created Paragraph instance.
func (rd *RootDoc) AddParagraph(text string) *docxpara.Paragraph {
	p := docxpara.AddParagraph(text)
	bodyElem := DocumentChild{
		Para: p,
	}
	rd.Document.Body.Children = append(rd.Document.Body.Children, bodyElem)

	return p
}

// AddEmptyParagraph adds a new empty paragraph to the document.
// It returns the created Paragraph instance.
//
// Returns:
//   - p: The created Paragraph instance.
func (rd *RootDoc) AddEmptyParagraph() *docxpara.Paragraph {
	p := &docxpara.Paragraph{
		Children: []*docxpara.ParagraphChild{},
	}

	bodyElem := DocumentChild{
		Para: p,
	}
	rd.Document.Body.Children = append(rd.Document.Body.Children, bodyElem)

	return p
}
