package ctypes

import (
	"encoding/xml"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestLongHexNum_MarshalXML(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"1A2B", `<LongHexNum w:val="1A2B"></LongHexNum>`, false},
		{"1a2b", `<LongHexNum w:val="1A2B"></LongHexNum>`, false},
		{"", `<LongHexNum w:val=""></LongHexNum>`, false},
		{"123", "", true},
		{"12345", "", true},
		{"1G2H", "", true},
	}

	for _, tt := range tests {
		longHexNum, err := stypes.LongHexNumFromStr(tt.input)
		if (err != nil) != tt.hasError {
			t.Errorf("stypes.LongHexNumFromStr(%q) error = %v, expected error = %v", tt.input, err, tt.hasError)
			continue
		}

		if err != nil {
			continue
		}

		ln := &LongHexNum{Val: longHexNum}
		output, err := xml.Marshal(ln)
		if (err != nil) != tt.hasError {
			t.Errorf("Marshal(%q) error = %v, expected error = %v", tt.input, err, tt.hasError)
		}

		if string(output) != tt.expected {
			t.Errorf("Marshal(%q) = %q, expected %q", tt.input, string(output), tt.expected)
		}
	}
}

func TestLongHexNum_UnmarshalXML(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{`<LongHexNum w:val="1A2B"></LongHexNum>`, "1A2B", false},
		{`<LongHexNum w:val="1a2b"></LongHexNum>`, "1A2B", false},
		{`<LongHexNum w:val=""></LongHexNum>`, "", false},
		{`<LongHexNum w:val="123"></LongHexNum>`, "", true},
		{`<LongHexNum w:val="12345"></LongHexNum>`, "", true},
		{`<LongHexNum w:val="1G2H"></LongHexNum>`, "", true},
	}

	for _, tt := range tests {
		var ln LongHexNum
		err := xml.Unmarshal([]byte(tt.input), &ln)
		if (err != nil) != tt.hasError {
			t.Errorf("Unmarshal(%q) error = %v, expected error = %v", tt.input, err, tt.hasError)
		}

		if string(ln.Val) != tt.expected {
			t.Errorf("Unmarshal(%q) = %q, expected %q", tt.input, string(ln.Val), tt.expected)
		}
	}
}
