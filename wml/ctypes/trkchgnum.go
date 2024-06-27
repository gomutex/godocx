package ctypes

import (
	"encoding/xml"
	"strconv"
)

// TrackChangeNum represents the complex type for track change numbering
type TrackChangeNum struct {
	ID       int     `xml:"id,attr"`
	Author   string  `xml:"author,attr"`
	Date     *string `xml:"date,attr,omitempty"`
	Original *string `xml:"original,attr,omitempty"`
}

func (t TrackChangeNum) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "w:id"}, Value: strconv.Itoa(t.ID)},
		{Name: xml.Name{Local: "w:author"}, Value: t.Author},
	}

	if t.Date != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:date"}, Value: *t.Date})
	}

	if t.Original != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:original"}, Value: *t.Original})
	}

	return e.EncodeElement("", start)
}
