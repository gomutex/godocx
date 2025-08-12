package ctypes

import "encoding/xml"

type Hyperlink struct {
	XMLName  xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main hyperlink,omitempty"`
	ID       string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr,omitempty"`
	Anchor   string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main anchor,attr,omitempty"`
	Run      *Run
	Children []ParagraphChild
}

func NewHyperlink() *Hyperlink {
	return &Hyperlink{}
}

func (h Hyperlink) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:hyperlink"

	if h.ID != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"}, Value: h.ID})
	}

	if h.Anchor != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:anchor"}, Value: h.Anchor})
	}

	if err = e.EncodeToken(start); err != nil {
		return err
	}

	if h.Run != nil {
		if err = h.Run.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:r"},
		}); err != nil {
			return err
		}
	}

	for _, child := range h.Children {
		if child.Run != nil {
			if err = child.Run.MarshalXML(e, xml.StartElement{
				Name: xml.Name{Local: "w:r"},
			}); err != nil {
				return err
			}
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (h *Hyperlink) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	h.XMLName = start.Name

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "id":
			h.ID = attr.Value
		case "anchor":
			h.Anchor = attr.Value
		}
	}

loop:
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "r":
				run := NewRun()
				if err = d.DecodeElement(run, &elem); err != nil {
					return err
				}
				if h.Run == nil {
					h.Run = run
				} else {
					h.Children = append(h.Children, ParagraphChild{
						Run: run,
					})
				}
			case "bookmarkStart":
				bookmarkStart := &BookmarkStart{}
				if err = d.DecodeElement(bookmarkStart, &elem); err != nil {
					return err
				}
				h.Children = append(h.Children, ParagraphChild{
					Bookmark: &Bookmark{
						Start: bookmarkStart,
					},
				})
			case "bookmarkEnd":
				bookmarkEnd := &BookmarkEnd{}
				if err = d.DecodeElement(bookmarkEnd, &elem); err != nil {
					return err
				}
				h.Children = append(h.Children, ParagraphChild{
					Bookmark: &Bookmark{
						End: bookmarkEnd,
					},
				})
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
