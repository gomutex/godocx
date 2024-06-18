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

	if rp.EALayout != nil {
		if err = rp.EALayout.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("East Asian Typography Settings: %w", err)
		}
	}

	if rp.Bold != nil {
		if err = rp.Bold.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:b"},
		}); err != nil {
			return fmt.Errorf("bold: %w", err)
		}
	}

	if rp.BoldCS != nil {
		if err = rp.Bold.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:bCs"},
		}); err != nil {
			return fmt.Errorf("bold complex script: %w", err)
		}
	}

	if rp.RightToLeft != nil {
		if err = rp.RightToLeft.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:rtl"},
		}); err != nil {
			return fmt.Errorf("right to left: %w", err)
		}
	}

	if rp.SpecVanish != nil {
		if err = rp.SpecVanish.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:specVanish"},
		}); err != nil {
			return fmt.Errorf("specVanish: %w", err)
		}
	}

	if rp.OMath != nil {
		if err = rp.OMath.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:oMath"},
		}); err != nil {
			return fmt.Errorf("Office Open XML Math: %w", err)
		}
	}

	if rp.CSFormat != nil {
		if err = rp.CSFormat.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:cs"},
		}); err != nil {
			return fmt.Errorf("complex script formatting: %w", err)
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

	if rp.Italic != nil {
		if err = rp.Italic.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:i"},
		}); err != nil {
			return fmt.Errorf("italic: %w", err)
		}
	}

	if rp.ItalicCS != nil {
		if err = rp.ItalicCS.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:iCs"},
		}); err != nil {
			return fmt.Errorf("Italic complex script: %w", err)
		}
	}

	if rp.Strike != nil {
		if err = rp.Strike.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:strike"},
		}); err != nil {
			return fmt.Errorf("strike: %w", err)
		}
	}

	if rp.Border != nil {
		if err = rp.Border.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:bdr"},
		}); err != nil {
			return fmt.Errorf("border: %w", err)
		}
	}

	if rp.DoubleStrike != nil {
		if err = rp.DoubleStrike.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:dstrike"},
		}); err != nil {
			return fmt.Errorf("double strike: %w", err)
		}
	}

	if rp.Outline != nil {
		if err = rp.Outline.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:outline"},
		}); err != nil {
			return fmt.Errorf("outline: %w", err)
		}
	}

	if rp.Shadow != nil {
		if err = rp.Shadow.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:shadow"},
		}); err != nil {
			return fmt.Errorf("shadow: %w", err)
		}
	}

	if rp.Caps != nil {
		if err = rp.Caps.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:caps"},
		}); err != nil {
			return fmt.Errorf("caps: %w", err)
		}
	}

	if rp.SmallCaps != nil {
		if err = rp.SmallCaps.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:smallCaps"},
		}); err != nil {
			return fmt.Errorf("small caps: %w", err)
		}
	}

	if rp.Emboss != nil {
		if err = rp.Emboss.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:emboss"},
		}); err != nil {
			return fmt.Errorf("emboss: %w", err)
		}
	}

	if rp.SnapToGrid != nil {
		if err = rp.SnapToGrid.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:snapToGrid"},
		}); err != nil {
			return fmt.Errorf("snap to grid: %w", err)
		}
	}

	if rp.Imprint != nil {
		if err = rp.Imprint.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:imprint"},
		}); err != nil {
			return fmt.Errorf("imprint: %w", err)
		}
	}

	if rp.Vanish != nil {
		if err = rp.Vanish.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:vanish"},
		}); err != nil {
			return fmt.Errorf("vanish: %w", err)
		}
	}

	if rp.WebHidden != nil {
		if err = rp.WebHidden.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:webHidden"},
		}); err != nil {
			return fmt.Errorf("web hidden text: %w", err)
		}
	}

	if rp.NoGrammar != nil {
		if err = rp.NoGrammar.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:noProof"},
		}); err != nil {
			return fmt.Errorf("no Grammar: %w", err)
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
