package docxrun

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/simpletypes"
)

// Helper functions to create pointers
func intPtr(i int) *int {
	return &i
}

func onOffPtr(value simpletypes.OnOff) *simpletypes.OnOff {
	return &value
}

func combineBracketsPtr(value simpletypes.CombineBrackets) *simpletypes.CombineBrackets {
	return &value
}

func TestEALayout_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		layout   EALayout
		expected string
	}{
		{
			name: "All attributes set",
			layout: EALayout{
				ID:           intPtr(1),
				Combine:      onOffPtr(simpletypes.OnOffOn),
				CombineBrkts: combineBracketsPtr(simpletypes.CombineBracketsRound),
				Vert:         onOffPtr(simpletypes.OnOffOff),
				VertCompress: onOffPtr(simpletypes.OnOffOn),
			},
			expected: `<w:eastAsianLayout w:id="1" w:combine="on" w:combineBrackets="round" w:vert="off" w:vertCompress="on"></w:eastAsianLayout>`,
		},
		{
			name: "Only ID set",
			layout: EALayout{
				ID: intPtr(2),
			},
			expected: `<w:eastAsianLayout w:id="2"></w:eastAsianLayout>`,
		},
		{
			name: "Only Combine set",
			layout: EALayout{
				Combine: onOffPtr(simpletypes.OnOffOn),
			},
			expected: `<w:eastAsianLayout w:combine="on"></w:eastAsianLayout>`,
		},
		{
			name: "Only CombineBrkts set",
			layout: EALayout{
				CombineBrkts: combineBracketsPtr(simpletypes.CombineBracketsSquare),
			},
			expected: `<w:eastAsianLayout w:combineBrackets="square"></w:eastAsianLayout>`,
		},
		{
			name: "Only Vert set",
			layout: EALayout{
				Vert: onOffPtr(simpletypes.OnOffOff),
			},
			expected: `<w:eastAsianLayout w:vert="off"></w:eastAsianLayout>`,
		},
		{
			name: "Only VertCompress set",
			layout: EALayout{
				VertCompress: onOffPtr(simpletypes.OnOffOn),
			},
			expected: `<w:eastAsianLayout w:vertCompress="on"></w:eastAsianLayout>`,
		},
		{
			name:     "No attributes set",
			layout:   EALayout{},
			expected: `<w:eastAsianLayout></w:eastAsianLayout>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			e := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:eastAsianLayout"}}
			err := tt.layout.MarshalXML(e, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}
			e.Flush()

			if result.String() != tt.expected {
				t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expected, result.String())
			}
		})
	}
}

func TestEALayout_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected EALayout
	}{
		{
			name:     "All attributes set",
			inputXML: `<w:eastAsianLayout w:id="1" w:combine="on" w:combineBrackets="round" w:vert="off" w:vertCompress="on"></w:eastAsianLayout>`,
			expected: EALayout{
				ID:           intPtr(1),
				Combine:      onOffPtr(simpletypes.OnOffOn),
				CombineBrkts: combineBracketsPtr(simpletypes.CombineBracketsRound),
				Vert:         onOffPtr(simpletypes.OnOffOff),
				VertCompress: onOffPtr(simpletypes.OnOffOn),
			},
		},
		{
			name:     "Only ID set",
			inputXML: `<w:eastAsianLayout w:id="2"></w:eastAsianLayout>`,
			expected: EALayout{
				ID: intPtr(2),
			},
		},
		{
			name:     "Only Combine set",
			inputXML: `<w:eastAsianLayout w:combine="on"></w:eastAsianLayout>`,
			expected: EALayout{
				Combine: onOffPtr(simpletypes.OnOffOn),
			},
		},
		{
			name:     "Only CombineBrkts set",
			inputXML: `<w:eastAsianLayout w:combineBrackets="square"></w:eastAsianLayout>`,
			expected: EALayout{
				CombineBrkts: combineBracketsPtr(simpletypes.CombineBracketsSquare),
			},
		},
		{
			name:     "Only Vert set",
			inputXML: `<w:eastAsianLayout w:vert="off"></w:eastAsianLayout>`,
			expected: EALayout{
				Vert: onOffPtr(simpletypes.OnOffOff),
			},
		},
		{
			name:     "Only VertCompress set",
			inputXML: `<w:eastAsianLayout w:vertCompress="on"></w:eastAsianLayout>`,
			expected: EALayout{
				VertCompress: onOffPtr(simpletypes.OnOffOn),
			},
		},
		{
			name:     "No attributes set",
			inputXML: `<w:eastAsianLayout></w:eastAsianLayout>`,
			expected: EALayout{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var layout EALayout
			err := xml.Unmarshal([]byte(tt.inputXML), &layout)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if !compareEALayout(layout, tt.expected) {
				t.Errorf("Expected %+v but got %+v", tt.expected, layout)
			}
		})
	}
}

func compareEALayout(a, b EALayout) bool {
	if (a.ID == nil && b.ID != nil) || (a.ID != nil && b.ID == nil) {
		return false
	}
	if a.ID != nil && *a.ID != *b.ID {
		return false
	}
	if (a.Combine == nil && b.Combine != nil) || (a.Combine != nil && b.Combine == nil) {
		return false
	}
	if a.Combine != nil && *a.Combine != *b.Combine {
		return false
	}
	if (a.CombineBrkts == nil && b.CombineBrkts != nil) || (a.CombineBrkts != nil && b.CombineBrkts == nil) {
		return false
	}
	if a.CombineBrkts != nil && *a.CombineBrkts != *b.CombineBrkts {
		return false
	}
	if (a.Vert == nil && b.Vert != nil) || (a.Vert != nil && b.Vert == nil) {
		return false
	}
	if a.Vert != nil && *a.Vert != *b.Vert {
		return false
	}
	if (a.VertCompress == nil && b.VertCompress != nil) || (a.VertCompress != nil && b.VertCompress == nil) {
		return false
	}
	if a.VertCompress != nil && *a.VertCompress != *b.VertCompress {
		return false
	}
	return true
}
