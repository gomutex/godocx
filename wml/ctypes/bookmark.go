package ctypes

import (
	"encoding/xml"
	"strconv"
)

type Bookmark struct {
	Start *BookmarkStart
	End   *BookmarkEnd
}

func (b *Bookmark) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	if b.Start != nil {
		return b.Start.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:bookmarkStart"},
		})
	}
	if b.End != nil {
		return b.End.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:bookmarkEnd"},
		})
	}
	return nil
}

func (b *Bookmark) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
loop:
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "bookmarkStart":
				b.Start = &BookmarkStart{}
				if err = d.DecodeElement(b.Start, &elem); err != nil {
					return err
				}
			case "bookmarkEnd":
				b.End = &BookmarkEnd{}
				if err = d.DecodeElement(b.End, &elem); err != nil {
					return err
				}
			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break loop
		}
	}

	return nil
}

type BookmarkStart struct {
	XMLName  xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main bookmarkStart,omitempty"`
	ID       int      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main id,attr,omitempty"`
	Name     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main name,attr,omitempty"`
	ColFirst int      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main colFirst,attr,omitempty"`
	ColLast  int      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main colLast,attr,omitempty"`
}

func (b *BookmarkStart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "w:id"}, Value: strconv.Itoa(b.ID)},
		{Name: xml.Name{Local: "w:name"}, Value: b.Name},
	}

	if b.ColFirst > 0 {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:colFirst"}, Value: strconv.Itoa(b.ColFirst)})
	}
	if b.ColLast > 0 {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:colLast"}, Value: strconv.Itoa(b.ColLast)})
	}

	return e.EncodeElement("", start)
}

type BookmarkEnd struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main bookmarkEnd,omitempty"`
	ID      int      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main id,attr,omitempty"`
}

func (b *BookmarkEnd) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "w:id"}, Value: strconv.Itoa(b.ID)},
	}

	return e.EncodeElement("", start)
}
