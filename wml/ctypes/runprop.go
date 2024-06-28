package ctypes

import (
	"encoding/xml"
	"fmt"

	"github.com/gomutex/godocx/wml/stypes"
)

// RunProperty represents the properties of a run of text within a paragraph.
type RunProperty struct {
	//1. Referenced Character Style
	Style *CTString `xml:"rStyle,omitempty"`

	//2. Run Fonts
	Fonts *RunFonts `xml:"rFonts,omitempty"`

	//3. Bold
	Bold *OnOff `xml:"b,omitempty"`

	//4.Complex Script Bold
	BoldCS *OnOff `xml:"bCs,omitempty"`

	// 5.Italics
	Italic *OnOff `xml:"i,omitempty"`

	//6.Complex Script Italics
	ItalicCS *OnOff `xml:"iCs,omitempty"`

	//7.Display All Characters As Capital Letters
	Caps *OnOff `xml:"caps,omitempty"`

	//8.Small Caps
	SmallCaps *OnOff `xml:"smallCaps,omitempty"`

	//9.Single Strikethrough
	Strike *OnOff `xml:"strike,omitempty"`

	//10.Double Strikethrough
	DoubleStrike *OnOff `xml:"dstrike,omitempty"`

	//11.Display Character Outline
	Outline *OnOff `xml:"outline,omitempty"`

	//12.Shadow
	Shadow *OnOff `xml:"shadow,omitempty"`

	//13.Embossing
	Emboss *OnOff `xml:"emboss,omitempty"`

	//14.Imprinting
	Imprint *OnOff `xml:"imprint,omitempty"`

	//15.Do Not Check Spelling or Grammar
	NoGrammar *OnOff `xml:"noProof,omitempty"`

	//16.Use Document Grid Settings For Inter-Character Spacing
	SnapToGrid *OnOff `xml:"snapToGrid,omitempty"`

	//17.Hidden Text
	Vanish *OnOff `xml:"vanish,omitempty"`

	//18.Web Hidden Text
	WebHidden *OnOff `xml:"webHidden,omitempty"`

	//19.Run Content Color
	Color *Color `xml:"color,omitempty"`

	//20. Character Spacing Adjustment
	Spacing *DecimalNum `xml:"spacing,omitempty"`

	//21.Expanded/Compressed Text
	ExpaComp *ExpaComp `xml:"w,omitempty"`

	//22.Font Kerning
	Kern *Uint64Elem `xml:"kern,omitempty"`

	//23. Vertically Raised or Lowered Text
	Position *DecimalNum `xml:"position,omitempty"`

	//24.Font Size
	Size *FontSize `xml:"sz,omitempty"`

	//25.Complex Script Font Size
	SizeCs *FontSizeCS `xml:"szCs,omitempty"`

	//26.Text Highlighting
	Highlight *CTString `xml:"highlight,omitempty"`

	//27.Underline
	Underline *GenSingleStrVal[stypes.Underline] `xml:"u,omitempty"`

	//28.Animated Text Effect
	Effect *Effect `xml:"effect,omitempty"`

	//29.Text Border
	Border *Border `xml:"bdr,omitempty"`

	//30.Run Shading
	Shading *Shading `xml:"shd,omitempty"`

	//31.Manual Run Width
	FitText *FitText `xml:"fitText,omitempty"`

	//32.Subscript/Superscript Text
	VertAlign *GenSingleStrVal[stypes.VerticalAlignRun] `xml:"vertAlign,omitempty"`

	//33.Right To Left Text
	RightToLeft *OnOff `xml:"rtl,omitempty"`

	//34.Use Complex Script Formatting on Run
	CSFormat *OnOff `xml:"cs,omitempty"`

	//35.Emphasis Mark
	Em *GenSingleStrVal[stypes.Em] `xml:"em,omitempty"`

	//36.Languages for Run Content
	Lang *Lang `xml:"lang,omitempty"`

	//37.East Asian Typography Settings
	EALayout *EALayout `xml:"eastAsianLayout,omitempty"`

	//38.Paragraph Mark Is Always Hidden
	SpecVanish *OnOff `xml:"specVanish,omitempty"`

	//39.Office Open XML Math
	OMath *OnOff `xml:"oMath,omitempty"`
}

// NewRunProperty creates a new RunProperty with default values.
func NewRunProperty() RunProperty {
	return RunProperty{}
}

type optBoolElems struct {
	elem    *OnOff
	XMLName string
}

// MarshalXML marshals RunProperty to XML.
func (rp RunProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:rPr"
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	// 1. Referenced Character Style
	if rp.Style != nil {
		if err = rp.Style.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:rStyle"},
		}); err != nil {
			return fmt.Errorf("style: %w", err)
		}
	}

	//2.Run Fonts
	if rp.Fonts != nil {
		if err = rp.Fonts.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("Fonts: %w", err)
		}
	}

	set1 := []optBoolElems{
		{rp.Bold, "w:b"},                //3.Bold
		{rp.BoldCS, "w:bCs"},            //4.Complex Script Bold
		{rp.Italic, "w:i"},              //5.Italics
		{rp.ItalicCS, "w:iCs"},          //6.Complex Script Italics
		{rp.Caps, "w:caps"},             //7.Display All Characters As Capital Letters
		{rp.SmallCaps, "w:smallCaps"},   //8.Small Caps
		{rp.Strike, "w:strike"},         //9.Single Strikethrough
		{rp.DoubleStrike, "w:dstrike"},  //10.Double Strikethrough
		{rp.Outline, "w:outline"},       //11.Display Character Outline
		{rp.Shadow, "w:shadow"},         //12.Shadow
		{rp.Emboss, "w:emboss"},         //13.Embossing
		{rp.Imprint, "w:imprint"},       //14.Imprinting
		{rp.NoGrammar, "w:noProof"},     //15.Do Not Check Spelling or Grammar
		{rp.SnapToGrid, "w:snapToGrid"}, //16.Use Document Grid Settings For Inter-Character Spacing
		{rp.Vanish, "w:vanish"},         //17.Hidden Text
		{rp.WebHidden, "w:webHidden"},   //18.Web Hidden Text
	}

	for _, entry := range set1 {
		if entry.elem == nil {
			continue
		}
		if err = entry.elem.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: entry.XMLName},
		}); err != nil {
			return fmt.Errorf("error in marshaling run property `%s`: %w", entry.XMLName, err)
		}
	}

	//19.Run Content Color
	if rp.Color != nil {
		if err = rp.Color.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("color: %w", err)
		}
	}

	//20. Character Spacing Adjustment
	if rp.Spacing != nil {
		if err = rp.Spacing.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:spacing"},
		}); err != nil {
			return fmt.Errorf("spacing: %w", err)
		}
	}

	//21.Expanded/Compressed Text
	if rp.ExpaComp != nil {
		if err = rp.ExpaComp.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:w"},
		}); err != nil {
			return fmt.Errorf("expand/compression text: %w", err)
		}
	}

	//22.Font Kerning
	if rp.Kern != nil {
		if err = rp.Kern.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:kern"},
		}); err != nil {
			return fmt.Errorf("kern: %w", err)
		}
	}

	//23. Vertically Raised or Lowered Text
	if rp.Position != nil {
		if err = rp.Position.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:position"},
		}); err != nil {
			return fmt.Errorf("position: %w", err)
		}
	}

	//24.Font Size
	if rp.Size != nil {
		if err = rp.Size.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("size: %w", err)
		}
	}

	//25.Complex Script Font Size
	if rp.SizeCs != nil {
		if err = rp.SizeCs.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("size complex script: %w", err)
		}
	}

	//26.Text Highlighting
	if rp.Highlight != nil {
		if err = rp.Highlight.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:highlight"},
		}); err != nil {
			return fmt.Errorf("highlight: %w", err)
		}
	}

	//27.Underline
	if rp.Underline != nil {
		if err = rp.Underline.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:u"},
		}); err != nil {
			return fmt.Errorf("underline: %w", err)
		}
	}

	//28.Animated Text Effect
	if rp.Effect != nil {
		if err = rp.Effect.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:effect"},
		}); err != nil {
			return fmt.Errorf("effect: %w", err)
		}
	}

	//29.Text Border
	if rp.Border != nil {
		if err = rp.Border.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:bdr"},
		}); err != nil {
			return fmt.Errorf("border: %w", err)
		}
	}

	//30.Run Shading
	if rp.Shading != nil {
		if err = rp.Shading.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("shading: %w", err)
		}
	}

	//31.Manual Run Width
	if rp.FitText != nil {
		if err = rp.FitText.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("fit text: %w", err)
		}
	}

	//32.Subscript/Superscript Text
	if rp.VertAlign != nil {
		if err = rp.VertAlign.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:vertAlign"},
		}); err != nil {
			return fmt.Errorf("vertical align: %w", err)
		}
	}

	//33.Right To Left Text
	if rp.RightToLeft != nil {
		if err = rp.RightToLeft.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:rtl"},
		}); err != nil {
			return fmt.Errorf("error in marshaling run property `%s`: %w", "RightToLeft", err)
		}
	}

	//34.Use Complex Script Formatting on Run
	if rp.CSFormat != nil {
		if err = rp.CSFormat.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:cs"},
		}); err != nil {
			return fmt.Errorf("error in marshaling run property `%s`: %w", "CSFormat", err)
		}
	}

	//35.Emphasis Mark
	if rp.Em != nil {
		if err = rp.Em.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:em"},
		}); err != nil {
			return fmt.Errorf("emphasis mark: %w", err)
		}
	}

	//36.Languages for Run Content
	if rp.Lang != nil {
		if err = rp.Lang.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("languages for Run Content: %w", err)
		}
	}

	//37.East Asian Typography Settings
	if rp.EALayout != nil {
		if err = rp.EALayout.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("East Asian Typography Settings: %w", err)
		}
	}

	//38.Paragraph Mark Is Always Hidden
	if rp.SpecVanish != nil {
		if err = rp.SpecVanish.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:specVanish"},
		}); err != nil {
			return fmt.Errorf("error in marshaling run property `%s`: %w", "specVanish", err)
		}
	}

	//39.Office Open XML Math
	if rp.OMath != nil {
		if err = rp.OMath.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:oMath"},
		}); err != nil {
			return fmt.Errorf("error in marshaling run property `%s`: %w", "oMath", err)
		}
	}

	return e.EncodeToken(start.End())
}
