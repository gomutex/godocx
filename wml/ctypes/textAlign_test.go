package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestTextAlign_MarshalXML(t *testing.T) {
	tests := []struct {
		name          string
		textAlignment TextAlign
		expected      string
		shouldMarshal bool // Indicates if the test should perform marshaling
	}{
		{
			name:          "WithVal",
			textAlignment: TextAlign{Val: stypes.TextAlignTop},
			expected:      `<w:textAlignment w:val="top"></w:textAlignment>`,
			shouldMarshal: true,
		},
		{
			name:          "WithoutVal",
			textAlignment: TextAlign{},
			expected:      ``,
			shouldMarshal: false,
		},
		{
			name:          "Empty",
			textAlignment: TextAlign{Val: stypes.TextAlign("")},
			expected:      ``,
			shouldMarshal: false, // Not expected to marshal because Val is empty
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldMarshal {
				xmlBytes, err := xml.Marshal(&tt.textAlignment)
				if err != nil {
					t.Fatalf("Error marshaling TextAlign: %v", err)
				}

				result := string(xmlBytes)
				if !strings.Contains(result, tt.expected) {
					t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expected, result)
				}
			} else {
				t.Skip("Skipping marshaling test for empty TextAlign")
			}
		})
	}
}

func TestTextAlign_UnmarshalXML(t *testing.T) {
	tests := []struct {
		xmlStr    string
		expected  TextAlign
		shouldErr bool
	}{
		{
			xmlStr:    `<w:textAlignment w:val="top"></w:textAlignment>`,
			expected:  TextAlign{Val: stypes.TextAlignTop},
			shouldErr: false,
		},
		// {
		// 	xmlStr:    `<w:textAlignment></w:textAlignment>`,
		// 	expected:  TextAlign{},
		// 	shouldErr: true,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.xmlStr, func(t *testing.T) {
			var textAlignment TextAlign

			err := xml.Unmarshal([]byte(tt.xmlStr), &textAlignment)
			if tt.shouldErr && err == nil {
				t.Fatal("err should not be nil")
			} else if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			} else {
				if textAlignment.Val != tt.expected.Val {
					t.Errorf("Expected Val %s but got %s", tt.expected.Val, textAlignment.Val)
				}
			}
		})
	}
}
