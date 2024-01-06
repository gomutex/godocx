package constants

import "bytes"

func TranslateNamespace(content []byte) []byte {
	namespaceTranslationDic := map[string]string{
		StrictNameSpaceDocumentPropertiesVariantTypes: NameSpaceDocumentPropertiesVariantTypes.Value,
		StrictNameSpaceDrawingMLMain:                  NameSpaceDrawingMLMain,
		StrictNameSpaceExtendedProperties:             NameSpaceExtendedProperties,
		StrictSourceRelationship:                      SourceRelationship.Value,
		StrictSourceRelationshipChart:                 SourceRelationshipChart,
		StrictSourceRelationshipComments:              SourceRelationshipComments,
		StrictSourceRelationshipExtendProperties:      SourceRelationshipExtendProperties,
		StrictSourceRelationshipImage:                 SourceRelationshipImage,
		StrictSourceRelationshipOfficeDocument:        SourceRelationshipOfficeDocument,
	}
	for s, n := range namespaceTranslationDic {
		content = replaceBytes(content, []byte(s), []byte(n), -1)
	}
	return content
}

// replaceBytes replace source bytes with given target.
func replaceBytes(s, source, target []byte, n int) []byte {
	if n == 0 {
		return s
	}

	if len(source) < len(target) {
		return bytes.Replace(s, source, target, n)
	}

	if n < 0 {
		n = len(s)
	}

	var wid, i, j, w int
	for i, j = 0, 0; i < len(s) && j < n; j++ {
		wid = bytes.Index(s[i:], source)
		if wid < 0 {
			break
		}

		w += copy(s[w:], s[i:i+wid])
		w += copy(s[w:], target)
		i += wid + len(source)
	}

	w += copy(s[w:], s[i:])
	return s[:w]
}
