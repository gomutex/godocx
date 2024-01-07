package oxml

import (
	"encoding/xml"
	"strconv"

	"github.com/gomutex/godocx/constants"
	"github.com/gomutex/godocx/oxml/elements"
)

type DocumentChild struct {
	Para  *elements.Paragraph
	Table *Table
}

type Body struct {
	XMLName  xml.Name         `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main body"`
	Children []*DocumentChild `xml:",any"`

	SectPr *CTSectPr
}

type CTSectPr struct {
	Type      *SectionType    `xml:"type,omitempty"`
	PgSz      *PageSize       `xml:"pgSz,omitempty"`
	PgMar     *PageMargins    `xml:"pgMar,omitempty"`
	PgNumType *PageNumbering  `xml:"pgNumType,omitempty"`
	FormProt  *FormProtection `xml:"formProt,omitempty"`
	TextDir   *TextDirection  `xml:"textDirection,omitempty"`
	DocGrid   *DocGrid        `xml:"docGrid,omitempty"` // Add DocGrid field

}

type DocGrid struct {
	Type      string `xml:"type,attr,omitempty"`
	LinePitch int    `xml:"linePitch,attr,omitempty"`
	CharSpace int    `xml:"charSpace,attr,omitempty"`
}

type SectionType struct {
	Val string `xml:"val,attr,omitempty"`
}

type PageSize struct {
	W int `xml:"w,attr,omitempty"`
	H int `xml:"h,attr,omitempty"`
}

type PageMargins struct {
	Left   int `xml:"left,attr,omitempty"`
	Right  int `xml:"right,attr,omitempty"`
	Gutter int `xml:"gutter,attr,omitempty"`
	Header int `xml:"header,attr,omitempty"`
	Top    int `xml:"top,attr,omitempty"`
	Footer int `xml:"footer,attr,omitempty"`
	Bottom int `xml:"bottom,attr,omitempty"` // Add Bottom attribute
}

type PageNumbering struct {
	Fmt string `xml:"fmt,attr,omitempty"`
}

type FormProtection struct {
	Val string `xml:"val,attr,omitempty"`
}

type TextDirection struct {
	Val string `xml:"val,attr,omitempty"`
}

func NewBody() *Body {
	return &Body{}
}

func (b *Body) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:body"

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	if b.Children != nil {
		for _, cElem := range b.Children {
			if cElem.Para != nil {
				if err = cElem.Para.MarshalXML(e, xml.StartElement{}); err != nil {
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

func (body *Body) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {

loop:
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name {
			case xml.Name{Space: constants.WMLNamespace, Local: "p"}, xml.Name{Space: constants.AltWMLNamespace, Local: "p"}:
				para := elements.NewParagraph()
				if err := d.DecodeElement(para, &elem); err != nil {
					return err
				}
				body.Children = append(body.Children, &DocumentChild{Para: para})

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
			break loop
		}
	}

	return nil
}

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

func (st *SectionType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	valElem := getSectionTypeVal(st.Val)
	start.Attr = append(start.Attr, valElem)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

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

func (pgNumType *PageNumbering) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:pgNumType"
	if pgNumType.Fmt != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:fmt"}, Value: pgNumType.Fmt})
	}
	return e.EncodeElement("", start)
}

func (formProt *FormProtection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:formProt"
	if formProt.Val != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: formProt.Val})
	}
	return e.EncodeElement("", start)
}

func (textDir *TextDirection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:textDirection"
	if textDir.Val != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: textDir.Val})
	}
	return e.EncodeElement("", start)
}
