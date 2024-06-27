package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/hdrftr"
	"github.com/gomutex/godocx/wml/stypes"
)

// Document Final Section Properties : w:sectPr
type SectionProp struct {
	HeaderReference *hdrftr.HeaderReference `xml:"headerReference,omitempty"`
	FooterReference *hdrftr.FooterReference `xml:"footerReference,omitempty"`
	PageSize        *PageSize               `xml:"pgSz,omitempty"`
	Type            *SectionType            `xml:"type,omitempty"`
	PageMargin      *PageMargin             `xml:"pgMar,omitempty"`
	PageNum         *PageNumbering          `xml:"pgNumType,omitempty"`
	FormProt        *FormProt               `xml:"formProt,omitempty"`
	TitlePg         *hdrftr.TitlePg         `xml:"titlePg,omitempty"`
	TextDir         *TextDirection          `xml:"textDirection,omitempty"`
	DocGrid         *DocGrid                `xml:"docGrid,omitempty"`
}

func NewSectionProper() *SectionProp {
	return &SectionProp{}
}

func (s SectionProp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:sectPr"

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if s.HeaderReference != nil {
		if err := s.HeaderReference.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if s.FooterReference != nil {
		if err := s.FooterReference.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if s.Type != nil {
		if err := s.Type.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if s.PageSize != nil {
		if err := s.PageSize.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if s.PageMargin != nil {
		if err = s.PageMargin.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if s.PageNum != nil {
		if err = s.PageNum.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if s.FormProt != nil {
		if err = s.FormProt.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if s.TitlePg != nil {
		if err = s.TitlePg.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if s.TextDir != nil {
		if s.TextDir.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if s.DocGrid != nil {
		if s.DocGrid.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// Section Type
type SectionType struct {
	Val stypes.SectionMark `xml:"val,attr,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for the SectionType type.
// It encodes the SectionType to its corresponding XML representation.
func (s SectionType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:type"
	if s.Val != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(s.Val)})
	}
	return e.EncodeElement("", start)
}

// <== SectionType Ends here ==>