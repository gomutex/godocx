package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/internal"
	"github.com/gomutex/godocx/wml/stypes"
)

func TestCellProperty_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    CellProperty
		expected string
	}{
		{
			name:     "Empty CellProperty",
			input:    CellProperty{},
			expected: `<w:tcPr></w:tcPr>`,
		},
		{
			name: "With Width",
			input: CellProperty{
				Width: NewTableWidth(100, stypes.TableWidthDxa),
			},
			expected: `<w:tcPr><w:tcW w:w="100" w:type="dxa"></w:tcW></w:tcPr>`,
		},
		{
			name: "With Shading",
			input: CellProperty{
				Shading: &Shading{Val: "clear"},
			},
			expected: `<w:tcPr><w:shd w:val="clear"></w:shd></w:tcPr>`,
		},
		{
			name: "With All Fields",
			input: CellProperty{
				CnfStyle:      &CTString{Val: "TestCnfStyle"},
				Width:         NewTableWidth(100, stypes.TableWidthDxa),
				GridSpan:      &DecimalNum{Val: 2},
				HMerge:        NewGenOptStrVal(stypes.MergeCellContinue),
				VMerge:        NewGenOptStrVal(stypes.MergeCellRestart),
				Borders:       &CellBorders{},
				Shading:       &Shading{Val: "clear"},
				NoWrap:        &OnOff{Val: internal.ToPtr(stypes.OnOffTrue)},
				Margins:       &CellMargins{},
				TextDirection: NewGenSingleStrVal(stypes.TextDirectionBtLr),
				FitText:       &OnOff{Val: internal.ToPtr(stypes.OnOffTrue)},
				VAlign:        NewGenSingleStrVal(stypes.VerticalJcCenter),
				HideMark:      &OnOff{Val: internal.ToPtr(stypes.OnOffTrue)},
				PrChange:      &TCPrChange{ID: 1, Author: "Author", Date: nil},
			},
			expected: `<w:tcPr>` +
				`<w:cnfStyle w:val="TestCnfStyle"></w:cnfStyle>` +
				`<w:tcW w:w="100" w:type="dxa"></w:tcW>` +
				`<w:gridSpan w:val="2"></w:gridSpan>` +
				`<w:hMerge w:val="continue"></w:hMerge>` +
				`<w:vMerge w:val="restart"></w:vMerge>` +
				`<w:tcBorders></w:tcBorders>` +
				`<w:shd w:val="clear"></w:shd>` +
				`<w:noWrap w:val="true"></w:noWrap>` +
				`<w:tcMar></w:tcMar>` +
				`<w:textDirection w:val="btLr"></w:textDirection>` +
				`<w:tcFitText w:val="true"></w:tcFitText>` +
				`<w:vAlign w:val="center"></w:vAlign>` +
				`<w:hideMark w:val="true"></w:hideMark>` +
				`<w:tcPrChange w:id="1" w:author="Author"><w:tcPr></w:tcPr></w:tcPrChange>` +
				`</w:tcPr>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)

			start := xml.StartElement{Name: xml.Name{Local: "w:tcPr"}}
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

func TestCellProperty_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name       string
		inputXML   string
		expected   CellProperty
		expectFail bool // Whether unmarshalling is expected to fail
	}{
		{
			name:     "Empty CellProperty",
			inputXML: `<w:tcPr></w:tcPr>`,
			expected: CellProperty{},
		},
		{
			name: "With Width",
			inputXML: `<w:tcPr>
						<w:tcW w:w="100" w:type="dxa"></w:tcW>
						</w:tcPr>`,
			expected: CellProperty{
				Width: NewTableWidth(100, stypes.TableWidthDxa),
			},
		},
		{
			name: "With Shading",
			inputXML: `<w:tcPr>
						<w:shd w:val="clear"></w:shd>
						</w:tcPr>`,
			expected: CellProperty{
				Shading: &Shading{Val: "clear"},
			},
		},
		{
			name: "With All Fields",
			inputXML: `<w:tcPr>` +
				`<w:cnfStyle w:val="TestCnfStyle"></w:cnfStyle>` +
				`<w:tcW w:w="100" w:type="dxa"></w:tcW>` +
				`<w:gridSpan w:val="2"></w:gridSpan>` +
				`<w:hMerge w:val="continue"></w:hMerge>` +
				`<w:vMerge w:val="restart"></w:vMerge>` +
				`<w:tcBorders></w:tcBorders>` +
				`<w:shd w:val="clear"></w:shd>` +
				`<w:noWrap w:val="1"></w:noWrap>` +
				`<w:tcMar></w:tcMar>` +
				`<w:textDirection w:val="btLr"></w:textDirection>` +
				`<w:tcFitText w:val="1"></w:tcFitText>` +
				`<w:vAlign w:val="center"></w:vAlign>` +
				`<w:hideMark w:val="1"></w:hideMark>` +
				`<w:cellIns w:id="1" w:author="Author"></w:cellIns>` +
				`<w:tcPrChange w:id="1" w:author="Author"><w:tcPr></w:tcPr></w:tcPrChange>` +
				`</w:tcPr>`,
			expected: CellProperty{
				CnfStyle:      &CTString{Val: "TestCnfStyle"},
				Width:         NewTableWidth(100, stypes.TableWidthDxa),
				GridSpan:      &DecimalNum{Val: 2},
				HMerge:        NewGenOptStrVal(stypes.MergeCellContinue),
				VMerge:        NewGenOptStrVal(stypes.MergeCellRestart),
				Borders:       &CellBorders{},
				Shading:       &Shading{Val: "clear"},
				NoWrap:        &OnOff{Val: internal.ToPtr(stypes.OnOffTrue)},
				Margins:       &CellMargins{},
				TextDirection: NewGenSingleStrVal(stypes.TextDirectionBtLr),
				FitText:       &OnOff{Val: internal.ToPtr(stypes.OnOffTrue)},
				VAlign:        NewGenSingleStrVal(stypes.VerticalJcCenter),
				HideMark:      &OnOff{Val: internal.ToPtr(stypes.OnOffTrue)},
				CellInsertion: &TrackChange{
					ID:     1,
					Author: "Author",
				},
				PrChange: &TCPrChange{ID: 1, Author: "Author", Date: nil},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoder := xml.NewDecoder(strings.NewReader(tt.inputXML))
			var result CellProperty

			err := decoder.Decode(&result)

			if tt.expectFail {
				if err == nil {
					t.Error("Expected unmarshaling to fail but it did not")
				}
				return
			}

			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if err := compareCellProperties(tt.expected, result); err != nil {
				t.Errorf("Unmarshaled CellProperty struct does not match expected: %v", err)
			}
		})

	}
}

// Helper function to compare CellProperty structs
func compareCellProperties(a, b CellProperty) error {
	if err := internal.ComparePtr("CnfStyle", a.CnfStyle, b.CnfStyle); err != nil {
		return err
	}
	if err := internal.ComparePtr("Width", a.Width, b.Width); err != nil {
		return err
	}
	if err := internal.ComparePtr("GridSpan", a.GridSpan, b.GridSpan); err != nil {
		return err
	}
	if err := internal.ComparePtr("HMerge", a.HMerge, b.HMerge); err != nil {
		return err
	}
	if err := internal.ComparePtr("VMerge", a.VMerge, b.VMerge); err != nil {
		return err
	}
	if err := internal.ComparePtr("Borders", a.Borders, b.Borders); err != nil {
		return err
	}
	if err := internal.ComparePtr("Shading", a.Shading, b.Shading); err != nil {
		return err
	}
	if err := internal.ComparePtr("NoWrap", a.NoWrap, b.NoWrap); err != nil {
		return err
	}
	if err := internal.ComparePtr("Margins", a.Margins, b.Margins); err != nil {
		return err
	}
	if err := internal.ComparePtr("TextDirection", a.TextDirection, b.TextDirection); err != nil {
		return err
	}
	if err := internal.ComparePtr("FitText", a.FitText, b.FitText); err != nil {
		return err
	}
	if err := internal.ComparePtr("VAlign", a.VAlign, b.VAlign); err != nil {
		return err
	}
	if err := internal.ComparePtr("HideMark", a.HideMark, b.HideMark); err != nil {
		return err
	}
	if err := internal.ComparePtr("CellInsertion", a.CellInsertion, b.CellInsertion); err != nil {
		return err
	}
	if err := internal.ComparePtr("CellDeletion", a.CellDeletion, b.CellDeletion); err != nil {
		return err
	}
	if err := internal.ComparePtr("CellMerge", a.CellMerge, b.CellMerge); err != nil {
		return err
	}
	if err := internal.ComparePtr("PrChange", a.PrChange, b.PrChange); err != nil {
		return err
	}
	return nil
}
