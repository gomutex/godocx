package packager

import (
	"encoding/xml"

	"github.com/gomutex/godocx/common/constants"
	"github.com/gomutex/godocx/doc"
)

// LoadRelationShips loads the relationships from the specified file.
func LoadRelationShips(fileName string, fileBytes []byte) (*doc.Relationships, error) {
	rels := doc.Relationships{Xmlns: constants.XMLNS_R}
	err := xml.Unmarshal(fileBytes, &rels)
	if err != nil {
		return nil, err
	}
	rels.RelativePath = fileName
	return &rels, nil
}

// LoadContentTypes loads the content type from the content types file
func LoadContentTypes(fileBytes []byte) (*doc.ContentTypes, error) {
	ct := doc.ContentTypes{}
	err := xml.Unmarshal(fileBytes, &ct)
	if err != nil {
		return nil, err
	}
	return &ct, nil
}
