package oxml

import (
	"encoding/xml"
	"sync"
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
