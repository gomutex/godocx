package stypes

import (
	"encoding/xml"
	"testing"
)

// TestShadingFromStr tests the ShadingFromStr function with a few example values.
func TestShadingFromStr(t *testing.T) {
	tests := []struct {
		input    string
		expected Shading
		hasError bool
	}{
		{"nil", ShdNil, false},
		{"solid", ShdSolid, false},
		{"horzStripe", ShdHorzStripe, false},
		{"pct50", ShdPct50, false},
		{"invalid", "", true},
	}

	for _, test := range tests {
		result, err := ShadingFromStr(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("ShadingFromStr(%q) error = %v, hasError %v", test.input, err, test.hasError)
		}
		if result != test.expected {
			t.Errorf("ShadingFromStr(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}

// TestUnmarshalXMLAttr tests the UnmarshalXMLAttr method for the Shading type.
func TestUnmarshalXMLAttr(t *testing.T) {
	tests := []struct {
		input    xml.Attr
		expected Shading
		hasError bool
	}{
		{xml.Attr{Name: xml.Name{Local: "pattern"}, Value: "diagStripe"}, ShdDiagStripe, false},
		{xml.Attr{Name: xml.Name{Local: "pattern"}, Value: "thinHorzCross"}, ShdThinHorzCross, false},
		{xml.Attr{Name: xml.Name{Local: "pattern"}, Value: "invalid"}, "", true},
	}

	for _, test := range tests {
		var result Shading
		err := result.UnmarshalXMLAttr(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("UnmarshalXMLAttr(%v) error = %v, hasError %v", test.input, err, test.hasError)
		}
		if result != test.expected {
			t.Errorf("UnmarshalXMLAttr(%v) = %v, want %v", test.input, result, test.expected)
		}
	}
}
