package ooxml

import (
	"encoding/xml"
	"strconv"

	"github.com/gomutex/godocx/common/constants"
	"github.com/gomutex/godocx/ooxml/wml"
)

// This element specifies the contents of the body of the document – the main document editing surface.
type Body struct {
	XMLName  xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main body"`
	Children []DocumentChild
	SectPr   *CTSectPr
}

// DocumentChild represents a child element within a Word document, which can be a Paragraph or a Table.
type DocumentChild struct {
	Para  *wml.Paragraph
	Table *wml.Table
}

// CTSectPr represents the section properties of a Word document.
type CTSectPr struct {
	Type      *SectionType    `xml:"type,omitempty"`
	PgSz      *PageSize       `xml:"pgSz,omitempty"`
	PgMar     *PageMargins    `xml:"pgMar,omitempty"`
	PgNumType *PageNumbering  `xml:"pgNumType,omitempty"`
	FormProt  *FormProtection `xml:"formProt,omitempty"`
	TextDir   *TextDirection  `xml:"textDirection,omitempty"`
	DocGrid   *DocGrid        `xml:"docGrid,omitempty"` // Add DocGrid field

}

// DocGrid represents the document grid settings.
type DocGrid struct {
	Type      string `xml:"type,attr,omitempty"`
	LinePitch int    `xml:"linePitch,attr,omitempty"`
	CharSpace int    `xml:"charSpace,attr,omitempty"`
}

// SectionType represents the type of section in a Word document.
type SectionType struct {
	Val string `xml:"val,attr,omitempty"`
}

// PageSize represents the page size of a Word document.
type PageSize struct {
	W int `xml:"w,attr,omitempty"`
	H int `xml:"h,attr,omitempty"`
}

// PageMargins represents the page margins of a Word document.
type PageMargins struct {
	Left   int `xml:"left,attr,omitempty"`
	Right  int `xml:"right,attr,omitempty"`
	Gutter int `xml:"gutter,attr,omitempty"`
	Header int `xml:"header,attr,omitempty"`
	Top    int `xml:"top,attr,omitempty"`
	Footer int `xml:"footer,attr,omitempty"`
	Bottom int `xml:"bottom,attr,omitempty"` // Add Bottom attribute
}

// PageNumbering represents the page numbering format in a Word document.
type PageNumbering struct {
	Fmt string `xml:"fmt,attr,omitempty"`
}

// FormProtection represents the form protection settings in a Word document.
type FormProtection struct {
	Val string `xml:"val,attr,omitempty"`
}

// TextDirection represents the text direction settings in a Word document.
type TextDirection struct {
	Val string `xml:"val,attr,omitempty"`
}

// Use this function to initialize a new Body before adding content to it.
func NewBody() *Body {
	return &Body{}
}

// MarshalXML implements the xml.Marshaler interface for the Body type.
// It encodes the Body to its corresponding XML representation.
func (b *Body) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:body"

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	if b.Children != nil {
		for _, child := range b.Children {
			if child.Para != nil {
				if err = child.Para.MarshalXML(e, xml.StartElement{}); err != nil {
					return err
				}
			}

			if child.Table != nil {
				if err = child.Table.MarshalXML(e, xml.StartElement{}); err != nil {
					return err
				}
			}
		}
	}

	if b.SectPr != nil {
		if err = e.EncodeElement(b.SectPr, xml.StartElement{Name: xml.Name{Local: "w:sectPr"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// MarshalXML implements the xml.Marshaler interface for the CTSectPr type.
// It encodes the CTSectPr to its corresponding XML representation.
func (sectPr *CTSectPr) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:sectPr"

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	if sectPr.Type != nil {
		if err = e.EncodeElement(sectPr.Type, xml.StartElement{Name: xml.Name{Local: "w:type"}}); err != nil {
			return err
		}
	}

	if sectPr.PgSz != nil {
		if err = e.EncodeElement(sectPr.PgSz, xml.StartElement{Name: xml.Name{Local: "w:pgSz"}}); err != nil {
			return err
		}
	}

	if sectPr.PgMar != nil {
		if err = e.EncodeElement(sectPr.PgMar, xml.StartElement{Name: xml.Name{Local: "w:pgMar"}}); err != nil {
			return err
		}
	}

	if sectPr.PgNumType != nil {
		if err = e.EncodeElement(sectPr.PgNumType, xml.StartElement{Name: xml.Name{Local: "w:pgNumType"}}); err != nil {
			return err
		}
	}

	if sectPr.FormProt != nil {
		if err = e.EncodeElement(sectPr.FormProt, xml.StartElement{Name: xml.Name{Local: "w:formProt"}}); err != nil {
			return err
		}
	}

	if sectPr.TextDir != nil {
		if err = e.EncodeElement(sectPr.TextDir, xml.StartElement{Name: xml.Name{Local: "w:textDirection"}}); err != nil {
			return err
		}
	}

	if sectPr.DocGrid != nil {
		if err = e.EncodeElement(sectPr.DocGrid, xml.StartElement{Name: xml.Name{Local: "w:docGrid"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// MarshalXML implements the xml.Marshaler interface for the DocGrid type.
// It encodes the DocGrid to its corresponding XML representation.
func (docGrid *DocGrid) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:docGrid"
	if docGrid.Type != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:type"}, Value: docGrid.Type})
	}
	if docGrid.LinePitch != 0 {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:linePitch"}, Value: strconv.Itoa(docGrid.LinePitch)})
	}
	if docGrid.CharSpace != 0 {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:charSpace"}, Value: strconv.Itoa(docGrid.CharSpace)})
	}
	return e.EncodeElement("", start)
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
			switch elem.Name {
			case xml.Name{Space: constants.WMLNamespace, Local: "p"}, xml.Name{Space: constants.AltWMLNamespace, Local: "p"}:
				para := wml.DefaultParagraph()
				if err := d.DecodeElement(para, &elem); err != nil {
					return err
				}
				body.Children = append(body.Children, DocumentChild{Para: para})
			case xml.Name{Space: constants.WMLNamespace, Local: "tbl"}, xml.Name{Space: constants.AltWMLNamespace, Local: "tbl"}:
				tbl := wml.DefaultTable()
				if err := d.DecodeElement(tbl, &elem); err != nil {
					return err
				}
				body.Children = append(body.Children, DocumentChild{Table: tbl})
			case xml.Name{Space: constants.WMLNamespace, Local: "sectPr"}, xml.Name{Space: constants.AltWMLNamespace, Local: "sectPr"}:
				body.SectPr = &CTSectPr{}
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

// getSectionTypeVal returns an xml.Attr representing the SectionType value.
// It maps a string value to the corresponding SectionType XML attribute.
func getSectionTypeVal(value string) xml.Attr {
	valElem := xml.Attr{}
	valElem.Name = xml.Name{Local: "w:val"}

	switch value {
	case "nextPage":
		valElem.Value = "nextPage"
	case "nextColumn":
		valElem.Value = "nextColumn"
	case "continuous":
		valElem.Value = "continuous"
	case "evenPage":
		valElem.Value = "evenPage"
	case "oddPage":
		valElem.Value = "oddPage"
	default:
		valElem.Value = ""
	}
	return valElem
}

// MarshalXML implements the xml.Marshaler interface for the SectionType type.
// It encodes the SectionType to its corresponding XML representation.
func (st *SectionType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	valElem := getSectionTypeVal(st.Val)
	start.Attr = append(start.Attr, valElem)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

// MarshalXML implements the xml.Marshaler interface for the PageSize type.
// It encodes the PageSize to its corresponding XML representation.
func (pgSz *PageSize) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:pgSz"
	if pgSz.W != 0 {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:w"}, Value: strconv.Itoa(pgSz.W)})
	}
	if pgSz.H != 0 {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:h"}, Value: strconv.Itoa(pgSz.H)})
	}
	return e.EncodeElement("", start)
}

// MarshalXML implements the xml.Marshaler interface for the PageMargins type.
// It encodes the PageMargins to its corresponding XML representation.
func (pgMar *PageMargins) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:pgMar"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:left"}, Value: strconv.Itoa(pgMar.Left)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:right"}, Value: strconv.Itoa(pgMar.Right)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:gutter"}, Value: strconv.Itoa(pgMar.Gutter)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:header"}, Value: strconv.Itoa(pgMar.Header)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:top"}, Value: strconv.Itoa(pgMar.Top)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:footer"}, Value: strconv.Itoa(pgMar.Footer)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:bottom"}, Value: strconv.Itoa(pgMar.Bottom)})
	return e.EncodeElement("", start)
}

// MarshalXML implements the xml.Marshaler interface for the PageNumbering type.
// It encodes the PageNumbering to its corresponding XML representation.
func (pgNumType *PageNumbering) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:pgNumType"
	if pgNumType.Fmt != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:fmt"}, Value: pgNumType.Fmt})
	}
	return e.EncodeElement("", start)
}

// MarshalXML implements the xml.Marshaler interface for the FormProtection type.
// It encodes the FormProtection to its corresponding XML representation.
func (formProt *FormProtection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:formProt"
	if formProt.Val != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: formProt.Val})
	}
	return e.EncodeElement("", start)
}

// MarshalXML implements the xml.Marshaler interface for the TextDirection type.
// It encodes the TextDirection to its corresponding XML representation.
func (textDir *TextDirection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:textDirection"
	if textDir.Val != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: textDir.Val})
	}
	return e.EncodeElement("", start)
}
