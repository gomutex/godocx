package docxpara

import (
	"encoding/xml"
	"fmt"

	"github.com/gomutex/godocx/elemtypes"
	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/docxrun"
	"github.com/gomutex/godocx/wml/formatting"
	"github.com/gomutex/godocx/wml/liststyle"
)

// Numbering Level Associated Paragraph Properties
type ParagraphProperty struct {
	// This element specifies the style ID of the paragraph style which shall be used to format the contents of this paragraph.
	Style *elemtypes.SingleStrVal `xml:"pStyle,omitempty"`

	//Keep Paragraph With Next Paragraph
	KeepNext *elemtypes.OptOnOffElem `xml:"keepNext,omitempty"`

	//Keep All Lines On One Page
	KeepLines *elemtypes.OptOnOffElem `xml:"keepLines,omitempty"`

	//Start Paragraph on Next Page
	PageBreakBefore *elemtypes.OptOnOffElem `xml:"pageBreakBefore,omitempty"`

	//Allow First/Last Line to Display on a Separate Page
	WindowControl *elemtypes.OptOnOffElem `xml:"widowControl,omitempty"`

	//Suppress Line Numbers for Paragraph
	SuppressLineNmbrs *elemtypes.OptOnOffElem `xml:"suppressLineNumbers,omitempty"`

	//Suppress Hyphenation for Paragraph
	SuppressAutoHyphens *elemtypes.OptOnOffElem `xml:"suppressAutoHyphens,omitempty"`

	//Use East Asian Typography Rules for First and Last Character per Line
	Kinsoku *elemtypes.OptOnOffElem `xml:"kinsoku,omitempty"`

	//Allow Line Breaking At Character Level
	WordWrap *elemtypes.OptOnOffElem `xml:"wordWrap,omitempty"`

	//Allow Punctuation to Extent Past Text Extents
	OverflowPunct *elemtypes.OptOnOffElem `xml:"overflowPunct,omitempty"`

	//Compress Punctuation at Start of a Line
	TopLinePunct *elemtypes.OptOnOffElem `xml:"topLinePunct,omitempty"`

	//Automatically Adjust Spacing of Latin and East Asian Text
	AutoSpaceDE *elemtypes.OptOnOffElem `xml:"autoSpaceDE,omitempty"`

	//Automatically Adjust Spacing of East Asian Text and Numbers
	AutoSpaceDN *elemtypes.OptOnOffElem `xml:"autoSpaceDN,omitempty"`

	//Right to Left Paragraph Layout
	Bidi *elemtypes.OptOnOffElem `xml:"bidi,omitempty"`

	//Automatically Adjust Right Indent When Using Document Grid
	AdjustRightInd *elemtypes.OptOnOffElem `xml:"adjustRightInd,omitempty"`

	//Use Document Grid Settings for Inter-Line Paragraph Spacing
	SnapToGrid *elemtypes.OptOnOffElem `xml:"snapToGrid,omitempty"`

	//Ignore Spacing Above and Below When Using Identical Styles
	CtxlSpacing *elemtypes.OptOnOffElem `xml:"contextualSpacing,omitempty"`

	// Use Left/Right Indents as Inside/Outside Indents
	MirrorIndents *elemtypes.OptOnOffElem `xml:"mirrorIndents,omitempty"`

	//This element specifies the shading applied to the contents of the paragraph.
	Shading *ctypes.Shading `xml:"shd,omitempty"`

	//Set of Custom Tab Stops
	Tabs ctypes.Tabs `xml:"tabs,omitempty"`

	DivID             *string
	Justification     *formatting.Justification    `xml:"jc,omitempty"`
	RunProperty       *docxrun.RunProperty         `xml:"rPr,omitempty"`
	NumberingProperty *liststyle.NumberingProperty `xml:"numPr,omitempty"`

	// TODO: Implement this
	// Text Frame Properties
	// FrameProp *FrameProp `xml:"framePr,omitempty"`
	// This element specifies that the current paragraph references a numbering definition instance in the current document.
	// NumPr *NumPr `xml:"numpr,omitempty"`
	//This element specifies the borders for the parent paragraph. Each child element shall specify a specific type of border (left, right, bottom, top, and between).
	// Border *Border `xml:"pBdr,omitempty"`
	// Spacing Spacing
	//Indent Indent
}

type onOffElems struct {
	elem    *elemtypes.OptOnOffElem
	XMLName string
}

func (pp *ParagraphProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	elem := xml.StartElement{Name: xml.Name{Local: "w:pPr"}}

	// Opening <w:pPr> element
	if err = e.EncodeToken(elem); err != nil {
		return err
	}

	ooElems := []onOffElems{
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
	}

	for _, entry := range ooElems {
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

	if pp.NumberingProperty != nil {
		if err = pp.NumberingProperty.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("NumberingProperty: %w", err)
		}
	}

	if err = e.EncodeToken(elem.End()); err != nil {
		return err
	}

	return nil
}

// NewParagraphStyle creates a new ParagraphStyle.
func NewParagraphStyle(val string) *elemtypes.SingleStrVal {
	return &elemtypes.SingleStrVal{Val: val}
}

// DefaultParagraphStyle creates the default ParagraphStyle with the value "Normal".
func DefaultParagraphStyle() *elemtypes.SingleStrVal {
	return &elemtypes.SingleStrVal{Val: "Normal"}
}

func DefaultParaProperty() *ParagraphProperty {
	return &ParagraphProperty{}
}
