package stypes

import (
	"encoding/xml"
	"testing"
)

// TestBorderStyleFromStr tests the BorderStyleFromStr function with a few example values.
func TestBorderStyleFromStr(t *testing.T) {
	tests := []struct {
		input    string
		expected BorderStyle
		hasError bool
	}{
		{"nil", BorderStyleNil, false},
		{"single", BorderStyleSingle, false},
		{"double", BorderStyleDouble, false},
		{"invalid", "", true},
	}

	for _, test := range tests {
		result, err := BorderStyleFromStr(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("BorderStyleFromStr(%q) error = %v, hasError %v", test.input, err, test.hasError)
		}
		if result != test.expected {
			t.Errorf("BorderStyleFromStr(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}

// TestUnmarshalXMLAttr tests the UnmarshalXMLAttr method for the BorderStyle type.
func TestBorderStyleUnmarshal(t *testing.T) {
	tests := []struct {
		input    xml.Attr
		expected BorderStyle
		hasError bool
	}{
		{xml.Attr{Name: xml.Name{Local: "border"}, Value: "dotted"}, BorderStyleDotted, false},
		{xml.Attr{Name: xml.Name{Local: "border"}, Value: "dashed"}, BorderStyleDashed, false},
		{xml.Attr{Name: xml.Name{Local: "border"}, Value: "invalid"}, "", true},
	}

	for _, test := range tests {
		var result BorderStyle
		err := result.UnmarshalXMLAttr(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("UnmarshalXMLAttr(%v) error = %v, hasError %v", test.input, err, test.hasError)
		}
		if result != test.expected {
			t.Errorf("UnmarshalXMLAttr(%v) = %v, want %v", test.input, result, test.expected)
		}
	}
}
