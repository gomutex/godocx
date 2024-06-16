package docxrun

import (
	"encoding/xml"
	"fmt"

	"github.com/gomutex/godocx/elemtypes"
	"github.com/gomutex/godocx/wml/formatting"
)

// RunProperty represents the properties of a run of text within a paragraph.
type RunProperty struct {
	Fonts  *RunFonts               `xml:"rFonts,omitempty"`
	Bold   *elemtypes.NullBoolElem `xml:"b,omitempty"`
	Color  *formatting.Color       `xml:"color,omitempty"`
	Size   *Sz                     `xml:"sz,omitempty"`
	SizeCs *SzCs                   `xml:"szCs,omitempty"`
	Style  *RunStyle               `xml:"rStyle,omitempty"`

	Shading   *formatting.Shading   `xml:"shd,omitempty"`
	Highlight *formatting.Highlight `xml:"highlight,omitempty"`
	Italic    *formatting.Italic    `xml:"i,omitempty"`

	Strike    *formatting.Strike    `xml:"strike,omitempty"`
	Underline *formatting.Underline `xml:"u,omitempty"`
	// Vanish           *Vanish
	// SpecVanish       *SpecVanish
	// VertAlign        *VertAlign
	// CharacterSpacing *CharacterSpacing
	// TextBorder       *TextBorder
	// Del              *Delete
	// Ins              *Insert
}

// NewRunProperty creates a new RunProperty with default values.
func NewRunProperty() RunProperty {
	return RunProperty{}
}

// MarshalXML marshals RunProperty to XML.
func (rp *RunProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:rPr"
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if rp.Fonts != nil {
		if err = rp.Fonts.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("Fonts: %w", err)
		}
	}

	if rp.Bold != nil {
		if err = rp.Bold.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:b"},
		}); err != nil {
			return fmt.Errorf("bold: %w", err)
		}
	}

	if rp.Color != nil {
		if err = rp.Color.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("color: %w", err)
		}
	}

	if rp.Style != nil {
		if err = rp.Style.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("style: %w", err)
		}
	}

	if rp.Size != nil {
		if err = rp.Size.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("size: %w", err)
		}
	}

	if rp.SizeCs != nil {
		if err = rp.SizeCs.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("sizeCs: %w", err)
		}
	}

	if rp.Shading != nil {
		if err = rp.Shading.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("shading: %w", err)
		}
	}

	if rp.Highlight != nil {
		if err = rp.Highlight.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("highlight: %w", err)
		}
	}

	if rp.Italic != nil {
		if err = rp.Italic.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("italic: %w", err)
		}
	}

	if rp.Strike != nil {
		if err = rp.Strike.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("strike: %w", err)
		}
	}

	if rp.Underline != nil {
		if err = rp.Underline.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("underline: %w", err)
		}
	}

	return e.EncodeToken(start.End())
}
