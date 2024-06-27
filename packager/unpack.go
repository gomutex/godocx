package packager

import (
	"archive/zip"
	"bytes"
	"fmt"
	"strings"

	"github.com/gomutex/godocx/common/constants"
	"github.com/gomutex/godocx/docx"
	"github.com/gomutex/godocx/internal"
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

		if fileList[fileName], err = internal.ReadFileFromZip(f); err != nil {
			return nil, err
		}
	}

	return fileList, nil
}

func Unpack(content *[]byte) (*docx.RootDoc, error) {

	rd := docx.NewRootDoc()

	fileIndex, err := ReadFromZip(content)
	if err != nil {
		return nil, err
	}

	// Load content type details
	ctBytes := fileIndex[constants.ConentTypeFileIdx]
	ct, err := LoadContentTypes(ctBytes)
	if err != nil {
		return nil, err
	}
	delete(fileIndex, constants.ConentTypeFileIdx)
	rd.ContentType = *ct

	rd.ImageCount = 0
	for fileName, fileContent := range fileIndex {
		if strings.HasPrefix(fileName, constants.MediaPath) {
			rd.ImageCount += 1
		}
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

	var docPath string
	var stylesPath string

	for _, relation := range rootRelations.Relationships {
		switch relation.Type {
		case constants.OFFICE_DOC_TYPE:
			docPath = relation.Target
		case constants.StylesType:
			stylesPath = relation.Target
		}
	}

	if docPath == "" {
		return nil, fmt.Errorf("root officeDocument type not found")
	}

	docRelURI, err := GetRelsURI(docPath)
	if err != nil {
		return nil, err
	}

	// Load document
	docFile := fileIndex[docPath]
	docObj, err := docx.LoadDocXml(docPath, docFile)
	if err != nil {
		return nil, err
	}
	delete(fileIndex, docPath)
	rd.Document = docObj

	//Load Styles
	stylesFile := fileIndex[stylesPath]
	stylesObj, err := docx.LoadStyles(stylesPath, stylesFile)
	if err != nil {
		return nil, err
	}
	delete(fileIndex, stylesPath)
	rd.DocStyles = stylesObj

	// Load Relationship details
	docRelFile := fileIndex[*docRelURI]
	docRelations, err := LoadRelationShips(*docRelURI, docRelFile)
	if err != nil {
		return nil, err
	}
	delete(fileIndex, *rootRelURI)
	rd.Document.DocRels = *docRelations
	rID := 0
	for range docRelations.Relationships {
		rID += 1
	}
	rd.Document.RID = rID

	return rd, nil
}
