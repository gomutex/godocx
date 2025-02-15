package docx

import (
	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/stypes"
)

type Run struct {
	root *RootDoc    // root is the root document to which this run belongs.
	ct   *ctypes.Run // ct is the underlying run element from the wml/ctypes package.
}

func newRun(root *RootDoc, ct *ctypes.Run) *Run {
	return &Run{root: root, ct: ct}
}

// getProp returns the run properties. If not initialized, it creates and returns a new instance.
func (r *Run) getProp() *ctypes.RunProperty {
	if r.ct.Property == nil {
		r.ct.Property = &ctypes.RunProperty{}
	}
	return r.ct.Property
}

// Sets the color of the Run.
//
// Example:
//
//	modifiedRun := run.Color("FF0000")
//
// Parameters:
//   - colorCode: A string representing the color code (e.g., "FF0000" for red).
//
// Returns:
//   - *Run: The modified Run instance with the updated color.
func (r *Run) Color(colorCode string) *Run {
	r.getProp().Color = ctypes.NewColor(colorCode)
	return r
}

// Sets the size of the Run.

// This method takes an integer parameter representing the desired font size.
// It updates the size property of the Run instance with the specified size,
// Example:

// 	modifiedRun := run.Size(12)

// Parameters:
//   - size: An integer representing the font size.

// Returns:
//   - *Run: The modified Run instance with the updated size.
func (r *Run) Size(size uint64) *Run {
	r.getProp().Size = ctypes.NewFontSize(size * 2)
	return r
}

// Font sets the font for the run.
func (r *Run) Font(font string) *Run {
	if r.getProp().Fonts == nil {
		r.getProp().Fonts = &ctypes.RunFonts{}
	}

	r.getProp().Fonts.Ascii = font
	r.getProp().Fonts.HAnsi = font
	return r
}

// Shading sets the shading properties (type, color, fill) for the run
func (r *Run) Shading(shdType stypes.Shading, color, fill string) *Run {
	r.getProp().Shading = ctypes.NewShading().SetShadingType(shdType).SetColor(color).SetFill(fill)
	return r
}

// AddHighlight sets the highlight color for the run.
func (r *Run) Highlight(color string) *Run {
	r.getProp().Highlight = ctypes.NewCTString(color)
	return r
}

// AddBold enables bold formatting for the run.
func (r *Run) Bold(value bool) *Run {
	r.getProp().Bold = ctypes.OnOffFromBool(value)
	return r
}

// Italic enables or disables italic formatting for the run.
func (r *Run) Italic(value bool) *Run {
	r.getProp().Italic = ctypes.OnOffFromBool(value)
	return r
}

// Specifies that the contents of this run shall be displayed with a single horizontal line through the center of the line.
func (r *Run) Strike(value bool) *Run {
	r.getProp().Strike = ctypes.OnOffFromBool(value)
	return r
}

// Specifies that the contents of this run shall be displayed with two horizontal lines through each character displayed on the line
func (r *Run) DoubleStrike(value bool) *Run {
	r.getProp().DoubleStrike = ctypes.OnOffFromBool(value)
	return r
}

// Display All Characters As Capital Letters
// Any lowercase characters in this text run shall be formatted for display only as their capital letter character equivalents
func (r *Run) Caps(value bool) *Run {
	r.getProp().Caps = ctypes.OnOffFromBool(value)
	return r
}

// Specifies that all small letter characters in this text run shall be formatted for display only as their capital letter character equivalents
func (r *Run) SmallCaps(value bool) *Run {
	r.getProp().Caps = ctypes.OnOffFromBool(value)
	return r
}

// Outline enables or disables outline formatting for the run.
func (r *Run) Outline(value bool) *Run {
	r.getProp().Outline = ctypes.OnOffFromBool(value)
	return r
}

// Shadow enables or disables shadow formatting for the run.
func (r *Run) Shadow(value bool) *Run {
	r.getProp().Shadow = ctypes.OnOffFromBool(value)
	return r
}

// Emboss enables or disables embossing formatting for the run.
func (r *Run) Emboss(value bool) *Run {
	r.getProp().Emboss = ctypes.OnOffFromBool(value)
	return r
}

// Imprint enables or disables imprint formatting for the run.
func (r *Run) Imprint(value bool) *Run {
	r.getProp().Imprint = ctypes.OnOffFromBool(value)
	return r
}

// Do Not Check Spelling or Grammar
func (r *Run) NoGrammer(value bool) *Run {
	r.getProp().NoGrammar = ctypes.OnOffFromBool(value)
	return r
}

// Use Document Grid Settings For Inter-Character Spacing
func (r *Run) SnapToGrid(value bool) *Run {
	r.getProp().SnapToGrid = ctypes.OnOffFromBool(value)
	return r
}

// Hidden Text
func (r *Run) HideText(value bool) *Run {
	r.getProp().Vanish = ctypes.OnOffFromBool(value)
	return r
}

// Spacing sets the spacing between characters in the run.
func (r *Run) Spacing(value int) *Run {
	r.getProp().Spacing = ctypes.NewDecimalNum(value)
	return r
}

// Underline sets the underline style for the run.
func (r *Run) Underline(value stypes.Underline) *Run {
	r.getProp().Underline = ctypes.NewGenSingleStrVal(value)
	return r
}

// Add a break element of `stypes.BreakType` to this run.
func (r *Run) AddBreak(breakType *stypes.BreakType) {
	// clear := stypes.BreakClearNone
	// switch breakType{
	// case stypes.BreakType:

	// }
	br := ctypes.Break{}

	if breakType != nil {
		br.BreakType = breakType
	}

	r.ct.Children = append(r.ct.Children, ctypes.RunChild{
		Break: &br,
	})
}

// Style sets the style of the run.
func (r *Run) Style(value string) *Run {
	r.getProp().Style = ctypes.NewRunStyle(value)
	return r
}

// VerticalAlign sets the vertical alignment for the run text.
//
// Parameter: A value from the stypes.VerticalAlignRun type indicating the desired vertical alignment. One of:
//
//	VerticalAlignRunBaseline, VerticalAlignRunSuperscript, VerticalAlignRunSubscript
//
// Returns: The modified Run instance with the updated vertical alignment.
func (r *Run) VerticalAlign(value stypes.VerticalAlignRun) *Run {
	r.getProp().VertAlign = ctypes.NewGenSingleStrVal(value)
	return r
}
