package docxrun

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/simpletypes"
)

func TestEm_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		em       Em
		expected string
	}{
		{
			name:     "Valid Value - none",
			em:       Em{Val: simpletypes.EmNone},
			expected: `<w:em w:val="none"></w:em>`,
		},
		{
			name:     "Valid Value - dot",
			em:       Em{Val: simpletypes.EmDot},
			expected: `<w:em w:val="dot"></w:em>`,
		},
		{
			name:     "Valid Value - comma",
			em:       Em{Val: simpletypes.EmComma},
			expected: `<w:em w:val="comma"></w:em>`,
		},
		{
			name:     "Valid Value - circle",
			em:       Em{Val: simpletypes.EmCircle},
			expected: `<w:em w:val="circle"></w:em>`,
		},
		{
			name:     "Valid Value - underDot",
			em:       Em{Val: simpletypes.EmUnderDot},
			expected: `<w:em w:val="underDot"></w:em>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:effect"}}

			err := tt.em.MarshalXML(encoder, start)
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

func TestEm_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected Em
	}{
		{
			name:     "Valid Value - none",
			inputXML: `<w:em w:val="none"></w:em>`,
			expected: Em{Val: simpletypes.EmNone},
		},
		{
			name:     "Valid Value - dot",
			inputXML: `<w:em w:val="dot"></w:em>`,
			expected: Em{Val: simpletypes.EmDot},
		},
		{
			name:     "Valid Value - comma",
			inputXML: `<w:em w:val="comma"></w:em>`,
			expected: Em{Val: simpletypes.EmComma},
		},
		{
			name:     "Valid Value - circle",
			inputXML: `<w:em w:val="circle"></w:em>`,
			expected: Em{Val: simpletypes.EmCircle},
		},
		{
			name:     "Valid Value - underDot",
			inputXML: `<w:em w:val="underDot"></w:em>`,
			expected: Em{Val: simpletypes.EmUnderDot},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var em Em
			err := xml.Unmarshal([]byte(tt.inputXML), &em)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if em.Val != tt.expected.Val {
				t.Errorf("Expected val %s but got %s", tt.expected.Val, em.Val)
			}
		})
	}
}
