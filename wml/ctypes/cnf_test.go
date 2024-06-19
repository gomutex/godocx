package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestCnf_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		cnf      Cnf
		expected string
	}{
		{
			name:     "WithVal",
			cnf:      Cnf{Val: "12345"},
			expected: `<w:cnfStyle w:val="12345"></w:cnfStyle>`,
		},
		{
			name:     "EmptyVal",
			cnf:      Cnf{Val: ""},
			expected: `<w:cnfStyle w:val=""></w:cnfStyle>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xmlBytes, err := xml.Marshal(&tt.cnf)
			if err != nil {
				t.Fatalf("Error marshaling Cnf: %v", err)
			}

			result := string(xmlBytes)
			if !strings.Contains(result, tt.expected) {
				t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expected, result)
			}
		})
	}
}

func TestCnf_UnmarshalXML(t *testing.T) {
	tests := []struct {
		xmlStr   string
		expected Cnf
	}{
		{
			xmlStr:   `<w:cnfStyle w:val="12345"></w:cnfStyle>`,
			expected: Cnf{Val: "12345"},
		},
		{
			xmlStr:   `<w:cnfStyle w:val=""></w:cnfStyle>`,
			expected: Cnf{Val: ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.xmlStr, func(t *testing.T) {
			var cnf Cnf

			err := xml.Unmarshal([]byte(tt.xmlStr), &cnf)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if cnf.Val != tt.expected.Val {
				t.Errorf("Expected Val %s but got %s", tt.expected.Val, cnf.Val)
			}
		})
	}
}
