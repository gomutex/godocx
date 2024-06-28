package ctypes

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"

	"github.com/gomutex/godocx/internal"
)

func TestSym_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Sym
		expected string
	}{
		{
			name:     "Test with Font only",
			input:    Sym{Font: internal.ToPtr("Arial")},
			expected: `<w:sym w:font="Arial"></w:sym>`,
		},
		{
			name:     "Test with Char only",
			input:    Sym{Char: internal.ToPtr("F0")},
			expected: `<w:sym w:char="F0"></w:sym>`,
		},
		{
			name:     "Test with Font and Char",
			input:    Sym{Font: internal.ToPtr("Times New Roman"), Char: internal.ToPtr("03")},
			expected: `<w:sym w:font="Times New Roman" w:char="03"></w:sym>`,
		},
		{
			name:     "Test with nil values",
			input:    Sym{},
			expected: `<w:sym></w:sym>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:sym"}}

			err := tt.input.MarshalXML(encoder, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			if err = encoder.Flush(); err != nil {
				t.Fatalf("Error flushing XML encoder: %v", err)
			}

			if result.String() != tt.expected {
				t.Errorf("Expected XML:\n%s\nGot:\n%s", tt.expected, result.String())
			}
		})
	}
}

func TestSym_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name       string
		inputXML   string
		expected   Sym
		expectFail bool // Whether unmarshalling is expected to fail
	}{
		{
			name:     "Test with Font attribute",
			inputXML: `<w:sym w:font="Verdana"></w:sym>`,
			expected: Sym{Font: internal.ToPtr("Verdana")},
		},
		{
			name:     "Test with Char attribute",
			inputXML: `<w:sym w:char="0E"></w:sym>`,
			expected: Sym{Char: internal.ToPtr("0E")},
		},
		{
			name:     "Test with Font and Char attributes",
			inputXML: `<w:sym w:font="Arial" w:char="F2"></w:sym>`,
			expected: Sym{Font: internal.ToPtr("Arial"), Char: internal.ToPtr("F2")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Sym
			err := xml.Unmarshal([]byte(tt.inputXML), &result)

			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected Sym %+v but got %+v", tt.expected, result)
			}
		})
	}
}
