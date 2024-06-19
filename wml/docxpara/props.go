package docxpara

import (
	"encoding/xml"
	"fmt"

	"github.com/gomutex/godocx/elemtypes"
	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/docxrun"
	"github.com/gomutex/godocx/wml/formatting"
	"github.com/gomutex/godocx/wml/sections"
)

// Numbering Level Associated Paragraph Properties
type ParagraphProp struct {
	// This element specifies the style ID of the paragraph style which shall be used to format the contents of this paragraph.
	Style *elemtypes.SingleStrVal `xml:"pStyle,omitempty"`

	//Keep Paragraph With Next Paragraph
	KeepNext *elemtypes.OptBinFlagElem `xml:"keepNext,omitempty"`

	//Keep All Lines On One Page
	KeepLines *elemtypes.OptBinFlagElem `xml:"keepLines,omitempty"`

	//Start Paragraph on Next Page
	PageBreakBefore *elemtypes.OptBinFlagElem `xml:"pageBreakBefore,omitempty"`

	//Allow First/Last Line to Display on a Separate Page
	WindowControl *elemtypes.OptBinFlagElem `xml:"widowControl,omitempty"`

	//Suppress Line Numbers for Paragraph
	SuppressLineNmbrs *elemtypes.OptBinFlagElem `xml:"suppressLineNumbers,omitempty"`

	//Suppress Hyphenation for Paragraph
	SuppressAutoHyphens *elemtypes.OptBinFlagElem `xml:"suppressAutoHyphens,omitempty"`

	//Use East Asian Typography Rules for First and Last Character per Line
	Kinsoku *elemtypes.OptBinFlagElem `xml:"kinsoku,omitempty"`

	//Allow Line Breaking At Character Level
	WordWrap *elemtypes.OptBinFlagElem `xml:"wordWrap,omitempty"`

	//Allow Punctuation to Extent Past Text Extents
	OverflowPunct *elemtypes.OptBinFlagElem `xml:"overflowPunct,omitempty"`

	//Compress Punctuation at Start of a Line
	TopLinePunct *elemtypes.OptBinFlagElem `xml:"topLinePunct,omitempty"`

	//Automatically Adjust Spacing of Latin and East Asian Text
	AutoSpaceDE *elemtypes.OptBinFlagElem `xml:"autoSpaceDE,omitempty"`

	//Automatically Adjust Spacing of East Asian Text and Numbers
	AutoSpaceDN *elemtypes.OptBinFlagElem `xml:"autoSpaceDN,omitempty"`

	//Right to Left Paragraph Layout
	Bidi *elemtypes.OptBinFlagElem `xml:"bidi,omitempty"`

	//Automatically Adjust Right Indent When Using Document Grid
	AdjustRightInd *elemtypes.OptBinFlagElem `xml:"adjustRightInd,omitempty"`

	//Use Document Grid Settings for Inter-Line Paragraph Spacing
	SnapToGrid *elemtypes.OptBinFlagElem `xml:"snapToGrid,omitempty"`

	//Ignore Spacing Above and Below When Using Identical Styles
	CtxlSpacing *elemtypes.OptBinFlagElem `xml:"contextualSpacing,omitempty"`

	// Use Left/Right Indents as Inside/Outside Indents
	MirrorIndents *elemtypes.OptBinFlagElem `xml:"mirrorIndents,omitempty"`

	// Prevent Text Frames From Overlapping
	SuppressOverlap *elemtypes.OptBinFlagElem `xml:"suppressOverlap,omitempty"`

	//This element specifies the shading applied to the contents of the paragraph.
	Shading *ctypes.Shading `xml:"shd,omitempty"`

	//Paragraph Text Flow Direction
	TextDirection *ctypes.TextDirection `xml:"textDirection,omitempty"`

	//Vertical Character Alignment on Line
	TextAlignment *ctypes.TextAlign `xml:"textAlignment,omitempty"`

	//Allow Surrounding Paragraphs to Tight Wrap to Text Box Contents
	TextboxTightWrap *ctypes.TextboxTightWrap `xml:"textboxTightWrap,omitempty"`

	//Set of Custom Tab Stops
	Tabs ctypes.Tabs `xml:"tabs,omitempty"`

	//Associated Outline Level
	OutlineLvl *ctypes.DecimalNum `xml:"outlineLvl,omitempty"`

	//Associated HTML div ID
	DivID *ctypes.DecimalNum `xml:"divId,omitempty"`

	//Paragraph Conditional Formatting
	CnfStyle *ctypes.Cnf `xml:"cnfStyle,omitempty"`

	//Run Properties for the Paragraph Mark
	RunProperty *docxrun.RunProperty `xml:"rPr,omitempty"`

	//Paragraph Alignment
	Justification *formatting.Justification `xml:"jc,omitempty"`

	//Revision Information for Paragraph Properties
	PPrChange *PPrChange `xml:"pPrChange,omitempty"`

	// Text Frame Properties
	FrameProp *FrameProp `xml:"framePr,omitempty"`

	//Numbering Definition Instance Reference
	NumProp *NumProp `xml:"numPr,omitempty"`

	//Paragraph Borders
	Border *ParaBorder `xml:"pBdr,omitempty"`

	//Spacing Between Lines and Above/Below Paragraph
	Spacing *Spacing `xml:"spacing,omitempty"`

	Indent *Indent              `xml:"ind,omitempty"`
	SectPr sections.SectionProp `xml:"sectPr,omitempty"`
}

type binElems struct {
	elem    *elemtypes.OptBinFlagElem
	XMLName string
}

func (pp *ParagraphProp) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	elem := xml.StartElement{Name: xml.Name{Local: "w:pPr"}}

	// Opening <w:pPr> element
	if err = e.EncodeToken(elem); err != nil {
		return err
	}

	bElems := []binElems{
		{pp.KeepNext, "w:keepNext"},
		{pp.KeepLines, "w:keepLines"},
		{pp.KeepLines, "w:keepLines"},
		{pp.PageBreakBefore, "w:pageBreakBefore"},
		{pp.WindowControl, "w:widowControl"},
		{pp.SuppressLineNmbrs, "w:suppressLineNumbers"},
		{pp.SuppressAutoHyphens, "w:suppressAutoHyphens"},
		{pp.Kinsoku, "w:kinsoku"},
		{pp.OverflowPunct, "w:overflowPunct"},
		{pp.TopLinePunct, "w:topLinePunct"},
		{pp.AutoSpaceDE, "w:autoSpaceDE"},
		{pp.AutoSpaceDE, "w:autoSpaceDE"},
		{pp.AutoSpaceDN, "w:autoSpaceDN"},
		{pp.Bidi, "w:bidi"},
		{pp.AdjustRightInd, "w:adjustRightInd"},
		{pp.SnapToGrid, "w:snapToGrid"},
		{pp.CtxlSpacing, "w:contextualSpacing"},
		{pp.MirrorIndents, "w:mirrorIndents"},
		{pp.WordWrap, "w:wordWrap"},
		{pp.SuppressOverlap, "w:suppressOverlap"},
	}

	for _, entry := range bElems {
		if entry.elem == nil {
			continue
		}
		if err = entry.elem.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: entry.XMLName},
		}); err != nil {
			return fmt.Errorf("error in marshaling paragraph property `%s`: %w", entry.XMLName, err)
		}
	}

	if pp.Style != nil {
		if err = pp.Style.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:pStyle"},
		}); err != nil {
			return fmt.Errorf("style: %w", err)
		}
	}

	if pp.RunProperty != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:rPr"}}
		if err = e.EncodeElement(pp.RunProperty, propsElement); err != nil {
			return err
		}
	}

	if pp.Justification != nil {
		if err = pp.Justification.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("Justification: %w", err)
		}
	}

	if pp.TextDirection != nil {
		if err = pp.TextDirection.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("TextDirection: %w", err)
		}
	}

	if pp.TextAlignment != nil {
		if err = pp.TextAlignment.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("TextAlignment: %w", err)
		}
	}

	if pp.TextboxTightWrap != nil {
		if err = pp.TextboxTightWrap.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("TextboxTightWrap: %w", err)
		}
	}

	if pp.OutlineLvl != nil {
		if err = pp.OutlineLvl.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:outlineLvl"},
		}); err != nil {
			return fmt.Errorf("OutlineLvl: %w", err)
		}
	}

	if pp.DivID != nil {
		if err = pp.DivID.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("DivID: %w", err)
		}
	}

	if pp.CnfStyle != nil {
		if err = pp.CnfStyle.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("CnfStyle: %w", err)
		}
	}

	if pp.NumProp != nil {
		if err = pp.NumProp.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("NumberingProperty: %w", err)
		}
	}

	if pp.PPrChange != nil {
		if err = pp.PPrChange.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("PPrChange: %w", err)
		}
	}

	if pp.FrameProp != nil {
		if err = pp.FrameProp.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("FrameProp: %w", err)
		}
	}

	if pp.Border != nil {
		if err = pp.Border.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("Border: %w", err)
		}
	}

	if pp.Spacing != nil {
		if err = pp.Spacing.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("Spacing: %w", err)
		}
	}

	if pp.Indent != nil {
		if err = pp.Indent.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("Indent: %w", err)
		}
	}

	return e.EncodeToken(elem.End())
}

// NewParagraphStyle creates a new ParagraphStyle.
func NewParagraphStyle(val string) *elemtypes.SingleStrVal {
	return &elemtypes.SingleStrVal{Val: val}
}

// DefaultParagraphStyle creates the default ParagraphStyle with the value "Normal".
func DefaultParagraphStyle() *elemtypes.SingleStrVal {
	return &elemtypes.SingleStrVal{Val: "Normal"}
}

func DefaultParaProperty() *ParagraphProp {
	return &ParagraphProp{}
}
