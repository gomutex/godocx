package opc

import (
	"archive/zip"
	"bytes"
	"fmt"
	"strings"

	"github.com/gomutex/godocx/constants"
	"github.com/gomutex/godocx/internal"
	"github.com/gomutex/godocx/oxml"
)

// ReadFromZip reads files from a zip archive.
func ReadFromZip(content *[]byte) (map[string][]byte, error) {
	zipReader, err := zip.NewReader(bytes.NewReader(*content), int64(len(*content)))
	if err != nil {
		return nil, err
	}

	var (
		fileList = make(map[string][]byte, len(zipReader.File))
	)

	for _, f := range zipReader.File {

		fileName := strings.ReplaceAll(f.Name, "\\", "/")

		if fileList[fileName], err = internal.ReadFile(f); err != nil {
			return nil, err
		}
	}

	return fileList, nil
}

func Unpack(content *[]byte) (*oxml.RootDoc, error) {

	rd := oxml.NewRootDoc()

	fileIndex, err := ReadFromZip(content)
	if err != nil {
		return nil, err
	}

	for fileName, fileContent := range fileIndex {
		rd.FileMap.Store(fileName, fileContent)
	}

	rootRelURI, err := GetRelsURI("")
	if err != nil {
		return nil, err
	}

	rootRelBytes := fileIndex[*rootRelURI]
	rootRelations, err := LoadRelationShips(*rootRelURI, rootRelBytes)
	if err != nil {
		return nil, err
	}
	delete(fileIndex, *rootRelURI)
	rd.RootRels = *rootRelations

	var docPath *string

	for _, relation := range rootRelations.Relationships {
		switch relation.Type {
		case constants.OFFICE_DOC_TYPE:
			docPath = &relation.Target
		}
	}

	if docPath == nil {
		return nil, fmt.Errorf("root officeDocument type not found")
	}

	docRelURI, err := GetRelsURI(*docPath)
	if err != nil {
		return nil, err
	}

	docRelFile := fileIndex[*docRelURI]
	docRelations, err := LoadRelationShips(*docRelURI, docRelFile)
	if err != nil {
		return nil, err
	}
	delete(fileIndex, *rootRelURI)
	rd.DocRels = *docRelations

	docFile := fileIndex[*docPath]
	doc, err := oxml.LoadDocXml(*docPath, docFile)
	if err != nil {
		return nil, err
	}
	delete(fileIndex, *docPath)

	rd.Document = doc

	return rd, nil
}
