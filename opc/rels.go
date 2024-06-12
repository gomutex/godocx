package opc

import (
	"encoding/xml"

	"github.com/gomutex/godocx/common/constants"
	"github.com/gomutex/godocx/ooxml"
)

// LoadRelationShips loads the relationships from the specified file.
func LoadRelationShips(fileName string, fileBytes []byte) (*ooxml.Relationships, error) {
	rels := ooxml.Relationships{Xmlns: constants.XMLNS_R}
	err := xml.Unmarshal(fileBytes, &rels)
	if err != nil {
		return nil, err
	}
	rels.RelativePath = fileName
	return &rels, nil
}
