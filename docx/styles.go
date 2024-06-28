package docx

import (
	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/stypes"
)

func (rd *RootDoc) GetStyleByID(styleID string, styleType stypes.StyleType) *ctypes.Style {
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
