package ctypes

import (
	"encoding/xml"
	"fmt"
	"strings"
	"testing"

	"github.com/gomutex/godocx/internal"
	"github.com/gomutex/godocx/wml/stypes"
)

func TestLatentStyle_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		latent   LatentStyle
		expected string
	}{
		{
			name: "All attributes set",
			latent: LatentStyle{
				DefLockedState:    internal.ToPtr(stypes.OnOffOn),
				DefUIPriority:     internal.ToPtr(99),
				DefSemiHidden:     internal.ToPtr(stypes.OnOffOn),
				DefUnhideWhenUsed: internal.ToPtr(stypes.OnOffOn),
				DefQFormat:        internal.ToPtr(stypes.OnOffOn),
				Count:             internal.ToPtr(3),
				LsdExceptions: []LsdException{
					{
						Name:           "Heading1",
						Locked:         internal.ToPtr(stypes.OnOffOn),
						UIPriority:     internal.ToPtr(99),
						SemiHidden:     internal.ToPtr(stypes.OnOffOn),
						UnhideWhenUsed: internal.ToPtr(stypes.OnOffOn),
						QFormat:        internal.ToPtr(stypes.OnOffOn),
					},
					{
						Name: "Heading2",
					},
				},
			},
			expected: `<w:latentStyles w:defLockedState="on" w:defUIPriority="99" w:defSemiHidden="on" w:defUnhideWhenUsed="on" w:defQFormat="on" w:count="3"><w:lsdException w:name="Heading1" w:locked="on" w:uiPriority="99" w:semiHidden="on" w:unhideWhenUsed="on" w:qFormat="on"></w:lsdException><w:lsdException w:name="Heading2"></w:lsdException></w:latentStyles>`,
		},
		{
			name: "Only required attributes set",
			latent: LatentStyle{
				LsdExceptions: []LsdException{
					{
						Name: "Heading1",
					},
				},
			},
			expected: `<w:latentStyles><w:lsdException w:name="Heading1"></w:lsdException></w:latentStyles>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			e := xml.NewEncoder(&result)
			err := tt.latent.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:latentStyles"}})
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
func TestLatentStyle_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected LatentStyle
	}{
		{
			name: "All attributes set",
			inputXML: `<w:latentStyles w:defLockedState="on" w:defUIPriority="99" w:defSemiHidden="on" w:defUnhideWhenUsed="on" w:defQFormat="on" w:count="3">
				<w:lsdException w:name="Heading1" w:locked="on" w:uiPriority="99" w:semiHidden="on" w:unhideWhenUsed="on" w:qFormat="on"></w:lsdException>
				<w:lsdException w:name="Heading2"></w:lsdException>
			</w:latentStyles>`,
			expected: LatentStyle{
				DefLockedState:    internal.ToPtr(stypes.OnOffOn),
				DefUIPriority:     internal.ToPtr(99),
				DefSemiHidden:     internal.ToPtr(stypes.OnOffOn),
				DefUnhideWhenUsed: internal.ToPtr(stypes.OnOffOn),
				DefQFormat:        internal.ToPtr(stypes.OnOffOn),
				Count:             internal.ToPtr(3),
				LsdExceptions: []LsdException{
					{
						Name:           "Heading1",
						Locked:         internal.ToPtr(stypes.OnOffOn),
						UIPriority:     internal.ToPtr(99),
						SemiHidden:     internal.ToPtr(stypes.OnOffOn),
						UnhideWhenUsed: internal.ToPtr(stypes.OnOffOn),
						QFormat:        internal.ToPtr(stypes.OnOffOn),
					},
					{
						Name: "Heading2",
					},
				},
			},
		},
		{
			name: "Only required attributes set",
			inputXML: `<w:latentStyles>
				<w:lsdException w:name="Heading1"></w:lsdException>
			</w:latentStyles>`,
			expected: LatentStyle{
				LsdExceptions: []LsdException{
					{
						Name: "Heading1",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var latent LatentStyle
			err := xml.Unmarshal([]byte(tt.inputXML), &latent)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if err := compareLatentStyles(latent, tt.expected); err != nil {
				t.Error(err)
			}
		})
	}
}

func compareLatentStyles(a, b LatentStyle) error {
	if err := internal.ComparePtr("DefLockedState", a.DefLockedState, b.DefLockedState); err != nil {
		return err
	}
	if err := internal.ComparePtr("DefUIPriority", a.DefUIPriority, b.DefUIPriority); err != nil {
		return err
	}
	if err := internal.ComparePtr("DefSemiHidden", a.DefSemiHidden, b.DefSemiHidden); err != nil {
		return err
	}
	if err := internal.ComparePtr("DefUnhideWhenUsed", a.DefUnhideWhenUsed, b.DefUnhideWhenUsed); err != nil {
		return err
	}
	if err := internal.ComparePtr("DefQFormat", a.DefQFormat, b.DefQFormat); err != nil {
		return err
	}
	if err := internal.ComparePtr("Count", a.Count, b.Count); err != nil {
		return err
	}

	if len(a.LsdExceptions) != len(b.LsdExceptions) {
		return fmt.Errorf("LsdExceptions length mismatch, expected %d but got %d", len(b.LsdExceptions), len(a.LsdExceptions))
	}

	for i := range a.LsdExceptions {
		if err := compareLsdExFields(a.LsdExceptions[i], b.LsdExceptions[i]); err != nil {
			return fmt.Errorf("LsdExceptions[%d]: %w", i, err)
		}
	}

	return nil
}
