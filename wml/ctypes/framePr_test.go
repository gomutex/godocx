package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/internal"
	"github.com/gomutex/godocx/wml/stypes"
)

func TestFrameProp_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    FrameProp
		expected string
	}{
		{
			name: "With all attributes",
			input: FrameProp{
				Width:      internal.ToPtr(int64(500)),
				Height:     internal.ToPtr(int64((300))),
				DropCap:    internal.ToPtr(stypes.DropCapMargin),
				Lines:      internal.ToPtr(3),
				VSpace:     internal.ToPtr(int64((50))),
				HSpace:     internal.ToPtr(int64((20))),
				Wrap:       internal.ToPtr(stypes.WrapAround),
				HAnchor:    internal.ToPtr(stypes.AnchorMargin),
				VAnchor:    internal.ToPtr(stypes.AnchorPage),
				AbsHPos:    internal.ToPtr(100),
				AbsVPos:    internal.ToPtr(200),
				XAlign:     internal.ToPtr(stypes.XAlignLeft),
				YAlign:     internal.ToPtr(stypes.YAlignCenter),
				HRule:      internal.ToPtr(stypes.HeightRuleExact),
				AnchorLock: internal.ToPtr(stypes.OnOffTrue),
			},
			expected: `<w:framePr w:w="500" w:h="300" w:dropCap="margin" w:lines="3" w:hSpace="20" w:vSpace="50" w:wrap="around" w:hAnchor="margin" w:vAnchor="page" w:x="100" w:y="200" w:xAlign="left" w:yAlign="center" w:hRule="exact" w:anchorLock="true"></w:framePr>`,
		},
		{
			name: "Without optional attributes",
			input: FrameProp{
				Width:  internal.ToPtr(int64(500)),
				Height: internal.ToPtr(int64(300)),
			},
			expected: `<w:framePr w:w="500" w:h="300"></w:framePr>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:framePr"}}

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

func TestFrameProp_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected FrameProp
	}{
		{
			name: "With all attributes",
			inputXML: `<w:framePr w:w="500" w:h="300" w:dropCap="margin" w:lines="3" w:vSpace="50" w:hSpace="20" ` +
				`w:wrap="around" w:hAnchor="margin" w:vAnchor="page" w:x="100" w:y="200" w:xAlign="left" w:yAlign="center" ` +
				`w:hRule="exact" w:anchorLock="true"></w:framePr>`,
			expected: FrameProp{
				Width:      internal.ToPtr(int64(500)),
				Height:     internal.ToPtr(int64(300)),
				DropCap:    internal.ToPtr(stypes.DropCapMargin),
				Lines:      internal.ToPtr(3),
				VSpace:     internal.ToPtr(int64(50)),
				HSpace:     internal.ToPtr(int64(20)),
				Wrap:       internal.ToPtr(stypes.WrapAround),
				HAnchor:    internal.ToPtr(stypes.AnchorMargin),
				VAnchor:    internal.ToPtr(stypes.AnchorPage),
				AbsHPos:    internal.ToPtr(100),
				AbsVPos:    internal.ToPtr(200),
				XAlign:     internal.ToPtr(stypes.XAlignLeft),
				YAlign:     internal.ToPtr(stypes.YAlignCenter),
				HRule:      internal.ToPtr(stypes.HeightRuleExact),
				AnchorLock: internal.ToPtr(stypes.OnOffTrue),
			},
		},
		{
			name:     "Without optional attributes",
			inputXML: `<w:framePr w:w="500" w:h="300"></w:framePr>`,
			expected: FrameProp{
				Width:  internal.ToPtr(int64(500)),
				Height: internal.ToPtr(int64(300)),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result FrameProp

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			// Compare each field individually due to pointer comparisons
			if err := compareFrameProps(result, tt.expected); err != nil {
				t.Error(err)
			}
		})
	}
}

// Helper function to compare FrameProp structs
func compareFrameProps(a, b FrameProp) error {
	if err := internal.ComparePtr("Width", a.Width, b.Width); err != nil {
		return err
	}
	if err := internal.ComparePtr("Height", a.Height, b.Height); err != nil {
		return err
	}
	if err := internal.ComparePtr("DropCap", a.DropCap, b.DropCap); err != nil {
		return err
	}
	if err := internal.ComparePtr("Lines", a.Lines, b.Lines); err != nil {
		return err
	}
	if err := internal.ComparePtr("VSpace", a.VSpace, b.VSpace); err != nil {
		return err
	}
	if err := internal.ComparePtr("HSpace", a.HSpace, b.HSpace); err != nil {
		return err
	}
	if err := internal.ComparePtr("Wrap", a.Wrap, b.Wrap); err != nil {
		return err
	}
	if err := internal.ComparePtr("HAnchor", a.HAnchor, b.HAnchor); err != nil {
		return err
	}
	if err := internal.ComparePtr("VAnchor", a.VAnchor, b.VAnchor); err != nil {
		return err
	}
	if err := internal.ComparePtr("AbsHPos", a.AbsHPos, b.AbsHPos); err != nil {
		return err
	}
	if err := internal.ComparePtr("AbsVPos", a.AbsVPos, b.AbsVPos); err != nil {
		return err
	}
	if err := internal.ComparePtr("XAlign", a.XAlign, b.XAlign); err != nil {
		return err
	}
	if err := internal.ComparePtr("YAlign", a.YAlign, b.YAlign); err != nil {
		return err
	}
	if err := internal.ComparePtr("HRule", a.HRule, b.HRule); err != nil {
		return err
	}
	if err := internal.ComparePtr("AnchorLock", a.AnchorLock, b.AnchorLock); err != nil {
		return err
	}
	return nil
}
