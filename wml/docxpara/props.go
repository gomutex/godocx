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

type OnOffElems struct {
	elem    *elemtypes.OptOnOffElem
	XMLName string
}

func (pp *ParagraphProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	elem := xml.StartElement{Name: xml.Name{Local: "w:pPr"}}

	// Opening <w:pPr> element
	if err = e.EncodeToken(elem); err != nil {
		return err
	}

	// onOffElems := []OnOffElems{
	// 	{pp.KeepNext, "w:keepNext"},
	// }

	if pp.Style != nil {
		if err = pp.Style.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:pStyle"},
		}); err != nil {
			return fmt.Errorf("style: %w", err)
		}
	}

	if pp.KeepLines != nil {
		if err = pp.KeepLines.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:keepLines"},
		}); err != nil {
			return fmt.Errorf("keepLines: %w", err)
		}
	}

	if pp.KeepNext != nil {
		if err = pp.KeepNext.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:keepNext"},
		}); err != nil {
			return fmt.Errorf("keep next: %w", err)
		}
	}

	if pp.PageBreakBefore != nil {
		if err = pp.PageBreakBefore.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:pageBreakBefore"},
		}); err != nil {
			return fmt.Errorf("pageBreakBefore: %w", err)
		}
	}

	if pp.WindowControl != nil {
		if err = pp.WindowControl.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:widowControl"},
		}); err != nil {
			return fmt.Errorf("widowControl: %w", err)
		}
	}

	if pp.SuppressLineNmbrs != nil {
		if err = pp.SuppressLineNmbrs.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:suppressLineNumbers"},
		}); err != nil {
			return fmt.Errorf("suppressLineNumbers: %w", err)
		}
	}

	if pp.SuppressAutoHyphens != nil {
		if err = pp.SuppressAutoHyphens.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:suppressAutoHyphens"},
		}); err != nil {
			return fmt.Errorf("suppressAutoHyphens: %w", err)
		}
	}

	if pp.Kinsoku != nil {
		if err = pp.Kinsoku.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:kinsoku"},
		}); err != nil {
			return fmt.Errorf("kinsoku: %w", err)
		}
	}

	if pp.OverflowPunct != nil {
		if err = pp.OverflowPunct.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:overflowPunct"},
		}); err != nil {
			return fmt.Errorf("overflowPunct: %w", err)
		}
	}

	if pp.TopLinePunct != nil {
		if err = pp.TopLinePunct.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:topLinePunct"},
		}); err != nil {
			return fmt.Errorf("topLinePunct: %w", err)
		}
	}

	if pp.AutoSpaceDE != nil {
		if err = pp.AutoSpaceDE.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:autoSpaceDE"},
		}); err != nil {
			return fmt.Errorf("autoSpaceDE: %w", err)
		}
	}

	if pp.AutoSpaceDN != nil {
		if err = pp.AutoSpaceDN.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:autoSpaceDN"},
		}); err != nil {
			return fmt.Errorf("autoSpaceDN: %w", err)
		}
	}

	if pp.Bidi != nil {
		if err = pp.Bidi.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:bidi"},
		}); err != nil {
			return fmt.Errorf("bidi: %w", err)
		}
	}

	if pp.AdjustRightInd != nil {
		if err = pp.AdjustRightInd.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:adjustRightInd"},
		}); err != nil {
			return fmt.Errorf("adjustRightInd: %w", err)
		}
	}

	if pp.SnapToGrid != nil {
		if err = pp.SnapToGrid.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:snapToGrid"},
		}); err != nil {
			return fmt.Errorf("snapToGrid: %w", err)
		}
	}

	if pp.CtxlSpacing != nil {
		if err = pp.CtxlSpacing.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:contextualSpacing"},
		}); err != nil {
			return fmt.Errorf("contextualSpacing: %w", err)
		}
	}

	if pp.MirrorIndents != nil {
		if err = pp.MirrorIndents.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:mirrorIndents"},
		}); err != nil {
			return fmt.Errorf("mirrorIndents: %w", err)
		}
	}

	if pp.WordWrap != nil {
		if err = pp.WordWrap.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:wordWrap"},
		}); err != nil {
			return fmt.Errorf("wordWrap: %w", err)
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
