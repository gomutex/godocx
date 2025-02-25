package docx

import (
	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/stypes"
)

type Hyperlink struct {
	root *RootDoc          // root is the root document to which this hyperlink belongs.
	ct   *ctypes.Hyperlink // ct is the underlying hyperlink element from the wml/ctypes package.
}

func newHyperlink(root *RootDoc, ct *ctypes.Hyperlink) *Hyperlink {
	return &Hyperlink{root: root, ct: ct}
}

// getProp returns the hyperlink properties. If not initialized, it creates and returns a new instance.
func (r *Hyperlink) getProp() *ctypes.RunProperty {
	if r.ct.Run.Property == nil {
		r.ct.Run.Property = &ctypes.RunProperty{}
	}
	return r.ct.Run.Property
}

// Sets the color of the Hyperlink.
//
// Example:
//
//	modifiedHyperlink := hyperlink.Color("FF0000")
//
// Parameters:
//   - colorCode: A string representing the color code (e.g., "FF0000" for red).
//
// Returns:
//   - *Hyperlink: The modified Hyperlink instance with the updated color.
func (r *Hyperlink) Color(colorCode string) *Hyperlink {
	r.getProp().Color = ctypes.NewColor(colorCode)
	return r
}

// Sets the size of the Hyperlink.

// This method takes an integer parameter representing the desired font size.
// It updates the size property of the Hyperlink instance with the specified size,
// Example:

// 	modifiedHyperlink := hyperlink.Size(12)

// Parameters:
//   - size: An integer representing the font size.

// Returns:
//   - *Hyperlink: The modified Hyperlink instance with the updated size.
func (r *Hyperlink) Size(size uint64) *Hyperlink {
	r.getProp().Size = ctypes.NewFontSize(size * 2)
	return r
}

// Font sets the font for the hyperlink.
func (r *Hyperlink) Font(font string) *Hyperlink {
	if r.getProp().Fonts == nil {
		r.getProp().Fonts = &ctypes.RunFonts{}
	}

	r.getProp().Fonts.Ascii = font
	r.getProp().Fonts.HAnsi = font
	return r
}

// Shading sets the shading properties (type, color, fill) for the hyperlink
func (r *Hyperlink) Shading(shdType stypes.Shading, color, fill string) *Hyperlink {
	r.getProp().Shading = ctypes.NewShading().SetShadingType(shdType).SetColor(color).SetFill(fill)
	return r
}

// AddHighlight sets the highlight color for the hyperlink.
func (r *Hyperlink) Highlight(color string) *Hyperlink {
	r.getProp().Highlight = ctypes.NewCTString(color)
	return r
}

// AddBold enables bold formatting for the hyperlink.
func (r *Hyperlink) Bold(value bool) *Hyperlink {
	r.getProp().Bold = ctypes.OnOffFromBool(value)
	return r
}

// Italic enables or disables italic formatting for the hyperlink.
func (r *Hyperlink) Italic(value bool) *Hyperlink {
	r.getProp().Italic = ctypes.OnOffFromBool(value)
	return r
}

// Specifies that the contents of this hyperlink shall be displayed with a single horizontal line through the center of the line.
func (r *Hyperlink) Strike(value bool) *Hyperlink {
	r.getProp().Strike = ctypes.OnOffFromBool(value)
	return r
}

// Specifies that the contents of this hyperlink shall be displayed with two horizontal lines through each character displayed on the line
func (r *Hyperlink) DoubleStrike(value bool) *Hyperlink {
	r.getProp().DoubleStrike = ctypes.OnOffFromBool(value)
	return r
}

// Display All Characters As Capital Letters
// Any lowercase characters in this text hyperlink shall be formatted for display only as their capital letter character equivalents
func (r *Hyperlink) Caps(value bool) *Hyperlink {
	r.getProp().Caps = ctypes.OnOffFromBool(value)
	return r
}

// Specifies that all small letter characters in this text hyperlink shall be formatted for display only as their capital letter character equivalents
func (r *Hyperlink) SmallCaps(value bool) *Hyperlink {
	r.getProp().Caps = ctypes.OnOffFromBool(value)
	return r
}

// Outline enables or disables outline formatting for the hyperlink.
func (r *Hyperlink) Outline(value bool) *Hyperlink {
	r.getProp().Outline = ctypes.OnOffFromBool(value)
	return r
}

// Shadow enables or disables shadow formatting for the hyperlink.
func (r *Hyperlink) Shadow(value bool) *Hyperlink {
	r.getProp().Shadow = ctypes.OnOffFromBool(value)
	return r
}

// Emboss enables or disables embossing formatting for the hyperlink.
func (r *Hyperlink) Emboss(value bool) *Hyperlink {
	r.getProp().Emboss = ctypes.OnOffFromBool(value)
	return r
}

// Imprint enables or disables imprint formatting for the hyperlink.
func (r *Hyperlink) Imprint(value bool) *Hyperlink {
	r.getProp().Imprint = ctypes.OnOffFromBool(value)
	return r
}

// Do Not Check Spelling or Grammar
func (r *Hyperlink) NoGrammer(value bool) *Hyperlink {
	r.getProp().NoGrammar = ctypes.OnOffFromBool(value)
	return r
}

// Use Document Grid Settings For Inter-Character Spacing
func (r *Hyperlink) SnapToGrid(value bool) *Hyperlink {
	r.getProp().SnapToGrid = ctypes.OnOffFromBool(value)
	return r
}

// Hidden Text
func (r *Hyperlink) HideText(value bool) *Hyperlink {
	r.getProp().Vanish = ctypes.OnOffFromBool(value)
	return r
}

// Spacing sets the spacing between characters in the hyperlink.
func (r *Hyperlink) Spacing(value int) *Hyperlink {
	r.getProp().Spacing = ctypes.NewDecimalNum(value)
	return r
}

// Underline sets the underline style for the hyperlink.
func (r *Hyperlink) Underline(value stypes.Underline) *Hyperlink {
	r.getProp().Underline = ctypes.NewGenSingleStrVal(value)
	return r
}

// Style sets the style of the Hyperlink.
func (r *Hyperlink) Style(value string) *Hyperlink {
	r.getProp().Style = ctypes.NewRunStyle(value)
	return r
}

// VerticalAlign sets the vertical alignment for the hyperlink text.
//
// Parameter: A value from the stypes.VerticalAlignRun type indicating the desired vertical alignment. One of:
//
//	VerticalAlignRunBaseline, VerticalAlignRunSuperscript, VerticalAlignRunSubscript
//
// Returns: The modified Hyperlink instance with the updated vertical alignment.
func (r *Hyperlink) VerticalAlign(value stypes.VerticalAlignRun) *Hyperlink {
	r.getProp().VertAlign = ctypes.NewGenSingleStrVal(value)
	return r
}
