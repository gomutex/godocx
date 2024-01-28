package oxml

import (
	"strconv"

	"github.com/gomutex/godocx/constants"
)

// addLinkRelation adds a link relation to the RootDoc's relationships.
// This function is used to establish a relationship between the current document
// and an external resource, typically a hyperlink.
//
// Parameters:
//   - rd: A pointer to the RootDoc struct representing the root document of the Word file.
//   - link: The target link URL to establish a relationship with.
//
// Returns:
//   - The ID of the newly created relationship, which can be used to reference this relationship.
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
