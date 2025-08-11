package docx

import (
	"encoding/xml"
	"sync"

	"github.com/gomutex/godocx/wml/ctypes"
)

// RootDoc represents the root document of an Office Open XML (OOXML) document.
// It contains information about the document path, file map, the document structure,
// and relationships with other parts of the document.
type RootDoc struct {
	Path        string        // Path represents the path of the document.
	FileMap     sync.Map      // FileMap is a synchronized map for managing files related to the document.
	RootRels    Relationships // RootRels represents relationships at the root level.
	ContentType ContentTypes
	Document    *Document      // Document is the main document structure.
	DocStyles   *ctypes.Styles // Document styles
	Numbering   *NumberingManager // Numbering manager for list instances

	rID        int // rId is used to generate unique relationship IDs.
	ImageCount uint
}

// NewRootDoc creates a new instance of the RootDoc structure.
func NewRootDoc() *RootDoc {
	root := &RootDoc{}
	root.Numbering = NewNumberingManager(root)
	return root
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
func LoadDocXml(rd *RootDoc, fileName string, fileBytes []byte) (*Document, error) {
	doc := Document{
		Root: rd,
	}
	err := xml.Unmarshal(fileBytes, &doc)
	if err != nil {
		return nil, err
	}

	doc.relativePath = fileName
	return &doc, nil
}

// Load styles.xml into Styles struct
func LoadStyles(fileName string, fileBytes []byte) (*ctypes.Styles, error) {
	styles := ctypes.Styles{}
	err := xml.Unmarshal(fileBytes, &styles)
	if err != nil {
		return nil, err
	}

	styles.RelativePath = fileName
	return &styles, nil
}

// NewListInstance creates a new numbering instance for the given abstract numbering ID.
// Returns the numId that can be used with paragraph.Numbering().
//
// Parameters:
//   - abstractNumId: An integer representing the abstract numbering definition ID (1-8 or custom).
//
// Returns:
//   - numId: An integer that can be used with paragraph.Numbering(numId, level).
//
// Example:
//
//	numId := doc.NewListInstance(1)
//	p.Numbering(numId, 0)
func (root *RootDoc) NewListInstance(abstractNumId int) int {
	return root.Numbering.NewListInstance(abstractNumId)
}
