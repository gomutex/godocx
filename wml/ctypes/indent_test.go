package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/internal"
)

func TestInd_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Indent
		expected string
	}{
		{
			name: "With all attributes",
			input: Indent{
				Left:           internal.ToPtr(720),
				LeftChars:      internal.ToPtr(2),
				Right:          internal.ToPtr(360),
				RightChars:     internal.ToPtr(1),
				Hanging:        internal.ToPtr(uint64(360)),
				HangingChars:   internal.ToPtr(1),
				FirstLine:      internal.ToPtr(uint64(720)),
				FirstLineChars: internal.ToPtr(2),
			},
			expected: `<w:ind w:left="720" w:leftChars="2" w:right="360" w:rightChars="1" w:hanging="360" w:hangingChars="1" w:firstLine="720" w:firstLineChars="2"></w:ind>`,
		},
		{
			name:     "Without attributes",
			input:    Indent{},
			expected: `<w:ind></w:ind>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)

			start := xml.StartElement{Name: xml.Name{Local: "w:ind"}}
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

func TestInd_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected Indent
	}{
		{
			name:     "With all attributes",
			inputXML: `<w:ind w:left="720" w:leftChars="2" w:right="360" w:rightChars="1" w:hanging="360" w:hangingChars="1" w:firstLine="720" w:firstLineChars="2"></w:ind>`,
			expected: Indent{
				Left:           internal.ToPtr(720),
				LeftChars:      internal.ToPtr(2),
				Right:          internal.ToPtr(360),
				RightChars:     internal.ToPtr(1),
				Hanging:        internal.ToPtr(uint64(360)),
				HangingChars:   internal.ToPtr(1),
				FirstLine:      internal.ToPtr(uint64(720)),
				FirstLineChars: internal.ToPtr(2),
			},
		},
		{
			name:     "Without attributes",
			inputXML: `<w:ind></w:ind>`,
			expected: Indent{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ind Indent
			err := xml.Unmarshal([]byte(tt.inputXML), &ind)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}
			//
			err = isEqualInd(tt.expected, ind)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

// Utility function to compare two Indent structs
func isEqualInd(a, b Indent) error {
	if err := internal.ComparePtr("Left", a.Left, b.Left); err != nil {
		return err
	}

	if err := internal.ComparePtr("LeftChars", a.LeftChars, b.LeftChars); err != nil {
		return err
	}

	if err := internal.ComparePtr("Right", a.Right, b.Right); err != nil {
		return err
	}

	if err := internal.ComparePtr("RightChars", a.RightChars, b.RightChars); err != nil {
		return err
	}

	if err := internal.ComparePtr("Hanging", a.Hanging, b.Hanging); err != nil {
		return err
	}

	if err := internal.ComparePtr("HangingChars", a.HangingChars, b.HangingChars); err != nil {
		return err
	}

	if err := internal.ComparePtr("FirstLine", a.FirstLine, b.FirstLine); err != nil {
		return err
	}

	if err := internal.ComparePtr("FirstLineChars", a.FirstLineChars, b.FirstLineChars); err != nil {
		return err
	}

	return nil
}
