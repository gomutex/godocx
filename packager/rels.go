package packager

import (
	"encoding/xml"

	"github.com/gomutex/godocx/common/constants"
	"github.com/gomutex/godocx/docx"
)

// LoadRelationShips loads the relationships from the specified file.
func LoadRelationShips(fileName string, fileBytes []byte) (*docx.Relationships, error) {
	rels := docx.Relationships{Xmlns: constants.XMLNS_R}
	err := xml.Unmarshal(fileBytes, &rels)
	if err != nil {
		return nil, err
	}
	rels.RelativePath = fileName
	return &rels, nil
}

// LoadContentTypes loads the content type from the content types file
func LoadContentTypes(fileBytes []byte) (*docx.ContentTypes, error) {
	ct := docx.ContentTypes{}
	err := xml.Unmarshal(fileBytes, &ct)
	if err != nil {
		return nil, err
	}
	return &ct, nil
}
