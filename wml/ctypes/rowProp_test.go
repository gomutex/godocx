package ctypes

import (
	"encoding/xml"
	"errors"
	"strings"
	"testing"

	"github.com/gomutex/godocx/internal"
	"github.com/gomutex/godocx/wml/stypes"
)

func TestRowProperty_MarshalXML(t *testing.T) {
	tests := []struct {
		name        string
		input       RowProperty
		expectFail  bool   // Whether marshaling is expected to fail
		expectedXML string // Expected XML output for validation
	}{
		{
			name: "Only Cnf populated",
			input: RowProperty{
				Cnf: &CTString{
					Val: "cnfval",
				},
			},
			expectFail:  false,
			expectedXML: `<w:trPr><w:cnfStyle w:val="cnfval"></w:cnfStyle></w:trPr>`,
		},
		{
			name: "Only DivId populated",
			input: RowProperty{
				DivId: &DecimalNum{},
			},
			expectFail:  false,
			expectedXML: `<w:trPr><w:divId w:val="0"></w:divId></w:trPr>`,
		},
		{
			name: "All fields populated",
			input: RowProperty{
				Cnf: &CTString{
					Val: "cnftest",
				},
				DivId:       &DecimalNum{},
				GridBefore:  &DecimalNum{},
				GridAfter:   &DecimalNum{},
				WidthBefore: NewTableWidth(500, "pct"),
				WidthAfter:  NewTableWidth(300, "dxa"),
				CantSplit: &OnOff{
					Val: internal.ToPtr(stypes.OnOffTrue),
				},
				Height: NewTableRowHeight(500, "atLeast"),
				Header: &OnOff{
					Val: internal.ToPtr(stypes.OnOffTrue),
				},
				CellSpacing: NewTableWidth(100, "dxa"),
				JC:          NewGenSingleStrVal(stypes.JustificationCenter),
				Hidden: &OnOff{
					Val: internal.ToPtr(stypes.OnOffFalse),
				},
			},
			expectFail:  false,
			expectedXML: `<w:trPr><w:cnfStyle w:val="cnftest"></w:cnfStyle><w:divId w:val="0"></w:divId><w:gridBefore w:val="0"></w:gridBefore><w:gridAfter w:val="0"></w:gridAfter><w:tblWBefore w:w="500" w:type="pct"></w:tblWBefore><w:tblWAfter w:w="300" w:type="dxa"></w:tblWAfter><w:cantSplit w:val="true"></w:cantSplit><w:trHeight w:val="500" w:hRule="atLeast"></w:trHeight><w:tblHeader w:val="true"></w:tblHeader><w:tblCellSpacing w:w="100" w:type="dxa"></w:tblCellSpacing><w:jc w:val="center"></w:jc><w:hidden w:val="false"></w:hidden></w:trPr>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)

			start := xml.StartElement{Name: xml.Name{Local: "w:trPr"}}
			err := tt.input.MarshalXML(encoder, start)

			encoder.Flush()

			if tt.expectFail {
				if err == nil {
					t.Error("Expected error but got none")
				} else {
					expectedErrorMsg := "incomplete table row property, missing choice elements"
					if !strings.Contains(err.Error(), expectedErrorMsg) {
						t.Errorf("Expected error message to contain '%s' but got '%s'", expectedErrorMsg, err.Error())
					}
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error during marshaling: %v", err)
				} else {
					// Check the generated XML against expected XML
					generatedXML := result.String()
					if generatedXML != tt.expectedXML {
						t.Errorf("Generated XML does not match expected XML.\nExpected:\n%s\n\nGot:\n%s", tt.expectedXML, generatedXML)
					}
				}
			}
		})
	}
}

func TestRowProperty_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name       string
		inputXML   string
		expected   RowProperty
		expectFail bool // Whether unmarshaling is expected to fail
	}{
		{
			name: "Only Cnf populated",
			inputXML: `<w:trPr>
						<w:cnfStyle val="cnftest"></w:cnfStyle>
					   </w:trPr>`,
			expected: RowProperty{
				Cnf: &CTString{
					Val: "cnftest",
				},
			},
		},
		{
			name: "Only DivId populated",
			inputXML: `<w:trPr>
						<w:divId></w:divId>
					   </w:trPr>`,
			expected: RowProperty{
				DivId: &DecimalNum{},
			},
		},
		{
			name: "All fields populated",
			inputXML: `<w:trPr>
						<w:cnfStyle val="cnftest"></w:cnfStyle>
						<w:divId></w:divId>
						<w:gridBefore>2</w:gridBefore>
						<w:gridAfter>3</w:gridAfter>
						<w:tblWBefore w:type="pct" w:w="500"></w:tblWBefore>
						<w:tblWAfter w:type="dxa" w:w="300"></w:tblWAfter>
						<w:cantSplit w:val="on"></w:cantSplit>
						<w:trHeight w:val="500" w:hRule="atLeast"></w:trHeight>
						<w:tblHeader w:val="on"></w:tblHeader>
						<w:tblCellSpacing w:type="dxa" w:w="100"></w:tblCellSpacing>
						<w:jc w:val="center"></w:jc>
						<w:hidden w:val="off"></w:hidden>
					   </w:trPr>`,
			expected: RowProperty{
				Cnf: &CTString{
					Val: "cnftest",
				},
				DivId:       &DecimalNum{},
				GridBefore:  &DecimalNum{Val: 2},
				GridAfter:   &DecimalNum{Val: 3},
				WidthBefore: NewTableWidth(500, "pct"),
				WidthAfter:  NewTableWidth(300, "dxa"),
				CantSplit: &OnOff{
					Val: internal.ToPtr(stypes.OnOffTrue),
				},
				Height: NewTableRowHeight(500, "atLeast"),
				Header: &OnOff{
					Val: internal.ToPtr(stypes.OnOffTrue),
				},
				CellSpacing: NewTableWidth(100, "dxa"),
				JC:          NewGenSingleStrVal(stypes.JustificationCenter),
				Hidden: &OnOff{
					Val: internal.ToPtr(stypes.OnOffFalse),
				},
			},
			expectFail: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result RowProperty

			err := xml.Unmarshal([]byte(tt.inputXML), &result)

			if tt.expectFail {
				if err == nil {
					t.Error("Expected unmarshaling to fail but it did not")
				}
				return
			}

			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			// Compare each field individually due to pointer comparisons
			if err := compareRowProperty(tt.expected, result); err != nil {
				t.Error(err)
			}
		})
	}
}

// Helper function to compare RowProperty structs
func compareRowProperty(a, b RowProperty) error {

	// Compare Cnf fields
	if err := internal.ComparePtr("Cnf", a.Cnf, b.Cnf); err != nil {
		return err
	}

	// Compare SingleIntVal fields
	if err := internal.ComparePtr("DivId", a.DivId, b.DivId); err != nil {
		return err
	}

	// Compare GridBefore and GridAfter fields
	if err := internal.ComparePtr("GridBefore", a.GridBefore, b.GridBefore); err != nil {
		return err
	}
	if err := internal.ComparePtr("GridAfter", a.GridAfter, b.GridAfter); err != nil {
		return err
	}

	// Compare WidthBefore and WidthAfter fields
	if err := compareTableWidth(a.WidthBefore, b.WidthBefore); err != nil {
		return err
	}
	if err := compareTableWidth(a.WidthAfter, b.WidthAfter); err != nil {
		return err
	}

	// Compare CantSplit fields
	if err := internal.ComparePtr("CantSplit", a.CantSplit, b.CantSplit); err != nil {
		return err
	}

	// Compare Height fields
	if err := compareTableRowHeight(a.Height, b.Height); err != nil {
		return err
	}

	// Compare Header fields
	if err := internal.ComparePtr("Header", a.Header, b.Header); err != nil {
		return err
	}

	// Compare CellSpacing fields
	if err := compareTableWidth(a.CellSpacing, b.CellSpacing); err != nil {
		return err
	}

	// Compare JC fields
	if err := internal.ComparePtr("JC", a.JC, b.JC); err != nil {
		return err
	}

	// Compare Hidden fields
	if err := internal.ComparePtr("Hidden", a.Hidden, b.Hidden); err != nil {
		return err
	}

	// Compare Ins fields
	if err := internal.ComparePtr("Ins", a.Ins, b.Ins); err != nil {
		return err
	}

	// Compare Del fields
	if err := internal.ComparePtr("Del", a.Del, b.Del); err != nil {
		return err
	}

	// Compare Change fields
	if err := internal.ComparePtr("Change", a.Change, b.Change); err != nil {
		return err
	}

	return nil
}

// Helper function to compare TableWidth structs
func compareTableWidth(a, b *TableWidth) error {
	if a == nil && b == nil {
		return nil
	}
	if a == nil || b == nil {
		return errors.New("TableWidth mismatch")
	}

	if err := internal.ComparePtr("TableWidth", a.Width, b.Width); err != nil {
		return err
	}

	if err := internal.ComparePtr("TableWidthType", a.WidthType, b.WidthType); err != nil {
		return err
	}
	return nil
}

// Helper function to compare TableRowHeight structs
func compareTableRowHeight(a, b *TableRowHeight) error {
	if a == nil && b == nil {
		return nil
	}
	if a == nil || b == nil {
		return errors.New("TableRowHeight mismatch")
	}
	if err := internal.ComparePtr("TableRowHeight Value", a.Val, b.Val); err != nil {
		return err
	}

	if err := internal.ComparePtr("TableRowHeight HRule", a.HRule, b.HRule); err != nil {
		return err
	}
	return nil
}
