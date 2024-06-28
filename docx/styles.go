package docx

import (
	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/stypes"
)

// GetStyleByID retrieves a style from the document styles collection based on the given style ID and type.
//
// Parameters:
//   - styleID: A string representing the ID of the style to retrieve.
//   - styleType: An stypes.StyleType indicating the type of style (e.g., paragraph, character, table).
//
// Returns:
//   - *ctypes.Style: A pointer to the style matching the provided ID and type, if found; otherwise, nil.
//
// This method searches through the document's style list to find a style with the specified ID and type.
// If no matching style is found or if the document styles collection is nil, it returns nil.
func (rd *RootDoc) GetStyleByID(styleID string, styleType stypes.StyleType) *ctypes.Style {
	if rd.DocStyles == nil {
		return nil
	}

	for _, style := range rd.DocStyles.StyleList {
		if style.ID == nil || style.Type == nil {
			continue
		}

		if *style.ID == styleID && *style.Type == styleType {
			return &style
		}
	}
	return nil
}
