package table

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestOverlap_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Overlap
		expected string
	}{
		{
			name:     "Test with Overlap Value `never`",
			input:    Overlap{Val: stypes.TblOverlapNever},
			expected: `<w:tblOverlap w:val="never"></w:tblOverlap>`,
		},
		{
			name:     "Test with Overlap Value `default`",
			input:    Overlap{Val: stypes.TblOverlapOverlap},
			expected: `<w:tblOverlap w:val="overlap"></w:tblOverlap>`,
		},
		{
			name:     "Test with Overlap Value `overlap`",
			input:    Overlap{Val: stypes.TblOverlapOverlap},
			expected: `<w:tblOverlap w:val="overlap"></w:tblOverlap>`,
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

func TestOverlap_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name       string
		inputXML   string
		expected   Overlap
		expectFail bool // Whether unmarshalling is expected to fail
	}{
		{
			name:     "Test with Overlap Value `never`",
			inputXML: `<w:tblOverlap w:val="never"></w:tblOverlap>`,
			expected: Overlap{Val: stypes.TblOverlapNever},
		},
		{
			name:     "Test with Overlap Value `default`",
			inputXML: `<w:tblOverlap w:val="overlap"></w:tblOverlap>`,
			expected: Overlap{Val: stypes.TblOverlapOverlap},
		},
		{
			name:     "Test with Overlap Value `overlap`",
			inputXML: `<w:tblOverlap w:val="overlap"></w:tblOverlap>`,
			expected: Overlap{Val: stypes.TblOverlapOverlap},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Overlap
			err := xml.Unmarshal([]byte(tt.inputXML), &result)

			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if result.Val != tt.expected.Val {
				t.Errorf("Expected Overlap value %s but got %s", tt.expected.Val, result.Val)
			}

		})
	}
}
