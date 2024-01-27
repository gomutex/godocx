package elements

import (
	"encoding/xml"

	"github.com/gomutex/godocx/constants"
)

type ParagraphProperty struct {
	DivID           *string
	KeepNext        *bool
	KeepLines       *bool
	PageBreakBefore *bool
	WidowControl    *bool
	Style           *ParagraphStyle
	Justification   *Justification

	NumberingProperty *NumberingProperty
}

func DefaultParaProperty() *ParagraphProperty {
	return &ParagraphProperty{}
}

func (pp *ParagraphProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	elem := xml.StartElement{Name: xml.Name{Local: "w:pPr"}}

	// Opening <w:pPr> element
	if err = e.EncodeToken(elem); err != nil {
		return err
	}

	// Encoding <w:pStyle> element
	if pp.Style != nil {

		if err = e.EncodeElement(pp.Style, xml.StartElement{Name: xml.Name{Local: "w:pStyle"}}); err != nil {
			return err
		}
	}

	// Encoding <w:jc> element
	if pp.Justification != nil {
		if err = e.EncodeElement(pp.Justification, xml.StartElement{Name: xml.Name{Local: "w:jc"}}); err != nil {
			return err
		}
	}

	// Encoding <w:numPr> element
	if pp.NumberingProperty != nil {
		if err = e.EncodeElement(pp.NumberingProperty, xml.StartElement{Name: xml.Name{Local: "w:numPr"}}); err != nil {
			return err
		}
	}

	// Closing </w:pPr> element
	if err = e.EncodeToken(elem.End()); err != nil {
		return err
	}

	return nil
}

func (pp *ParagraphProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name {
			case xml.Name{Space: constants.WMLNamespace, Local: "pStyle"}, xml.Name{Space: constants.AltWMLNamespace, Local: "pStyle"}:
				if err = d.DecodeElement(&pp.Style, &t); err != nil {
					return err
				}
			case xml.Name{Space: constants.WMLNamespace, Local: "jc"}, xml.Name{Space: constants.AltWMLNamespace, Local: "jc"}:
				if err = d.DecodeElement(&pp.Justification, &t); err != nil {
					return err
				}
			case xml.Name{Space: constants.WMLNamespace, Local: "numPr"}, xml.Name{Space: constants.AltWMLNamespace, Local: "numPr"}:
				if err = d.DecodeElement(&pp.NumberingProperty, &t); err != nil {
					return err
				}
			default:
				// fmt.Println(t.Name)
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			// Check if it's the end of the w:pPr element
			if t == start.End() {
				return nil
			}
		}
	}
}
