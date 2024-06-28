package ctypes

import (
	"encoding/xml"
	"strconv"
)

// PageMargin represents the page margins of a Word document.
type PageMargin struct {
	Left   *int `xml:"left,attr,omitempty"`
	Right  *int `xml:"right,attr,omitempty"`
	Gutter *int `xml:"gutter,attr,omitempty"`
	Header *int `xml:"header,attr,omitempty"`
	Top    *int `xml:"top,attr,omitempty"`
	Footer *int `xml:"footer,attr,omitempty"`
	Bottom *int `xml:"bottom,attr,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for the PageMargin type.
// It encodes the PageMargin to its corresponding XML representation.
func (p PageMargin) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:pgMar"

	start.Attr = []xml.Attr{}

	attrs := []struct {
		local string
		value *int
	}{
		{"w:left", p.Left},
		{"w:right", p.Right},
		{"w:gutter", p.Gutter},
		{"w:header", p.Header},
		{"w:top", p.Top},
		{"w:footer", p.Footer},
		{"w:bottom", p.Bottom},
	}

	for _, attr := range attrs {
		if attr.value == nil {
			continue
		}
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: attr.local}, Value: strconv.Itoa(*attr.value)})
	}

	return e.EncodeElement("", start)
}
