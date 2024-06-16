package sections

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/simpletypes"
)

func TestTextDirection_MarshalXML(t *testing.T) {
	tests := []struct {
		name          string
		textDir       TextDirection
		expected      string
		shouldMarshal bool // Indicates if the test should perform marshaling
	}{
		{
			name:          "WithVal",
			textDir:       TextDirection{Val: simpletypes.TextDirectionLrTb},
			expected:      `<w:textDirection w:val="lrTb"></w:textDirection>`,
			shouldMarshal: true,
		},
		{
			name:          "WithoutVal",
			textDir:       TextDirection{},
			expected:      `<w:textDirection></w:textDirection>`,
			shouldMarshal: true,
		},
		{
			name:          "Empty",
			textDir:       TextDirection{Val: simpletypes.TextDirection("")},
			expected:      ``,
			shouldMarshal: false, // Not expected to marshal because Val is empty
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldMarshal {
				xmlBytes, err := xml.Marshal(&tt.textDir)
				if err != nil {
					t.Fatalf("Error marshaling TextDirection: %v", err)
				}

				result := string(xmlBytes)
				if !strings.Contains(result, tt.expected) {
					t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expected, result)
				}
			} else {
				t.Skip("Skipping marshaling test for empty TextDirection")
			}
		})
	}
}

func TestTextDirection_UnmarshalXML(t *testing.T) {
	tests := []struct {
		xmlStr   string
		expected TextDirection
	}{
		{
			xmlStr:   `<w:textDirection w:val="lrTb"></w:textDirection>`,
			expected: TextDirection{Val: simpletypes.TextDirectionLrTb},
		},
		{
			xmlStr:   `<w:textDirection></w:textDirection>`,
			expected: TextDirection{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.xmlStr, func(t *testing.T) {
			var textDir TextDirection

			err := xml.Unmarshal([]byte(tt.xmlStr), &textDir)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if textDir.Val != tt.expected.Val {
				t.Errorf("Expected Val %s but got %s", tt.expected.Val, textDir.Val)
			}
		})
	}
}
