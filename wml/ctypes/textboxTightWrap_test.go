package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestTextboxTightWrap_MarshalXML(t *testing.T) {
	tests := []struct {
		name          string
		tightWrap     TextboxTightWrap
		expected      string
		shouldMarshal bool // Indicates if the test should perform marshaling
	}{
		{
			name:          "None",
			tightWrap:     TextboxTightWrap{Val: stypes.TextboxTightWrapNone},
			expected:      `<w:textboxTightWrap w:val="none"></w:textboxTightWrap>`,
			shouldMarshal: true,
		},
		{
			name:          "AllLines",
			tightWrap:     TextboxTightWrap{Val: stypes.TextboxTightWrapAllLines},
			expected:      `<w:textboxTightWrap w:val="allLines"></w:textboxTightWrap>`,
			shouldMarshal: true,
		},
		{
			name:          "FirstAndLastLine",
			tightWrap:     TextboxTightWrap{Val: stypes.TextboxTightWrapFirstAndLastLine},
			expected:      `<w:textboxTightWrap w:val="firstAndLastLine"></w:textboxTightWrap>`,
			shouldMarshal: true,
		},
		{
			name:          "FirstLineOnly",
			tightWrap:     TextboxTightWrap{Val: stypes.TextboxTightWrapFirstLineOnly},
			expected:      `<w:textboxTightWrap w:val="firstLineOnly"></w:textboxTightWrap>`,
			shouldMarshal: true,
		},
		{
			name:          "LastLineOnly",
			tightWrap:     TextboxTightWrap{Val: stypes.TextboxTightWrapLastLineOnly},
			expected:      `<w:textboxTightWrap w:val="lastLineOnly"></w:textboxTightWrap>`,
			shouldMarshal: true,
		},
		{
			name:          "Empty",
			tightWrap:     TextboxTightWrap{Val: stypes.TextboxTightWrap("")},
			expected:      ``,
			shouldMarshal: false, // Not expected to marshal because Val is empty
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldMarshal {
				xmlBytes, err := xml.Marshal(&tt.tightWrap)
				if err != nil {
					t.Fatalf("Error marshaling TextboxTightWrap: %v", err)
				}

				result := string(xmlBytes)
				if !strings.Contains(result, tt.expected) {
					t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expected, result)
				}
			} else {
				t.Skip("Skipping marshaling test for empty TextboxTightWrap")
			}
		})
	}
}

func TestTextboxTightWrap_UnmarshalXML(t *testing.T) {
	tests := []struct {
		xmlStr   string
		expected TextboxTightWrap
	}{
		{
			xmlStr:   `<w:textboxTightWrap w:val="none"></w:textboxTightWrap>`,
			expected: TextboxTightWrap{Val: stypes.TextboxTightWrapNone},
		},
		{
			xmlStr:   `<w:textboxTightWrap w:val="allLines"></w:textboxTightWrap>`,
			expected: TextboxTightWrap{Val: stypes.TextboxTightWrapAllLines},
		},
		{
			xmlStr:   `<w:textboxTightWrap w:val="firstAndLastLine"></w:textboxTightWrap>`,
			expected: TextboxTightWrap{Val: stypes.TextboxTightWrapFirstAndLastLine},
		},
		{
			xmlStr:   `<w:textboxTightWrap w:val="firstLineOnly"></w:textboxTightWrap>`,
			expected: TextboxTightWrap{Val: stypes.TextboxTightWrapFirstLineOnly},
		},
		{
			xmlStr:   `<w:textboxTightWrap w:val="lastLineOnly"></w:textboxTightWrap>`,
			expected: TextboxTightWrap{Val: stypes.TextboxTightWrapLastLineOnly},
		},
		{
			xmlStr:   `<w:textboxTightWrap></w:textboxTightWrap>`,
			expected: TextboxTightWrap{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.xmlStr, func(t *testing.T) {
			var tightWrap TextboxTightWrap

			err := xml.Unmarshal([]byte(tt.xmlStr), &tightWrap)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if tightWrap.Val != tt.expected.Val {
				t.Errorf("Expected Val %s but got %s", tt.expected.Val, tightWrap.Val)
			}
		})
	}
}
