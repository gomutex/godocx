package sections

import (
	"encoding/xml"
	"strconv"
)

// Section Properties : w:sectPr
type SectionProp struct {
	PageSize   *PageSize       `xml:"pgSz,omitempty"`
	Type       *SectionType    `xml:"type,omitempty"`
	PageMargin *PageMargin     `xml:"pgMar,omitempty"`
	PageNum    *PageNumbering  `xml:"pgNumType,omitempty"`
	FormProt   *FormProtection `xml:"formProt,omitempty"`
	TextDir    *TextDirection  `xml:"textDirection,omitempty"`
	DocGrid    *DocGrid        `xml:"docGrid,omitempty"`
}

func NewSectionProper() *SectionProp {
	return &SectionProp{}
}

func (s *SectionProp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:sectPr"

	err := e.EncodeToken(start)
	if err != nil {
		return err
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

// PageMargin represents the page margins of a Word document.
type PageMargin struct {
	Left   int `xml:"left,attr,omitempty"`
	Right  int `xml:"right,attr,omitempty"`
	Gutter int `xml:"gutter,attr,omitempty"`
	Header int `xml:"header,attr,omitempty"`
	Top    int `xml:"top,attr,omitempty"`
	Footer int `xml:"footer,attr,omitempty"`
	Bottom int `xml:"bottom,attr,omitempty"` // Add Bottom attribute
}

// FormProtection represents the form protection settings in a Word document.
type FormProtection struct {
	Val string `xml:"val,attr,omitempty"`
}

// TextDirection represents the text direction settings in a Word document.
type TextDirection struct {
	Val string `xml:"val,attr,omitempty"`
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

// MarshalXML implements the xml.Marshaler interface for the SectionType type.
// It encodes the SectionType to its corresponding XML representation.
func (st *SectionType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:type"
	valElem := getSectionTypeVal(st.Val)
	start.Attr = append(start.Attr, valElem)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

// MarshalXML implements the xml.Marshaler interface for the PageMargin type.
// It encodes the PageMargin to its corresponding XML representation.
func (pgMar *PageMargin) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
