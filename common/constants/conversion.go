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

var NSToLocal = map[string]string{
	// Document Properties
	"http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes": "vt",

	// DrawingML
	"http://schemas.openxmlformats.org/drawingml/2006/main":    "a",
	"http://schemas.openxmlformats.org/drawingml/2006/chart":   "c",
	"http://schemas.openxmlformats.org/drawingml/2006/diagram": "dgm",

	// DrawingML Compatibility
	"http://schemas.openxmlformats.org/drawingml/2006/compatibility": "compat",

	// DrawingML Picture
	"http://schemas.openxmlformats.org/drawingml/2006/picture": "pic",

	// Microsoft Drawing
	"http://schemas.microsoft.com/office/drawing/2010/main":          "a14",
	"http://schemas.microsoft.com/office/drawing/2010/compatibility": "compat14",
	"http://schemas.microsoft.com/office/drawing/2016/SVG/main":      "asvg",

	// Microsoft Office Relationships
	"http://schemas.openxmlformats.org/officeDocument/2006/relationships": "r",
	"http://schemas.microsoft.com/office/2006/relationships":              "r2006",
	"http://schemas.microsoft.com/office/2011/relationships":              "r2011",

	// Microsoft Office Compatibility Relationships
	"http://schemas.openxmlformats.org/officeDocument/2006/relationships/slicer": "sle",
	"http://schemas.microsoft.com/office/drawing/2010/relationships":             "r14",

	// Markup Compatibility
	"http://schemas.openxmlformats.org/markup-compatibility/2006": "mc",

	// Theme
	"http://schemas.openxmlformats.org/drawingml/2006/main/theme": "thm",

	// Word Processing
	"http://schemas.openxmlformats.org/wordprocessingml/2006/main": "w",

	// Word Processing Styles
	"http://schemas.openxmlformats.org/officeDocument/2006/styles": "s",
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
