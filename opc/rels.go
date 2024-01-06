package opc

import (
	"encoding/xml"

	"github.com/gomutex/godocx/constants"
	"github.com/gomutex/godocx/oxml"
)

func LoadRelationShips(fileName string, fileBytes []byte) (*oxml.Relationships, error) {
	rels := oxml.Relationships{Xmlns: constants.XMLNS_R}
	err := xml.Unmarshal(fileBytes, &rels)
	if err != nil {
		return nil, err
	}
	rels.RelativePath = fileName
	return &rels, nil
}
