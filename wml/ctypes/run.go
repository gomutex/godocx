package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/dml"
	"github.com/gomutex/godocx/internal"
	"github.com/gomutex/godocx/wml/stypes"
)

// A Run is part of a paragraph that has its own style. It could be
type Run struct {
	// Attributes
	RsidRPr *stypes.LongHexNum // Revision Identifier for Run Properties
	RsidR   *stypes.LongHexNum // Revision Identifier for Run
	RsidDel *stypes.LongHexNum // Revision Identifier for Run Deletion

	// Sequence:

	//1. Run Properties
	Property *RunProperty

	// 2. Choice - Run Inner content
	Children []RunChild
}

type RunChild struct {
	//specifies that a break shall be placed at the current location in the run content
	Break *Break `xml:"br,omitempty"`

	//specifies that this run contains literal text which shall be displayed in the document
	Text *Text `xml:"t,omitempty"`

	//specifies that this run contains literal text which shall be displayed in the document
	DelText *Text `xml:"delText,omitempty"`

	//Field Code
	InstrText *Text `xml:"instrText,omitempty"`

	//Deleted Field Code
	DelInstrText *Text `xml:"delInstrText,omitempty"`

	//Non Breaking Hyphen Character
	NoBreakHyphen *Empty `xml:"noBreakHyphen,omitempty"`

	//Non Breaking Hyphen Character
	SoftHyphen *Empty `xml:"softHyphen,omitempty"`

	//Date Block - Short Day Format
	DayShort *Empty `xml:"dayShort,omitempty"`

	//Date Block - Short Month Format
	MonthShort *Empty `xml:"monthShort,omitempty"`

	//Date Block - Short Year Format
	YearShort *Empty `xml:"yearShort,omitempty"`

	//Date Block - Long Day Format
	DayLong *Empty `xml:"dayLong,omitempty"`

	//Date Block - Long Month Format
	MonthLong *Empty `xml:"monthLong,omitempty"`

	//Date Block - Long Year Format
	YearLong *Empty `xml:"yearLong,omitempty"`

	//Comment Information Block
	AnnotationRef *Empty `xml:"annotationRef,omitempty"`

	//Footnote Reference Mark
	FootnoteRef *Empty `xml:"footnoteRef,omitempty"`

	//Endnote Reference Mark
	EndnoteRef *Empty `xml:"endnoteRef,omitempty"`

	//Footnote/Endnote Separator Mark
	Separator *Empty `xml:"separator,omitempty"`

	//Continuation Separator Mark
	ContSeparator *Empty `xml:"continuationSeparator,omitempty"`

	//Symbol Character
	Sym *Sym `xml:"sym,omitempty"`

	//Page Number Block
	PgNumBlock *Empty `xml:"pgNum,omitempty"`

	//Carriage Return
	CarrRtn *Empty `xml:"cr,omitempty"`

	//Tab Character
	Tab *Empty `xml:"tab,omitempty"`

	//TODO:
	// 	w:object    Inline Embedded Object
	// w:pict    VML Object
	// w:fldChar    Complex Field Character
	// w:ruby    Phonetic Guide
	// w:footnoteReference    Footnote Reference
	// w:endnoteReference    Endnote Reference
	// w:commentReference    Comment Content Reference Mark

	//Comment Content Reference Mark
	CmntRef *Markup `xml:"commentReference,omitempty"`

	//DrawingML Object
	Drawing *dml.Drawing `xml:"drawing,omitempty"`

	//Absolute Position Tab Character
	PTab *PTab `xml:"ptab,omitempty"`

	//Position of Last Calculated Page Break
	LastRenPgBrk *Empty `xml:"lastRenderedPageBreak,omitempty"`
}

func NewRun() *Run {
	return &Run{}
}

// Get the Run property. If it is not initialized, create one and return it.
func (r *Run) getProp() *RunProperty {
	if r.Property == nil {
		r.Property = &RunProperty{}
	}
	return r.Property
}

// Sets the color of the Run.
//
// Example:
//
//	run := NewRun()
//	modifiedRun := run.Color("FF0000")
//
// Parameters:
//   - colorCode: A string representing the color code (e.g., "FF0000" for red).
//
// Returns:
//   - *Run: The modified Run instance with the updated color.
func (r *Run) Color(colorCode string) *Run {
	r.getProp().Color = NewColor(colorCode)
	return r
}

// Sets the size of the Run.

// This method takes an integer parameter representing the desired font size.
// It updates the size property of the Run instance with the specified size,
// Example:

// 	run := NewRun()
// 	modifiedRun := run.Size(12)

// Parameters:
//   - size: An integer representing the font size.

// Returns:
//   - *Run: The modified Run instance with the updated size.
func (r *Run) Size(size uint64) *Run {
	r.getProp().Size = NewFontSize(size * 2)
	return r
}

func (r *Run) Shading(shdType stypes.Shading, color, fill string) *Run {
	r.getProp().Shading = NewShading().SetShadingType(shdType).SetColor(color).SetFill(fill)
	return r
}

// AddHighlight sets the highlight color for the run.
func (r *Run) Highlight(color string) *Run {
	r.getProp().Highlight = NewCTString(color)
	return r
}

// AddBold enables bold formatting for the run.
func (r *Run) Bold(value bool) *Run {
	r.getProp().Bold = OnOffFromBool(value)
	return r
}

func (r *Run) Italic(value bool) *Run {
	r.getProp().Italic = OnOffFromBool(value)
	return r
}

// Specifies that the contents of this run shall be displayed with a single horizontal line through the center of the line.
func (r *Run) Strike(value bool) *Run {
	r.getProp().Strike = OnOffFromBool(value)
	return r
}

// Specifies that the contents of this run shall be displayed with two horizontal lines through each character displayed on the line
func (r *Run) DoubleStrike(value bool) *Run {
	r.getProp().DoubleStrike = OnOffFromBool(value)
	return r
}

// Display All Characters As Capital Letters
//
// Any lowercase characters in this text run shall be formatted for display only as their capital letter character equivalents
func (r *Run) Caps(value bool) *Run {
	r.getProp().Caps = OnOffFromBool(value)
	return r
}

// Specifies that all small letter characters in this text run shall be formatted for display only as their capital letter character equivalents
func (r *Run) SmallCaps(value bool) *Run {
	r.getProp().Caps = OnOffFromBool(value)
	return r
}

func (r *Run) Outline(value bool) *Run {
	r.getProp().Outline = OnOffFromBool(value)
	return r
}

func (r *Run) Shadow(value bool) *Run {
	r.getProp().Shadow = OnOffFromBool(value)
	return r
}

func (r *Run) Emboss(value bool) *Run {
	r.getProp().Emboss = OnOffFromBool(value)
	return r
}

func (r *Run) Imprint(value bool) *Run {
	r.getProp().Imprint = OnOffFromBool(value)
	return r
}

// Do Not Check Spelling or Grammar
func (r *Run) NoGrammer(value bool) *Run {
	r.getProp().NoGrammar = OnOffFromBool(value)
	return r
}

// Use Document Grid Settings For Inter-Character Spacing
func (r *Run) SnapToGrid(value bool) *Run {
	r.getProp().SnapToGrid = OnOffFromBool(value)
	return r
}

// Hidden Text
func (r *Run) HideText(value bool) *Run {
	r.getProp().Vanish = OnOffFromBool(value)
	return r
}

func (r *Run) Spacing(value int) *Run {
	r.getProp().Spacing = NewDecimalNum(value)
	return r
}

func (r *Run) Underline(value stypes.Underline) *Run {
	r.getProp().Underline = NewGenSingleStrVal(value)
	return r
}

func (r *Run) Style(value string) *Run {
	r.getProp().Style = NewRunStyle(value)
	return r
}

func (r Run) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:r"

	if r.RsidRPr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidRPr"}, Value: string(*r.RsidRPr)})
	}
	if r.RsidR != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidR"}, Value: string(*r.RsidR)})
	}
	if r.RsidDel != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidDel"}, Value: string(*r.RsidDel)})
	}

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	// 1. Property
	if r.Property != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:rPr"}}
		if err = e.EncodeElement(r.Property, propsElement); err != nil {
			return err
		}
	}

	// 2. Remaining Child elemens
	if err = r.MarshalChild(e); err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (r *Run) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	// Decode attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "rsidRPr":
			r.RsidRPr = internal.ToPtr(stypes.LongHexNum(attr.Value))
		case "rsidR":
			r.RsidR = internal.ToPtr(stypes.LongHexNum(attr.Value))
		case "rsidDel":
			r.RsidDel = internal.ToPtr(stypes.LongHexNum(attr.Value))
		}
	}

loop:
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "t":
				txt := NewText()
				if err = d.DecodeElement(txt, &elem); err != nil {
					return err
				}

				r.Children = append(r.Children, RunChild{Text: txt})
			case "rPr":
				r.Property = &RunProperty{}
				if err = d.DecodeElement(r.Property, &elem); err != nil {
					return err
				}
			case "tab":
				tabElem := &Empty{}
				if err = d.DecodeElement(tabElem, &elem); err != nil {
					return err
				}

				r.Children = append(r.Children, RunChild{
					Tab: tabElem,
				})
			case "br":
				br := Break{}
				if err = d.DecodeElement(&br, &elem); err != nil {
					return err
				}

				r.Children = append(r.Children, RunChild{
					Break: &br,
				})
			case "drawing":
				drawingElem := &dml.Drawing{}
				if err = d.DecodeElement(drawingElem, &elem); err != nil {
					return err
				}

				r.Children = append(r.Children, RunChild{
					Drawing: drawingElem,
				})
			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break loop
		}
	}

	return nil
}

// Sym represents a symbol character in a document.
type Sym struct {
	Font *string `xml:"font,attr,omitempty"`
	Char *string `xml:"char,attr,omitempty"`
}

func NewSym(font, char string) *Sym {
	return &Sym{
		Font: &font,
		Char: &char,
	}
}

func (s Sym) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	if s.Font != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:font"}, Value: *s.Font})
	}

	if s.Char != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:char"}, Value: *s.Char})
	}

	return e.EncodeElement("", start)
}

func (r *Run) MarshalChild(e *xml.Encoder) error {
	var err error
	for _, child := range r.Children {
		switch {
		case child.Break != nil:
			err = child.Break.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:br"}})
		case child.Text != nil:
			err = child.Text.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:t"}})
		case child.DelText != nil:
			err = child.DelText.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:delText"}})
		case child.InstrText != nil:
			err = child.InstrText.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:instrText"}})
		case child.DelInstrText != nil:
			err = child.DelInstrText.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:delInstrText"}})
		case child.NoBreakHyphen != nil:
			err = child.NoBreakHyphen.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:noBreakHyphen"}})
		case child.SoftHyphen != nil:
			err = child.SoftHyphen.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:softHyphen"}})
		case child.DayShort != nil:
			err = child.DayShort.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:dayShort"}})
		case child.MonthShort != nil:
			err = child.MonthShort.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:monthShort"}})
		case child.YearShort != nil:
			err = child.YearShort.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:yearShort"}})
		case child.DayLong != nil:
			err = child.DayLong.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:dayLong"}})
		case child.MonthLong != nil:
			err = child.MonthLong.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:monthLong"}})
		case child.YearLong != nil:
			err = child.YearLong.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:yearLong"}})
		case child.AnnotationRef != nil:
			err = child.AnnotationRef.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:annotationRef"}})
		case child.FootnoteRef != nil:
			err = child.FootnoteRef.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:footnoteRef"}})
		case child.EndnoteRef != nil:
			err = child.EndnoteRef.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:endnoteRef"}})
		case child.Separator != nil:
			err = child.Separator.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:separator"}})
		case child.ContSeparator != nil:
			err = child.ContSeparator.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:continuationSeparator"}})
		case child.Sym != nil:
			err = child.Sym.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:sym"}})
		case child.PgNumBlock != nil:
			err = child.PgNumBlock.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:pgNum"}})
		case child.CarrRtn != nil:
			err = child.CarrRtn.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:cr"}})
		case child.Tab != nil:
			err = child.Tab.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:tab"}})
		case child.Drawing != nil:
			err = child.Drawing.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:drawing"}})
		case child.LastRenPgBrk != nil:
			err = child.LastRenPgBrk.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:lastRenderedPageBreak"}})
		case child.PTab != nil:
			err = child.PTab.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:ptab"}})
		case child.CmntRef != nil:
			err = child.CmntRef.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:commentReference"}})

		}

		if err != nil {
			return err
		}
	}
	return nil
}

// Add a break element of `stypes.BreakType` to this run.
func (r *Run) AddBreak(breakType *stypes.BreakType) {
	// clear := stypes.BreakClearNone
	// switch breakType{
	// case stypes.BreakType:

	// }
	br := Break{}

	if breakType != nil {
		br.BreakType = breakType
	}

	r.Children = append(r.Children, RunChild{
		Break: &br,
	})
}
