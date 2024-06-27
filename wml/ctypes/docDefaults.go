package ctypes

import (
	"encoding/xml"
	"fmt"
)

// Document Default Paragraph and Run Properties
type DocDefault struct {
	//Sequence

	//1.Default Run Properties
	RunProp *RunPropDefault `xml:"rPrDefault,omitempty"`

	//2.Default Paragraph Properties
	ParaProp *ParaPropDefault `xml:"pPrDefault,omitempty"`
}

type RunPropDefault struct {
	RunProp *RunProperty `xml:"rPr,omitempty"`
}

type ParaPropDefault struct {
	ParaProp *ParagraphProp `xml:"pPr,omitempty"`
}

func (d *DocDefault) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:docDefaults"

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if d.RunProp != nil {
		if err := d.RunProp.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:rPrDefault"}}); err != nil {
			return fmt.Errorf("DocDefault RunProp: %w", err)
		}
	}

	if d.ParaProp != nil {
		if err := d.ParaProp.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:pPrDefault"}}); err != nil {
			return fmt.Errorf("DocDefault ParaProp: %w", err)
		}
	}

	return e.EncodeToken(start.End())
}

func (r *RunPropDefault) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:rPrDefault"
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if r.RunProp != nil {
		if err := e.EncodeElement(r.RunProp, xml.StartElement{
			Name: xml.Name{Local: "w:rPr"},
		}); err != nil {
			return fmt.Errorf("RunPropDefault: %w", err)
		}
	}

	return e.EncodeToken(start.End())
}

func (p *ParaPropDefault) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:pPrDefault"
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if p.ParaProp != nil {
		if err := e.EncodeElement(p.ParaProp, xml.StartElement{
			Name: xml.Name{Local: "w:pPr"},
		}); err != nil {
			return fmt.Errorf("ParaPropDefault: %w", err)
		}
	}

	return e.EncodeToken(start.End())
}
