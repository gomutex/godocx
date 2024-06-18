package docxrun

import (
	"encoding/xml"
	"fmt"

	"github.com/gomutex/godocx/elemtypes"
	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/formatting"
)

// RunProperty represents the properties of a run of text within a paragraph.
type RunProperty struct {
	Fonts     *RunFonts             `xml:"rFonts,omitempty"`
	Color     *formatting.Color     `xml:"color,omitempty"`
	Size      *FontSize             `xml:"sz,omitempty"`
	SizeCs    *FontSizeCS           `xml:"szCs,omitempty"`
	Shading   *ctypes.Shading       `xml:"shd,omitempty"`
	Highlight *formatting.Highlight `xml:"highlight,omitempty"`
	Underline *formatting.Underline `xml:"u,omitempty"`
	Effect    *Effect               `xml:"effect,omitempty"`
	ExpaComp  *ExpaComp             `xml:"w,omitempty"`
	Border    *TextBorder           `xml:"bdr,omitempty"`
	FitText   *FitText              `xml:"fitText,omitempty"`
	VertAlign *VertAlign            `xml:"vertAlign,omitempty"`
	Em        *Em                   `xml:"em,omitempty"`
	Lang      *Lang                 `xml:"lang,omitempty"`
	EALayout  *EALayout             `xml:"eastAsianLayout,omitempty"`

	Bold         *elemtypes.OptBoolElem     `xml:"b,omitempty"`
	BoldCS       *elemtypes.OptBoolElem     `xml:"bCs,omitempty"`
	Italic       *elemtypes.OptBoolElem     `xml:"i,omitempty"`
	ItalicCS     *elemtypes.OptBoolElem     `xml:"iCs,omitempty"`
	Strike       *elemtypes.OptBoolElem     `xml:"strike,omitempty"`
	DoubleStrike *elemtypes.OptBoolElem     `xml:"dstrike,omitempty"`
	Outline      *elemtypes.OptBoolElem     `xml:"outline,omitempty"`
	Shadow       *elemtypes.OptBoolElem     `xml:"shadow,omitempty"`
	Caps         *elemtypes.OptBoolElem     `xml:"caps,omitempty"`
	SmallCaps    *elemtypes.OptBoolElem     `xml:"smallCaps,omitempty"`
	Emboss       *elemtypes.OptBoolElem     `xml:"emboss,omitempty"`
	Imprint      *elemtypes.OptBoolElem     `xml:"imprint,omitempty"`
	NoGrammar    *elemtypes.OptBoolElem     `xml:"noProof,omitempty"`
	SnapToGrid   *elemtypes.OptBoolElem     `xml:"snapToGrid,omitempty"`
	Vanish       *elemtypes.OptBoolElem     `xml:"vanish,omitempty"`
	WebHidden    *elemtypes.OptBoolElem     `xml:"webHidden,omitempty"`
	RightToLeft  *elemtypes.OptBoolElem     `xml:"rtl,omitempty"`
	CSFormat     *elemtypes.OptBoolElem     `xml:"cs,omitempty"`
	SpecVanish   *elemtypes.OptBoolElem     `xml:"specVanish,omitempty"`
	OMath        *elemtypes.OptBoolElem     `xml:"oMath,omitempty"`
	Kern         *elemtypes.SingleUint64Val `xml:"kern,omitempty"`
	Spacing      *elemtypes.SingleIntVal    `xml:"spacing,omitempty"`
	Style        *elemtypes.SingleStrVal    `xml:"rStyle,omitempty"`
	Position     *elemtypes.SingleIntVal    `xml:"position,omitempty"`
}

// NewRunProperty creates a new RunProperty with default values.
func NewRunProperty() RunProperty {
	return RunProperty{}
}

type optBoolElems struct {
	elem    *elemtypes.OptBoolElem
	XMLName string
}

// MarshalXML marshals RunProperty to XML.
func (rp *RunProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:rPr"
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	obElems := []optBoolElems{
		{rp.Bold, "w:b"},
		{rp.BoldCS, "w:bCs"},
		{rp.RightToLeft, "w:rtl"},
		{rp.SpecVanish, "w:specVanish"},
		{rp.OMath, "w:oMath"},
		{rp.CSFormat, "w:cs"},
		{rp.Italic, "w:i"},
		{rp.ItalicCS, "w:iCs"},
		{rp.Strike, "w:strike"},
		{rp.DoubleStrike, "w:dstrike"},
		{rp.Outline, "w:outline"},
		{rp.Shadow, "w:shadow"},
		{rp.Caps, "w:caps"},
		{rp.SmallCaps, "w:smallCaps"},
		{rp.SnapToGrid, "w:snapToGrid"},
		{rp.Emboss, "w:emboss"},
		{rp.Imprint, "w:imprint"},
		{rp.Vanish, "w:vanish"},
		{rp.WebHidden, "w:webHidden"},
		{rp.NoGrammar, "w:noProof"},
	}

	for _, entry := range obElems {
		if entry.elem == nil {
			continue
		}
		if err = entry.elem.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: entry.XMLName},
		}); err != nil {
			return fmt.Errorf("error in marshaling run property `%s`: %w", entry.XMLName, err)
		}
	}

	if rp.Fonts != nil {
		if err = rp.Fonts.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("Fonts: %w", err)
		}
	}

	if rp.EALayout != nil {
		if err = rp.EALayout.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("East Asian Typography Settings: %w", err)
		}
	}

	if rp.Color != nil {
		if err = rp.Color.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("color: %w", err)
		}
	}

	if rp.Em != nil {
		if err = rp.Em.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("emphasis mark: %w", err)
		}
	}

	if rp.Style != nil {
		if err = rp.Style.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:rStyle"},
		}); err != nil {
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
			return fmt.Errorf("size complex script: %w", err)
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

	if rp.FitText != nil {
		if err = rp.FitText.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("fit text: %w", err)
		}
	}

	if rp.VertAlign != nil {
		if err = rp.VertAlign.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("vertical align: %w", err)
		}
	}

	if rp.Border != nil {
		if err = rp.Border.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:bdr"},
		}); err != nil {
			return fmt.Errorf("border: %w", err)
		}
	}

	if rp.Spacing != nil {
		if err = rp.Spacing.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:spacing"},
		}); err != nil {
			return fmt.Errorf("spacing: %w", err)
		}
	}

	if rp.ExpaComp != nil {
		if err = rp.ExpaComp.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:w"},
		}); err != nil {
			return fmt.Errorf("expand/compression text: %w", err)
		}
	}

	if rp.Kern != nil {
		if err = rp.Kern.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:kern"},
		}); err != nil {
			return fmt.Errorf("kern: %w", err)
		}
	}

	if rp.Position != nil {
		if err = rp.Position.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:position"},
		}); err != nil {
			return fmt.Errorf("position: %w", err)
		}
	}

	if rp.Effect != nil {
		if err = rp.Effect.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:effect"},
		}); err != nil {
			return fmt.Errorf("effect: %w", err)
		}
	}

	if rp.Underline != nil {
		if err = rp.Underline.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("underline: %w", err)
		}
	}

	return e.EncodeToken(start.End())
}
