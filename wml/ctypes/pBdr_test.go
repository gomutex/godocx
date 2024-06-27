package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/internal"
)

func TestParaBorder_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    ParaBorder
		expected string
	}{
		{
			name: "All fields set",
			input: ParaBorder{
				Top:     &Border{Val: "single", Color: internal.ToPtr("auto")},
				Left:    &Border{Val: "dashed", Color: internal.ToPtr("FF0000")},
				Right:   &Border{Val: "double", Color: internal.ToPtr("00FF00")},
				Bottom:  &Border{Val: "dotted", Color: internal.ToPtr("0000FF")},
				Between: &Border{Val: "wave", Color: internal.ToPtr("123456")},
				Bar:     &Border{Val: "thick", Color: internal.ToPtr("654321")},
			},
			expected: `<w:pBdr><w:top w:val="single" w:color="auto"></w:top><w:left w:val="dashed" w:color="FF0000"></w:left><w:right w:val="double" w:color="00FF00"></w:right><w:bottom w:val="dotted" w:color="0000FF"></w:bottom><w:between w:val="wave" w:color="123456"></w:between><w:bar w:val="thick" w:color="654321"></w:bar></w:pBdr>`,
		},
		{
			name: "Some fields set",
			input: ParaBorder{
				Top:    &Border{Val: "single", Color: internal.ToPtr("auto")},
				Bottom: &Border{Val: "dotted", Color: internal.ToPtr("0000FF")},
			},
			expected: `<w:pBdr><w:top w:val="single" w:color="auto"></w:top><w:bottom w:val="dotted" w:color="0000FF"></w:bottom></w:pBdr>`,
		},
		{
			name:     "No fields set",
			input:    ParaBorder{},
			expected: `<w:pBdr></w:pBdr>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)

			start := xml.StartElement{Name: xml.Name{Local: "w:pBdr"}}
			if err := tt.input.MarshalXML(encoder, start); err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			encoder.Flush()

			got := strings.TrimSpace(result.String())
			if got != tt.expected {
				t.Errorf("Expected XML:\n%s\nGot:\n%s", tt.expected, got)
			}
		})
	}
}

func TestParaBorder_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected ParaBorder
	}{
		{
			name: "All fields set",
			inputXML: `<w:pBdr>
						<w:top w:val="single" w:color="auto"></w:top>
						<w:left w:val="dashed" w:color="FF0000"></w:left>
						<w:right w:val="double" w:color="00FF00"></w:right>
						<w:bottom w:val="dotted" w:color="0000FF"></w:bottom>
						<w:between w:val="wave" w:color="123456"></w:between>
						<w:bar w:val="thick" w:color="654321"></w:bar>
					  </w:pBdr>`,
			expected: ParaBorder{
				Top:     &Border{Val: "single", Color: internal.ToPtr("auto")},
				Left:    &Border{Val: "dashed", Color: internal.ToPtr("FF0000")},
				Right:   &Border{Val: "double", Color: internal.ToPtr("00FF00")},
				Bottom:  &Border{Val: "dotted", Color: internal.ToPtr("0000FF")},
				Between: &Border{Val: "wave", Color: internal.ToPtr("123456")},
				Bar:     &Border{Val: "thick", Color: internal.ToPtr("654321")},
			},
		},
		{
			name: "Some fields set",
			inputXML: `<w:pBdr>
						<w:top w:val="single" w:color="auto"></w:top>
						<w:bottom w:val="dotted" w:color="0000FF"></w:bottom>
					  </w:pBdr>`,
			expected: ParaBorder{
				Top:    &Border{Val: "single", Color: internal.ToPtr("auto")},
				Bottom: &Border{Val: "dotted", Color: internal.ToPtr("0000FF")},
			},
		},
		{
			name: "No fields set",
			inputXML: `<w:pBdr>
					  </w:pBdr>`,
			expected: ParaBorder{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result ParaBorder
			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			// Compare unmarshaled result with expected
			if !paraBorderEqual(result, tt.expected) {
				t.Errorf("Unmarshaled ParaBorder does not match expected:\nExpected: %+v\nGot: %+v", tt.expected, result)
			}
		})
	}
}

// Helper function to compare ParaBorder structs
func paraBorderEqual(p1, p2 ParaBorder) bool {
	if (p1.Top == nil && p2.Top != nil) || (p1.Top != nil && p2.Top == nil) {
		return false
	}
	if p1.Top != nil && p2.Top != nil {
		if p1.Top.Val != p2.Top.Val || *p1.Top.Color != *p2.Top.Color {
			return false
		}
	}

	if (p1.Left == nil && p2.Left != nil) || (p1.Left != nil && p2.Left == nil) {
		return false
	}
	if p1.Left != nil && p2.Left != nil {
		if p1.Left.Val != p2.Left.Val || *p1.Left.Color != *p2.Left.Color {
			return false
		}
	}

	if (p1.Right == nil && p2.Right != nil) || (p1.Right != nil && p2.Right == nil) {
		return false
	}
	if p1.Right != nil && p2.Right != nil {
		if p1.Right.Val != p2.Right.Val || *p1.Right.Color != *p2.Right.Color {
			return false
		}
	}

	if (p1.Bottom == nil && p2.Bottom != nil) || (p1.Bottom != nil && p2.Bottom == nil) {
		return false
	}
	if p1.Bottom != nil && p2.Bottom != nil {
		if p1.Bottom.Val != p2.Bottom.Val || *p1.Bottom.Color != *p2.Bottom.Color {
			return false
		}
	}

	if (p1.Between == nil && p2.Between != nil) || (p1.Between != nil && p2.Between == nil) {
		return false
	}
	if p1.Between != nil && p2.Between != nil {
		if p1.Between.Val != p2.Between.Val || *p1.Between.Color != *p2.Between.Color {
			return false
		}
	}

	if (p1.Bar == nil && p2.Bar != nil) || (p1.Bar != nil && p2.Bar == nil) {
		return false
	}
	if p1.Bar != nil && p2.Bar != nil {
		if p1.Bar.Val != p2.Bar.Val || *p1.Bar.Color != *p2.Bar.Color {
			return false
		}
	}

	return true
}
