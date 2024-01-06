package oxml

import (
	"strconv"

	"github.com/gomutex/godocx/constants"
)

func (rd *RootDoc) addLinkRelation(link string) string {

	rel := &Relationship{
		ID:         "rId" + strconv.Itoa(rd.rId),
		TargetMode: "External",
		Type:       constants.SourceRelationshipHyperLink,
		Target:     link,
	}

	rd.rId += 1

	rd.DocRels.Relationships = append(rd.DocRels.Relationships, rel)

	return rel.ID
}
