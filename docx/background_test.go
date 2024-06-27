package docx

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestBackground_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Background
		expected string
	}{
		{
			name: "With all attributes",
			input: Background{
				Color:      StringPtr("FFFFFF"),
				ThemeColor: ThemeColorPtr(stypes.ThemeColorAccent1),
				ThemeTint:  StringPtr("500"),
				ThemeShade: StringPtr("200"),
			},
			expected: `<w:background w:color="FFFFFF" w:themeColor="accent1" w:themeTint="500" w:themeShade="200"></w:background>`,
		},
		{
			name: "Without optional attributes",
			input: Background{
				Color: StringPtr("000000"),
			},
			expected: `<w:background w:color="000000"></w:background>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:background"}}

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

func TestBackground_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected Background
	}{
		{
			name:     "With all attributes",
			inputXML: `<w:background w:color="FFFFFF" w:themeColor="accent1" w:themeTint="500" w:themeShade="200"></w:background>`,
			expected: Background{
				Color:      StringPtr("FFFFFF"),
				ThemeColor: ThemeColorPtr(stypes.ThemeColorAccent1),
				ThemeTint:  StringPtr("500"),
				ThemeShade: StringPtr("200"),
			},
		},
		{
			name:     "Without optional attributes",
			inputXML: `<w:background w:color="000000"></w:background>`,
			expected: Background{
				Color: StringPtr("000000"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Background

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
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
		})
	}
}

func StringPtr(s string) *string {
	return &s
}

func ThemeColorPtr(t stypes.ThemeColor) *stypes.ThemeColor {
	return &t
}
