package docxpara

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/internal"
)

func TestPPrChange_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    PPrChange
		expected string
	}{
		{
			name: "With all attributes",
			input: PPrChange{
				ID:       123,
				Author:   "John Doe",
				Date:     internal.ToPtr("2024-06-19"),
				ParaProp: &ParagraphProp{
					// Initialize ParagraphProp fields here if needed
				},
			},
			expected: `<w:pPrChange id="123" author="John Doe" date="2024-06-19"><w:pPr></w:pPr></w:pPrChange>`,
		},
		{
			name: "Without date attribute",
			input: PPrChange{
				ID:       456,
				Author:   "Jane Smith",
				ParaProp: &ParagraphProp{
					// Initialize ParagraphProp fields here if needed
				},
			},
			expected: `<w:pPrChange id="456" author="Jane Smith"><w:pPr></w:pPr></w:pPrChange>`,
		},
		{
			name: "Without paraProp",
			input: PPrChange{
				ID:     789,
				Author: "Alice Brown",
				Date:   internal.ToPtr("2024-06-20"),
			},
			expected: `<w:pPrChange id="789" author="Alice Brown" date="2024-06-20"></w:pPrChange>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)

			start := xml.StartElement{Name: xml.Name{Local: "w:pPrChange"}}
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
