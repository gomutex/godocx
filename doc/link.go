package doc

import (
	"strconv"

	"github.com/gomutex/godocx/common/constants"
)

func (doc *Document) addLinkRelation(link string) string {

	rID := doc.IncRelationID()

	rel := &Relationship{
		ID:         "rId" + strconv.Itoa(rID),
		TargetMode: "External",
		Type:       constants.SourceRelationshipHyperLink,
		Target:     link,
	}

	doc.DocRels.Relationships = append(doc.DocRels.Relationships, rel)

	return "rId" + strconv.Itoa(rID)
}

func (doc *Document) addRelation(relType string, fileName string) string {
	rID := doc.IncRelationID()
	rel := &Relationship{
		ID:     "rId" + strconv.Itoa(rID),
		Type:   relType,
		Target: fileName,
	}

	doc.DocRels.Relationships = append(doc.DocRels.Relationships, rel)

	return "rId" + strconv.Itoa(rID)
}
