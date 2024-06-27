package ctypes

import (
	"encoding/xml"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestFormProt_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		formProt FormProt
		expected string
	}{
		{
			name:     "WithVal",
			formProt: FormProt{Val: stypes.OnOff("on")},
			expected: `<w:formProt w:val="on"></w:formProt>`,
		},
		{
			name:     "WithoutVal",
			formProt: FormProt{},
			expected: `<w:formProt></w:formProt>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xmlBytes, err := xml.Marshal(&tt.formProt)
			if err != nil {
				t.Fatalf("Error marshaling FormProt: %v", err)
			}

			result := string(xmlBytes)
			if result != tt.expected {
				t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expected, result)
			}
		})
	}
}

func TestFormProt_UnmarshalXML(t *testing.T) {
	tests := []struct {
		xmlStr   string
		expected FormProt
	}{
		{
			xmlStr:   `<w:formProt w:val="on"></w:formProt>`,
			expected: FormProt{Val: stypes.OnOff("on")},
		},
		{
			xmlStr:   `<w:formProt></w:formProt>`,
			expected: FormProt{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.xmlStr, func(t *testing.T) {
			var formProt FormProt

			err := xml.Unmarshal([]byte(tt.xmlStr), &formProt)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if formProt.Val != tt.expected.Val {
				t.Errorf("Expected Val %s but got %s", tt.expected.Val, formProt.Val)
			}
		})
	}
}

func TestFormProt_UnmarshalXML_EmptyElement(t *testing.T) {
	xmlStr := `<w:formProt/>`

	var formProt FormProt

	err := xml.Unmarshal([]byte(xmlStr), &formProt)
	if err != nil {
		t.Fatalf("Error unmarshaling XML: %v", err)
	}

	if formProt.Val != "" {
		t.Errorf("Expected empty Val but got %s", formProt.Val)
	}
}
