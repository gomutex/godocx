package docxpara

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/elemtypes"
	"github.com/gomutex/godocx/internal"
	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/docxrun"
	"github.com/gomutex/godocx/wml/hdrftr"
	"github.com/gomutex/godocx/wml/sections"
	"github.com/gomutex/godocx/wml/stypes"
)

func areParagraphPropertiesEqual(p1, p2 ParagraphProp) bool {
	return p1.Style.Val == p2.Style.Val &&
		p1.Justification.Val == p2.Justification.Val
}

func TestParagraphProp(t *testing.T) {
	xmlString := `<w:pPr xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
  <w:pStyle w:val="Heading1"></w:pStyle>
  <w:jc w:val="center"></w:jc>
</w:pPr>`

	var parsedParagraphProp ParagraphProp
	err := xml.Unmarshal([]byte(xmlString), &parsedParagraphProp)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to ParagraphProp: %v", err)
	}

	jc, err := ctypes.NewJustification("center")
	if err != nil {
		t.Fatal(err)
	}

	expectedParagraphProp := ParagraphProp{
		Style:         NewParagraphStyle("Heading1"),
		Justification: jc,
	}

	if !areParagraphPropertiesEqual(expectedParagraphProp, parsedParagraphProp) {
		t.Errorf("Expected ParagraphProp %v, got %v", expectedParagraphProp, parsedParagraphProp)
	}
}

func TestNewParagraphStyle(t *testing.T) {
	expected := "TestStyle"
	style := NewParagraphStyle(expected)

	if style.Val != expected {
		t.Errorf("NewParagraphStyle() = %s; want %s", style.Val, expected)
	}
}

func TestDefaultParagraphStyle(t *testing.T) {
	expected := "Normal"
	style := DefaultParagraphStyle()

	if style.Val != expected {
		t.Errorf("DefaultParagraphStyle() = %s; want %s", style.Val, expected)
	}
}
func TestParagraphProp_MarshalUnmarshal(t *testing.T) {
	paraProp := ParagraphProp{
		Style: &elemtypes.SingleStrVal{
			Val: "teststyle",
		},
		NumProp: &NumProp{
			ILvl: ctypes.NewDecimalNum(10),
		},
		Indent: &Indent{
			Left: internal.ToPtr(10),
		},
		Spacing: &Spacing{
			After: internal.ToPtr(uint64(10)),
		},
		Border: &ParaBorder{
			Top: &ctypes.Border{
				Val: "single",
			},
		},
		FrameProp: &FrameProp{
			Width: internal.ToPtr(int64(10)),
		},
		KeepNext:            &elemtypes.OptBinFlagElem{},
		KeepLines:           &elemtypes.OptBinFlagElem{},
		PageBreakBefore:     &elemtypes.OptBinFlagElem{},
		WindowControl:       &elemtypes.OptBinFlagElem{},
		SuppressLineNmbrs:   &elemtypes.OptBinFlagElem{},
		SuppressAutoHyphens: &elemtypes.OptBinFlagElem{},
		Kinsoku:             &elemtypes.OptBinFlagElem{},
		WordWrap:            &elemtypes.OptBinFlagElem{},
		OverflowPunct:       &elemtypes.OptBinFlagElem{},
		TopLinePunct:        &elemtypes.OptBinFlagElem{},
		AutoSpaceDE:         &elemtypes.OptBinFlagElem{},
		AutoSpaceDN:         &elemtypes.OptBinFlagElem{},
		Bidi:                &elemtypes.OptBinFlagElem{},
		AdjustRightInd:      &elemtypes.OptBinFlagElem{},
		SnapToGrid:          &elemtypes.OptBinFlagElem{},
		CtxlSpacing:         &elemtypes.OptBinFlagElem{},
		MirrorIndents:       &elemtypes.OptBinFlagElem{},
		SuppressOverlap:     &elemtypes.OptBinFlagElem{},
		Shading: &ctypes.Shading{
			Val: stypes.ShdDiagStripe,
		},
		TextDirection: &ctypes.TextDirection{
			Val: stypes.TextDirectionBtLr,
		},
		TextAlignment: &ctypes.TextAlign{
			Val: stypes.TextAlignAuto,
		},
		TextboxTightWrap: &ctypes.TextboxTightWrap{
			Val: stypes.TextboxTightWrapAllLines,
		},
		Tabs: ctypes.Tabs{
			Tab: []ctypes.Tab{
				{
					Val: internal.ToPtr(stypes.CustTabStopBar),
				},
			},
		},
		OutlineLvl: &ctypes.DecimalNum{
			Val: 10000,
		},
		DivID: &ctypes.DecimalNum{
			Val: 2222222,
		},
		CnfStyle: &ctypes.Cnf{
			Val: "cnftest",
		},
		RunProperty: &docxrun.RunProperty{
			Shading: ctypes.DefaultShading(),
		},
		Justification: &ctypes.Justification{
			Val: "center",
		},
		PPrChange: &PPrChange{
			ID:     1,
			Author: "authortest",
		},
		SectPr: &sections.SectionProp{
			TitlePg: &hdrftr.TitlePg{},
		},
	}

	marshaledXML, err := xml.Marshal(paraProp)
	if err != nil {
		t.Fatalf("Error marshaling XML: %v", err)
	}

	// Unmarshal the marshaled XML back into a new ParagraphProp instance
	var unmarshaledProp ParagraphProp
	err = xml.Unmarshal(marshaledXML, &unmarshaledProp)
	if err != nil {
		t.Fatalf("Error unmarshaling XML: %v", err)
	}

	// Compare using the custom equality function
	if err := isEqualParagraphProps(paraProp, unmarshaledProp); err != nil {
		t.Error(err)
	}

	// For further validation, you can convert the XML back to string and compare
	var result strings.Builder
	encoder := xml.NewEncoder(&result)
	err = encoder.Encode(paraProp)
	if err != nil {
		t.Fatalf("Error encoding XML: %v", err)
	}

	if result.String() != string(marshaledXML) {
		t.Errorf("Expected XML:\n%s\nGot:\n%s", string(marshaledXML), result.String())
	}
}

// Define your helper function to compare two ParagraphProp structs
func isEqualParagraphProps(a, b ParagraphProp) error {
	// Compare each field using internal.ComparePtr for pointers and direct comparison for non-pointers
	if err := internal.ComparePtr("Style", a.Style, b.Style); err != nil {
		return err
	}
	if err := internal.ComparePtr("NumProp", a.NumProp, b.NumProp); err != nil {
		return err
	}
	if err := internal.ComparePtr("Indent", a.Indent, b.Indent); err != nil {
		return err
	}
	if err := internal.ComparePtr("Spacing", a.Spacing, b.Spacing); err != nil {
		return err
	}
	if err := internal.ComparePtr("Border", a.Border, b.Border); err != nil {
		return err
	}
	if err := internal.ComparePtr("FrameProp", a.FrameProp, b.FrameProp); err != nil {
		return err
	}
	if err := internal.ComparePtr("KeepNext", a.KeepNext, b.KeepNext); err != nil {
		return err
	}
	if err := internal.ComparePtr("KeepLines", a.KeepLines, b.KeepLines); err != nil {
		return err
	}
	if err := internal.ComparePtr("PageBreakBefore", a.PageBreakBefore, b.PageBreakBefore); err != nil {
		return err
	}
	if err := internal.ComparePtr("WindowControl", a.WindowControl, b.WindowControl); err != nil {
		return err
	}
	if err := internal.ComparePtr("SuppressLineNmbrs", a.SuppressLineNmbrs, b.SuppressLineNmbrs); err != nil {
		return err
	}
	if err := internal.ComparePtr("SuppressAutoHyphens", a.SuppressAutoHyphens, b.SuppressAutoHyphens); err != nil {
		return err
	}
	if err := internal.ComparePtr("Kinsoku", a.Kinsoku, b.Kinsoku); err != nil {
		return err
	}
	if err := internal.ComparePtr("WordWrap", a.WordWrap, b.WordWrap); err != nil {
		return err
	}
	if err := internal.ComparePtr("OverflowPunct", a.OverflowPunct, b.OverflowPunct); err != nil {
		return err
	}
	if err := internal.ComparePtr("TopLinePunct", a.TopLinePunct, b.TopLinePunct); err != nil {
		return err
	}
	if err := internal.ComparePtr("AutoSpaceDE", a.AutoSpaceDE, b.AutoSpaceDE); err != nil {
		return err
	}
	if err := internal.ComparePtr("AutoSpaceDN", a.AutoSpaceDN, b.AutoSpaceDN); err != nil {
		return err
	}
	if err := internal.ComparePtr("Bidi", a.Bidi, b.Bidi); err != nil {
		return err
	}
	if err := internal.ComparePtr("AdjustRightInd", a.AdjustRightInd, b.AdjustRightInd); err != nil {
		return err
	}
	if err := internal.ComparePtr("SnapToGrid", a.SnapToGrid, b.SnapToGrid); err != nil {
		return err
	}
	if err := internal.ComparePtr("CtxlSpacing", a.CtxlSpacing, b.CtxlSpacing); err != nil {
		return err
	}
	if err := internal.ComparePtr("MirrorIndents", a.MirrorIndents, b.MirrorIndents); err != nil {
		return err
	}
	if err := internal.ComparePtr("SuppressOverlap", a.SuppressOverlap, b.SuppressOverlap); err != nil {
		return err
	}

	if err := internal.ComparePtr("Shading", a.Shading, b.Shading); err != nil {
		return err
	}

	if err := internal.ComparePtr("TextDirection", a.TextDirection, b.TextDirection); err != nil {
		return err
	}

	if err := internal.ComparePtr("TextAlignment", a.TextAlignment, b.TextAlignment); err != nil {
		return err
	}

	if err := internal.ComparePtr("TextboxTightWrap", a.TextboxTightWrap, b.TextboxTightWrap); err != nil {
		return err
	}

	if err := internal.ComparePtr("OutlineLvl", a.OutlineLvl, b.OutlineLvl); err != nil {
		return err
	}

	if err := internal.ComparePtr("DivID", a.DivID, b.DivID); err != nil {
		return err
	}

	return nil
}
