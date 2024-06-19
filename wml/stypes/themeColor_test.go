package stypes

import (
	"encoding/xml"
	"testing"
)

// TestThemeColorFromStr tests the ThemeColorFromStr function with a few example values.
func TestThemeColorFromStr(t *testing.T) {
	tests := []struct {
		input    string
		expected ThemeColor
		hasError bool
	}{
		{"dark1", ThemeColorDark1, false},
		{"light2", ThemeColorLight2, false},
		{"accent5", ThemeColorAccent5, false},
		{"invalid", "", true},
	}

	for _, test := range tests {
		result, err := ThemeColorFromStr(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("ThemeColorFromStr(%q) error = %v, hasError %v", test.input, err, test.hasError)
		}
		if result != test.expected {
			t.Errorf("ThemeColorFromStr(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}

// TestUnmarshalXMLAttr tests the UnmarshalXMLAttr method for the ThemeColor type.
func TestThemeColorUnmarshal(t *testing.T) {
	tests := []struct {
		input    xml.Attr
		expected ThemeColor
		hasError bool
	}{
		{xml.Attr{Name: xml.Name{Local: "color"}, Value: "accent3"}, ThemeColorAccent3, false},
		{xml.Attr{Name: xml.Name{Local: "color"}, Value: "background2"}, ThemeColorBackground2, false},
		{xml.Attr{Name: xml.Name{Local: "color"}, Value: "invalid"}, "", true},
	}

	for _, test := range tests {
		var result ThemeColor
		err := result.UnmarshalXMLAttr(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("UnmarshalXMLAttr(%v) error = %v, hasError %v", test.input, err, test.hasError)
		}
		if result != test.expected {
			t.Errorf("UnmarshalXMLAttr(%v) = %v, want %v", test.input, result, test.expected)
		}
	}
}
