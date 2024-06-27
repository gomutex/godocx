package ctypes

import (
	"encoding/xml"
	"fmt"
)

type ParaBorder struct {
	Top     *Border `xml:"top,omitempty"`
	Left    *Border `xml:"left,omitempty"`
	Right   *Border `xml:"right,omitempty"`
	Bottom  *Border `xml:"bottom,omitempty"`
	Between *Border `xml:"between,omitempty"`
	Bar     *Border `xml:"bar,omitempty"`
}

func (p ParaBorder) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:pBdr"
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if p.Top != nil {
		if err = p.Top.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:top"}}); err != nil {
			return fmt.Errorf("Paragraph border-Top: %w", err)
		}
	}

	if p.Left != nil {
		if err = p.Left.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:left"}}); err != nil {
			return fmt.Errorf("Paragraph border-Left: %w", err)
		}
	}

	if p.Right != nil {
		if err = p.Right.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:right"}}); err != nil {
			return fmt.Errorf("Paragraph border-Right: %w", err)
		}
	}

	if p.Bottom != nil {
		if err = p.Bottom.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:bottom"}}); err != nil {
			return fmt.Errorf("Paragraph border-Bottom: %w", err)
		}
	}

	if p.Between != nil {
		if err = p.Between.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:between"}}); err != nil {
			return fmt.Errorf("Paragraph border-Between: %w", err)
		}
	}

	if p.Bar != nil {
		if err = p.Bar.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:bar"}}); err != nil {
			return fmt.Errorf("Paragraph border-Bar: %w", err)
		}
	}

	return e.EncodeToken(start.End())
}
