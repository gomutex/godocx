package docx

import (
	"strconv"

	"github.com/gomutex/godocx/common/constants"
)

// addLinkRelation adds a hyperlink relationship to the document's relationships collection.
//
// Parameters:
//   - link: A string representing the target URL or location of the hyperlink.
//
// Returns:
//   - string: The ID ("rId" + relation ID) of the added relationship.
//
// This function generates a new relationship ID, creates a Relationship object with the specified link as the target,
// and appends it to the document's relationships collection (DocRels.Relationships). It returns the generated ID of the relationship.
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

// addRelation adds a generic relationship to the document's relationships collection.
//
// Parameters:
//   - relType: A string representing the type of relationship (e.g., constants.SourceRelationshipImage).
//   - fileName: A string representing the target file name or location related to the relationship.
//
// Returns:
//   - string: The ID ("rId" + relation ID) of the added relationship.
//
// This function generates a new relationship ID, creates a Relationship object with the specified type and target,
// and appends it to the document's relationships collection (DocRels.Relationships). It returns the generated ID of the relationship.
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
