package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/internal"
)

func TestNumProp_MarshalXML(t *testing.T) {
	date := "2023-06-18T12:34:56Z"
	tests := []struct {
		name     string
		input    NumProp
		expected string
	}{
		{
			name: "With all attributes",
			input: NumProp{
				ILvl:  &DecimalNum{Val: 1},
				NumID: &DecimalNum{Val: 42},
				NumChange: &TrackChangeNum{
					ID:       123,
					Author:   "John Doe",
					Date:     &date,
					Original: internal.ToPtr("original"),
				},
				Ins: &TrackChange{
					ID:     124,
					Author: "Jane Doe",
					Date:   &date,
				},
			},
			expected: `<w:numPr><w:ilvl w:val="1"></w:ilvl><w:numId w:val="42"></w:numId><w:numberingChange w:id="123" w:author="John Doe" w:date="2023-06-18T12:34:56Z" w:original="original"></w:numberingChange><w:ins w:id="124" w:author="Jane Doe" w:date="2023-06-18T12:34:56Z"></w:ins></w:numPr>`,
		},
		{
			name: "Without optional attributes",
			input: NumProp{
				ILvl:  &DecimalNum{Val: 1},
				NumID: &DecimalNum{Val: 42},
			},
			expected: `<w:numPr><w:ilvl w:val="1"></w:ilvl><w:numId w:val="42"></w:numId></w:numPr>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{}

			err := tt.input.MarshalXML(encoder, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			// Finalize encoding
			encoder.Flush()

			if result.String() != tt.expected {
				t.Errorf("Expected XML:\n%s\nGot:\n%s", tt.expected, result.String())
			}
		})
	}
}

func TestNumProp_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected NumProp
	}{
		{
			name: "All fields present",
			inputXML: `<w:numPr>
				<w:ilvl w:val="2"/>
				<w:numId w:val="5"/>
				<w:numberingChange w:id="1" w:author="author" w:date="2024-06-19T12:00:00Z"/>
				<w:ins w:id="2" w:author="author" w:date="2024-06-19T12:00:00Z"/>
			</w:numPr>`,
			expected: NumProp{
				ILvl:  &DecimalNum{Val: 2},
				NumID: &DecimalNum{Val: 5},
				NumChange: &TrackChangeNum{
					ID:     1,
					Author: "author",
					Date:   internal.ToPtr("2024-06-19T12:00:00Z"),
				},
				Ins: &TrackChange{
					ID:     2,
					Author: "author",
					Date:   internal.ToPtr("2024-06-19T12:00:00Z"),
				},
			},
		},
		{
			name: "Some fields missing",
			inputXML: `<w:numPr>
				<w:ilvl w:val="1"/>
			</w:numPr>`,
			expected: NumProp{
				ILvl: &DecimalNum{Val: 1},
			},
		},
		{
			name:     "Empty struct",
			inputXML: `<w:numPr></w:numPr>`,
			expected: NumProp{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result NumProp
			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			// Validate ILvl field
			if tt.expected.ILvl != nil {
				if result.ILvl == nil {
					t.Errorf("Expected ILvl %d but got nil", tt.expected.ILvl.Val)
				} else if *result.ILvl != *tt.expected.ILvl {
					t.Errorf("Expected ILvl %d but got %d", tt.expected.ILvl.Val, result.ILvl.Val)
				}
			} else if result.ILvl != nil {
				t.Errorf("Expected ILvl nil but got %d", result.ILvl.Val)
			}

			// Validate NumID field
			if tt.expected.NumID != nil {
				if result.NumID == nil {
					t.Errorf("Expected NumID %d but got nil", tt.expected.NumID.Val)
				} else if *result.NumID != *tt.expected.NumID {
					t.Errorf("Expected NumID %d but got %d", tt.expected.NumID.Val, result.NumID.Val)
				}
			} else if result.NumID != nil {
				t.Errorf("Expected NumID nil but got %d", result.NumID.Val)
			}

			// Validate NumChange field
			if tt.expected.NumChange != nil {
				if result.NumChange == nil {
					t.Errorf("Expected NumChange but got nil")
				} else if !compareTrackChangeNum(*tt.expected.NumChange, *result.NumChange) {
					t.Errorf("Expected NumChange %+v but got %+v", *tt.expected.NumChange, *result.NumChange)
				}
			} else if result.NumChange != nil {
				t.Errorf("Expected NumChange nil but got %+v", *result.NumChange)
			}

			// Validate Ins field
			if tt.expected.Ins != nil {
				if result.Ins == nil {
					t.Errorf("Expected Ins but got nil")
				} else if !compareTrackChange(*tt.expected.Ins, *result.Ins) {
					t.Errorf("Expected Ins %+v but got %+v", *tt.expected.Ins, *result.Ins)
				}
			} else if result.Ins != nil {
				t.Errorf("Expected Ins nil but got %+v", *result.Ins)
			}
		})
	}
}

func compareTrackChangeNum(a, b TrackChangeNum) bool {
	return a.ID == b.ID && a.Author == b.Author && *a.Date == *b.Date
}

func compareTrackChange(a, b TrackChange) bool {
	return a.ID == b.ID && a.Author == b.Author && *a.Date == *b.Date
}
