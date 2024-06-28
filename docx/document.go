package docx

import (
	"encoding/xml"

	"github.com/gomutex/godocx/internal"
	"github.com/gomutex/godocx/wml/stypes"
)

var docAttrs = map[string]string{
	"xmlns:w":      "http://schemas.openxmlformats.org/wordprocessingml/2006/main",
	"xmlns:o":      "urn:schemas-microsoft-com:office:office",
	"xmlns:r":      "http://schemas.openxmlformats.org/officeDocument/2006/relationships",
	"xmlns:v":      "urn:schemas-microsoft-com:vml",
	"xmlns:w10":    "urn:schemas-microsoft-com:office:word",
	"xmlns:wp":     "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing",
	"xmlns:wps":    "http://schemas.microsoft.com/office/word/2010/wordprocessingShape",
	"xmlns:wpg":    "http://schemas.microsoft.com/office/word/2010/wordprocessingGroup",
	"xmlns:mc":     "http://schemas.openxmlformats.org/markup-compatibility/2006",
	"xmlns:wp14":   "http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing",
	"xmlns:w14":    "http://schemas.microsoft.com/office/word/2010/wordml",
	"xmlns:w15":    "http://schemas.microsoft.com/office/word/2012/wordml",
	"mc:Ignorable": "w14 wp14 w15",
}

// This element specifies the contents of a main document part in a WordprocessingML document.
type Document struct {
	// Reference to the RootDoc
	Root *RootDoc

	// Elements
	Background *Background
	Body       *Body

	// Non elements - helper fields
	DocRels      Relationships // DocRels represents relationships specific to the document.
	RID          int
	relativePath string
}

// IncRelationID increments the relation ID of the document and returns the new ID.
// This method is used to generate unique IDs for relationships within the document.
func (doc *Document) IncRelationID() int {
	doc.RID += 1
	return doc.RID
}

// MarshalXML implements the xml.Marshaler interface for the Document type.
func (doc Document) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:document"

	for key, value := range docAttrs {
		attr := xml.Attr{Name: xml.Name{Local: key}, Value: value}
		start.Attr = append(start.Attr, attr)
	}

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	if doc.Background != nil {
		if err = doc.Background.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if doc.Body != nil {
		bodyElement := xml.StartElement{Name: xml.Name{Local: "w:body"}}
		if err = e.EncodeElement(doc.Body, bodyElement); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (d *Document) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) (err error) {

	for {
		currentToken, err := decoder.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "body":
				body := NewBody(d.Root)
				if err := decoder.DecodeElement(body, &elem); err != nil {
					return err
				}
				d.Body = body
			case "background":
				bg := NewBackground()
				if err := decoder.DecodeElement(bg, &elem); err != nil {
					return err
				}
				d.Background = bg
			default:
				if err = decoder.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			return nil
		}
	}

}

// Return newly Paragraph object containing only a page break.
func (rd *RootDoc) AddPageBreak() *Paragraph {
	p := rd.AddEmptyParagraph()
	p.AddRun().AddBreak(internal.ToPtr(stypes.BreakTypePage))

	return p
}
