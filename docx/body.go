package docx

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/ctypes"
)

// This element specifies the contents of the body of the document – the main document editing surface.
type Body struct {
	root     *RootDoc
	XMLName  xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main body"`
	Children []DocumentChild
	SectPr   *ctypes.SectionProp
}

// DocumentChild represents a child element within a Word document, which can be a Paragraph or a Table.
type DocumentChild struct {
	Para  *Paragraph
	Table *Table
}

// Use this function to initialize a new Body before adding content to it.
func NewBody(root *RootDoc) *Body {
	return &Body{
		root: root,
	}
}

// MarshalXML implements the xml.Marshaler interface for the Body type.
// It encodes the Body to its corresponding XML representation.
func (b Body) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:body"

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	if b.Children != nil {
		for _, child := range b.Children {
			if child.Para != nil {
				if err = child.Para.ct.MarshalXML(e, xml.StartElement{}); err != nil {
					return err
				}
			}

			if child.Table != nil {
				if err = child.Table.ct.MarshalXML(e, xml.StartElement{}); err != nil {
					return err
				}
			}
		}
	}

	if b.SectPr != nil {
		if err = b.SectPr.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements the xml.Unmarshaler interface for the Body type.
// It decodes the XML representation of the Body.
func (body *Body) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {

	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "p":
				para := newParagraph(body.root)
				if err := para.unmarshalXML(d, elem); err != nil {
					return err
				}
				body.Children = append(body.Children, DocumentChild{Para: para})
			case "tbl":
				tbl := NewTable(body.root)
				if err := tbl.unmarshalXML(d, elem); err != nil {
					return err
				}
				body.Children = append(body.Children, DocumentChild{Table: tbl})
			case "sectPr":
				body.SectPr = ctypes.NewSectionProper()
				if err := d.DecodeElement(body.SectPr, &elem); err != nil {
					return err
				}
			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			return nil
		}
	}
}
