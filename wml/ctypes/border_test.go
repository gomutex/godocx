package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestBorder_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Border
		expected string
	}{
		{
			name: "With all attributes",
			input: Border{
				Val:        stypes.BorderStyleSingle,
				Color:      StringPtr("FF0000"),
				ThemeColor: themeColorPointer(stypes.ThemeColorAccent1),
				ThemeTint:  StringPtr("500"),
				ThemeShade: StringPtr("200"),
				Space:      StringPtr("0"),
				Shadow:     OnOffPtr(stypes.OnOffTrue),
				Frame:      OnOffPtr(stypes.OnOffTrue),
			},
			expected: `<w:bdr w:val="single" w:color="FF0000" w:themeColor="accent1" w:themeTint="500" w:themeShade="200" w:space="0" w:shadow="true" w:frame="true"></w:bdr>`,
		},
		{
			name: "Without optional attributes",
			input: Border{
				Val: stypes.BorderStyleDouble,
			},
			expected: `<w:bdr w:val="double"></w:bdr>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:bdr"}}

			err := tt.input.MarshalXML(encoder, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			encoder.Flush()

			if result.String() != tt.expected {
				t.Errorf("Expected XML:\n%s\nGot:\n%s", tt.expected, result.String())
			}
		})
	}
}

func TestBorder_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected Border
	}{
		{
			name: "With all attributes",
			inputXML: `<w:bdr w:val="single" w:color="FF0000" w:themeColor="accent1" w:themeTint="500" ` +
				`w:themeShade="200" w:space="0" w:shadow="true" w:frame="true"></w:bdr>`,
			expected: Border{
				Val:        stypes.BorderStyleSingle,
				Color:      StringPtr("FF0000"),
				ThemeColor: themeColorPointer(stypes.ThemeColorAccent1),
				ThemeTint:  StringPtr("500"),
				ThemeShade: StringPtr("200"),
				Space:      StringPtr("0"),
				Shadow:     OnOffPtr(stypes.OnOffTrue),
				Frame:      OnOffPtr(stypes.OnOffTrue),
			},
		},
		{
			name:     "Without optional attributes",
			inputXML: `<w:bdr w:val="double"></w:bdr>`,
			expected: Border{
				Val: stypes.BorderStyleDouble,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Border

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			// Compare each field individually due to pointer comparisons
			if result.Val != tt.expected.Val {
				t.Errorf("Expected Val %s but got %s", tt.expected.Val, result.Val)
			}
			if tt.expected.Color != nil {
				if result.Color == nil {
					t.Errorf("Expected Color %s but got nil", *tt.expected.Color)
				} else if *result.Color != *tt.expected.Color {
					t.Errorf("Expected Color %s but got %s", *tt.expected.Color, *result.Color)
				}
			} else if result.Color != nil {
				t.Errorf("Expected nil but got %s", *result.Color)
			}
			if tt.expected.ThemeColor != nil {
				if result.ThemeColor == nil {
					t.Errorf("Expected ThemeColor %s but got nil", *tt.expected.ThemeColor)
				} else if *result.ThemeColor != *tt.expected.ThemeColor {
					t.Errorf("Expected ThemeColor %s but got %s", *tt.expected.ThemeColor, *result.ThemeColor)
				}
			} else if result.ThemeColor != nil {
				t.Errorf("Expected nil but got %s", *result.ThemeColor)
			}
			if tt.expected.ThemeTint != nil {
				if result.ThemeTint == nil {
					t.Errorf("Expected ThemeTint %s but got nil", *tt.expected.ThemeTint)
				} else if *result.ThemeTint != *tt.expected.ThemeTint {
					t.Errorf("Expected ThemeTint %s but got %s", *tt.expected.ThemeTint, *result.ThemeTint)
				}
			} else if result.ThemeTint != nil {
				t.Errorf("Expected nil but got %s", *result.ThemeTint)
			}
			if tt.expected.ThemeShade != nil {
				if result.ThemeShade == nil {
					t.Errorf("Expected ThemeShade %s but got nil", *tt.expected.ThemeShade)
				} else if *result.ThemeShade != *tt.expected.ThemeShade {
					t.Errorf("Expected ThemeShade %s but got %s", *tt.expected.ThemeShade, *result.ThemeShade)
				}
			} else if result.ThemeShade != nil {
				t.Errorf("Expected nil but got %s", *result.ThemeShade)
			}
			if tt.expected.Space != nil {
				if result.Space == nil {
					t.Errorf("Expected Space %s but got nil", *tt.expected.Space)
				} else if *result.Space != *tt.expected.Space {
					t.Errorf("Expected Space %s but got %s", *tt.expected.Space, *result.Space)
				}
			} else if result.Space != nil {
				t.Errorf("Expected nil but got %s", *result.Space)
			}
			if tt.expected.Shadow != nil {
				if result.Shadow == nil {
					t.Errorf("Expected Shadow %s but got nil", *tt.expected.Shadow)
				} else if *result.Shadow != *tt.expected.Shadow {
					t.Errorf("Expected Shadow %s but got %s", *tt.expected.Shadow, *result.Shadow)
				}
			} else if result.Shadow != nil {
				t.Errorf("Expected nil but got %s", *result.Shadow)
			}
			if tt.expected.Frame != nil {
				if result.Frame == nil {
					t.Errorf("Expected Frame %s but got nil", *tt.expected.Frame)
				} else if *result.Frame != *tt.expected.Frame {
					t.Errorf("Expected Frame %s but got %s", *tt.expected.Frame, *result.Frame)
				}
			} else if result.Frame != nil {
				t.Errorf("Expected nil but got %s", *result.Frame)
			}
		})
	}
}

func StringPtr(s string) *string {
	return &s
}

func OnOffPtr(o stypes.OnOff) *stypes.OnOff {
	return &o
}

func themeColorPointer(t stypes.ThemeColor) *stypes.ThemeColor {
	return &t
}
