package docxpara

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
				XAlign:     internal.ToPtr(stypes.AlignLeft),
				YAlign:     internal.ToPtr(stypes.AlignCenter),
				HRule:      internal.ToPtr(stypes.HeightRuleExact),
				AnchorLock: internal.ToPtr(stypes.BinFlagTrue),
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
				XAlign:     internal.ToPtr(stypes.AlignLeft),
				YAlign:     internal.ToPtr(stypes.AlignCenter),
				HRule:      internal.ToPtr(stypes.HeightRuleExact),
				AnchorLock: internal.ToPtr(stypes.BinFlagTrue),
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
			if !compareFrameProps(result, tt.expected) {
				t.Errorf("Expected:\n%+v\nGot:\n%+v", tt.expected, result)
			}
		})
	}
}

// Helper function to compare FrameProp structs
// Helper function to compare FrameProp structs
func compareFrameProps(a, b FrameProp) bool {
	// Check each pointer for nil before comparing their dereferenced values
	return (a.Width == nil && b.Width == nil || (a.Width != nil && b.Width != nil && *a.Width == *b.Width)) &&
		(a.Height == nil && b.Height == nil || (a.Height != nil && b.Height != nil && *a.Height == *b.Height)) &&
		(a.DropCap == nil && b.DropCap == nil || (a.DropCap != nil && b.DropCap != nil && *a.DropCap == *b.DropCap)) &&
		(a.Lines == nil && b.Lines == nil || (a.Lines != nil && b.Lines != nil && *a.Lines == *b.Lines)) &&
		(a.VSpace == nil && b.VSpace == nil || (a.VSpace != nil && b.VSpace != nil && *a.VSpace == *b.VSpace)) &&
		(a.HSpace == nil && b.HSpace == nil || (a.HSpace != nil && b.HSpace != nil && *a.HSpace == *b.HSpace)) &&
		(a.Wrap == nil && b.Wrap == nil || (a.Wrap != nil && b.Wrap != nil && *a.Wrap == *b.Wrap)) &&
		(a.HAnchor == nil && b.HAnchor == nil || (a.HAnchor != nil && b.HAnchor != nil && *a.HAnchor == *b.HAnchor)) &&
		(a.VAnchor == nil && b.VAnchor == nil || (a.VAnchor != nil && b.VAnchor != nil && *a.VAnchor == *b.VAnchor)) &&
		(a.AbsHPos == nil && b.AbsHPos == nil || (a.AbsHPos != nil && b.AbsHPos != nil && *a.AbsHPos == *b.AbsHPos)) &&
		(a.AbsVPos == nil && b.AbsVPos == nil || (a.AbsVPos != nil && b.AbsVPos != nil && *a.AbsVPos == *b.AbsVPos)) &&
		(a.XAlign == nil && b.XAlign == nil || (a.XAlign != nil && b.XAlign != nil && *a.XAlign == *b.XAlign)) &&
		(a.YAlign == nil && b.YAlign == nil || (a.YAlign != nil && b.YAlign != nil && *a.YAlign == *b.YAlign)) &&
		(a.HRule == nil && b.HRule == nil || (a.HRule != nil && b.HRule != nil && *a.HRule == *b.HRule)) &&
		(a.AnchorLock == nil && b.AnchorLock == nil || (a.AnchorLock != nil && b.AnchorLock != nil && *a.AnchorLock == *b.AnchorLock))
}
