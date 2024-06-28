package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/internal"

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

	expectedParagraphProp := ParagraphProp{
		Style:         NewParagraphStyle("Heading1"),
		Justification: NewGenSingleStrVal(stypes.JustificationCenter),
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
		Style: &CTString{
			Val: "teststyle",
		},
		NumProp: &NumProp{
			ILvl: NewDecimalNum(10),
		},
		Indent: &Indent{
			Left: internal.ToPtr(10),
		},
		Spacing: &Spacing{
			After: internal.ToPtr(uint64(10)),
		},
		Border: &ParaBorder{
			Top: &Border{
				Val: "single",
			},
		},
		FrameProp: &FrameProp{
			Width: internal.ToPtr(int64(10)),
		},
		KeepNext:            &OnOff{},
		KeepLines:           &OnOff{},
		PageBreakBefore:     &OnOff{},
		WindowControl:       &OnOff{},
		SuppressLineNmbrs:   &OnOff{},
		SuppressAutoHyphens: &OnOff{},
		Kinsoku:             &OnOff{},
		WordWrap:            &OnOff{},
		OverflowPunct:       &OnOff{},
		TopLinePunct:        &OnOff{},
		AutoSpaceDE:         &OnOff{},
		AutoSpaceDN:         &OnOff{},
		Bidi:                &OnOff{},
		AdjustRightInd:      &OnOff{},
		SnapToGrid:          &OnOff{},
		CtxlSpacing:         &OnOff{},
		MirrorIndents:       &OnOff{},
		SuppressOverlap:     &OnOff{},
		Shading: &Shading{
			Val: stypes.ShdDiagStripe,
		},
		TextDirection:    NewGenSingleStrVal(stypes.TextDirectionBtLr),
		TextAlignment:    NewGenSingleStrVal(stypes.TextAlignAuto),
		TextboxTightWrap: NewGenSingleStrVal(stypes.TextboxTightWrapAllLines),
		Tabs: Tabs{
			Tab: []Tab{
				{
					Val: stypes.CustTabStopBar,
				},
			},
		},
		OutlineLvl: &DecimalNum{
			Val: 10000,
		},
		DivID: &DecimalNum{
			Val: 2222222,
		},
		CnfStyle: &CTString{
			Val: "cnftest",
		},
		RunProperty: &RunProperty{
			Shading: DefaultShading(),
		},
		Justification: NewGenSingleStrVal(stypes.JustificationCenter),
		PPrChange: &PPrChange{
			ID:     1,
			Author: "authortest",
		},
		SectPr: &SectionProp{
			TitlePg: NewGenSingleStrVal(stypes.OnOffTrue),
		},
	}

	// marshaledXML, err := xml.Marshal(paraProp)
	// if err != nil {
	// 	t.Fatalf("Error marshaling XML: %v", err)
	// }

	var result strings.Builder
	encoder := xml.NewEncoder(&result)
	start := xml.StartElement{Name: xml.Name{Local: "w:pPr"}}

	err := paraProp.MarshalXML(encoder, start)
	if err != nil {
		t.Fatalf("Error marshaling XML: %v", err)
	}

	encoder.Flush()
	if err != nil {
		t.Errorf("Error flushing encoder: %v", err)
	}

	marshaledXML := result.String()

	// Unmarshal the marshaled XML back into a new ParagraphProp instance
	var unmarshaledProp ParagraphProp
	err = xml.Unmarshal([]byte(marshaledXML), &unmarshaledProp)
	if err != nil {
		t.Fatalf("Error unmarshaling XML: %v", err)
	}

	// Compare using the custom equality function
	if err := isEqualParagraphProps(paraProp, unmarshaledProp); err != nil {
		t.Error(err)
	}

	var rresult strings.Builder
	rencoder := xml.NewEncoder(&rresult)

	err = paraProp.MarshalXML(rencoder, start)
	if err != nil {
		t.Fatalf("Error marshaling XML: %v", err)
	}

	rencoder.Flush()
	if err != nil {
		t.Errorf("Error flushing encoder: %v", err)
	}

	if result.String() != rresult.String() {
		t.Errorf("Expected XML :\n%s\nGot:\n%s", result.String(), rresult.String())
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

// <== ParaProp Tests end here ==>

func TestPPrChange_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    PPrChange
		expected string
	}{
		{
			name: "With all attributes",
			input: PPrChange{
				ID:       123,
				Author:   "John Doe",
				Date:     internal.ToPtr("2024-06-19"),
				ParaProp: &ParagraphProp{
					// Initialize ParagraphProp fields here if needed
				},
			},
			expected: `<w:pPrChange id="123" author="John Doe" date="2024-06-19"><w:pPr></w:pPr></w:pPrChange>`,
		},
		{
			name: "Without date attribute",
			input: PPrChange{
				ID:       456,
				Author:   "Jane Smith",
				ParaProp: &ParagraphProp{
					// Initialize ParagraphProp fields here if needed
				},
			},
			expected: `<w:pPrChange id="456" author="Jane Smith"><w:pPr></w:pPr></w:pPrChange>`,
		},
		{
			name: "Without paraProp",
			input: PPrChange{
				ID:     789,
				Author: "Alice Brown",
				Date:   internal.ToPtr("2024-06-20"),
			},
			expected: `<w:pPrChange id="789" author="Alice Brown" date="2024-06-20"></w:pPrChange>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)

			start := xml.StartElement{Name: xml.Name{Local: "w:pPrChange"}}
			if err := tt.input.MarshalXML(encoder, start); err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			// Finalize encoding
			encoder.Flush()

			got := strings.TrimSpace(result.String())
			if got != tt.expected {
				t.Errorf("Expected XML:\n%s\nGot:\n%s", tt.expected, got)
			}
		})
	}
}

// <== PPrChange Test end here ==>
