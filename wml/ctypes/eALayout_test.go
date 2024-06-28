package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/internal"
	"github.com/gomutex/godocx/wml/stypes"
)

func TestEALayout_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		layout   EALayout
		expected string
	}{
		{
			name: "All attributes set",
			layout: EALayout{
				ID:           internal.ToPtr(1),
				Combine:      internal.ToPtr(stypes.OnOffOn),
				CombineBrkts: internal.ToPtr(stypes.CombineBracketsRound),
				Vert:         internal.ToPtr(stypes.OnOffOff),
				VertCompress: internal.ToPtr(stypes.OnOffOn),
			},
			expected: `<w:eastAsianLayout w:id="1" w:combine="on" w:combineBrackets="round" w:vert="off" w:vertCompress="on"></w:eastAsianLayout>`,
		},
		{
			name: "Only ID set",
			layout: EALayout{
				ID: internal.ToPtr(2),
			},
			expected: `<w:eastAsianLayout w:id="2"></w:eastAsianLayout>`,
		},
		{
			name: "Only Combine set",
			layout: EALayout{
				Combine: internal.ToPtr(stypes.OnOffOn),
			},
			expected: `<w:eastAsianLayout w:combine="on"></w:eastAsianLayout>`,
		},
		{
			name: "Only CombineBrkts set",
			layout: EALayout{
				CombineBrkts: internal.ToPtr(stypes.CombineBracketsSquare),
			},
			expected: `<w:eastAsianLayout w:combineBrackets="square"></w:eastAsianLayout>`,
		},
		{
			name: "Only Vert set",
			layout: EALayout{
				Vert: internal.ToPtr(stypes.OnOffOff),
			},
			expected: `<w:eastAsianLayout w:vert="off"></w:eastAsianLayout>`,
		},
		{
			name: "Only VertCompress set",
			layout: EALayout{
				VertCompress: internal.ToPtr(stypes.OnOffOn),
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
			err := tt.layout.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:eastAsianLayout"}})
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
				ID:           internal.ToPtr(1),
				Combine:      internal.ToPtr(stypes.OnOffOn),
				CombineBrkts: internal.ToPtr(stypes.CombineBracketsRound),
				Vert:         internal.ToPtr(stypes.OnOffOff),
				VertCompress: internal.ToPtr(stypes.OnOffOn),
			},
		},
		{
			name:     "Only ID set",
			inputXML: `<w:eastAsianLayout w:id="2"></w:eastAsianLayout>`,
			expected: EALayout{
				ID: internal.ToPtr(2),
			},
		},
		{
			name:     "Only Combine set",
			inputXML: `<w:eastAsianLayout w:combine="on"></w:eastAsianLayout>`,
			expected: EALayout{
				Combine: internal.ToPtr(stypes.OnOffOn),
			},
		},
		{
			name:     "Only CombineBrkts set",
			inputXML: `<w:eastAsianLayout w:combineBrackets="square"></w:eastAsianLayout>`,
			expected: EALayout{
				CombineBrkts: internal.ToPtr(stypes.CombineBracketsSquare),
			},
		},
		{
			name:     "Only Vert set",
			inputXML: `<w:eastAsianLayout w:vert="off"></w:eastAsianLayout>`,
			expected: EALayout{
				Vert: internal.ToPtr(stypes.OnOffOff),
			},
		},
		{
			name:     "Only VertCompress set",
			inputXML: `<w:eastAsianLayout w:vertCompress="on"></w:eastAsianLayout>`,
			expected: EALayout{
				VertCompress: internal.ToPtr(stypes.OnOffOn),
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
