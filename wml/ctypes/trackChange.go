package ctypes

import (
	"encoding/xml"
	"strconv"
)

// TrackChange represents the complex type for track change
type TrackChange struct {
	ID     int     `xml:"id,attr"`
	Author string  `xml:"author,attr"`
	Date   *string `xml:"date,attr,omitempty"`
}

func (t TrackChange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "w:id"}, Value: strconv.Itoa(t.ID)},
		{Name: xml.Name{Local: "w:author"}, Value: t.Author},
	}

	if t.Date != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:date"}, Value: *t.Date})
	}

	return e.EncodeElement("", start)
}
