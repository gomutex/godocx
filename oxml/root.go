package oxml

import (
	"encoding/xml"
	"sync"

	"github.com/gomutex/godocx/oxml/elements"
)

type RootDoc struct {
	Path     string
	FileMap  sync.Map
	Document *Document
	RootRels Relationships
	DocRels  Relationships

	rId int
}

func NewRootDoc() *RootDoc {
	return &RootDoc{}
}

func LoadDocXml(fileName string, fileBytes []byte) (*Document, error) {
	doc := Document{}
	err := xml.Unmarshal(fileBytes, &doc)
	if err != nil {
		return nil, err
	}

	doc.relativePath = fileName
	return &doc, nil
}

func (rd *RootDoc) AddParagraph(text string) *elements.Paragraph {
	p := &elements.Paragraph{
		Children: make([]*elements.ParagraphChild, 0),
	}

	bodyElem := DocumentChild{
		Para: p,
	}
	rd.Document.Body.Children = append(rd.Document.Body.Children, &bodyElem)

	p.AddText("Hello, world!")

	return p
}

func (rd *RootDoc) AddEmptyParagraph() *elements.Paragraph {
	p := &elements.Paragraph{
		Children: make([]*elements.ParagraphChild, 0),
	}

	bodyElem := DocumentChild{
		Para: p,
	}
	rd.Document.Body.Children = append(rd.Document.Body.Children, &bodyElem)

	return p
}
