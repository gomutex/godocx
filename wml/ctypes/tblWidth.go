package ctypes

import (
	"encoding/xml"
	"strconv"

	"github.com/gomutex/godocx/wml/stypes"
)

// TableWidth represents the width of a table in a document.
type TableWidth struct {
	Width     *int               `xml:"w,attr,omitempty"`
	WidthType *stypes.TableWidth `xml:"type,attr,omitempty"`
}

func NewTableWidth(width int, widthType stypes.TableWidth) *TableWidth {
	return &TableWidth{
		Width:     &width,
		WidthType: &widthType,
	}
}

func (t TableWidth) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {

	if t.Width != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:w"}, Value: strconv.Itoa(*t.Width)})
	}

	if t.WidthType != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:type"}, Value: string(*t.WidthType)})
	}

	return e.EncodeElement("", start)
}
