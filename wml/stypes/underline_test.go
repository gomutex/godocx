package stypes

import (
	"encoding/xml"
	"testing"
)

func TestUnderlineFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected Underline
	}{
		{"none", UnderlineNone},
		{"single", UnderlineSingle},
		{"words", UnderlineWords},
		{"double", UnderlineDouble},
		{"dotted", UnderlineDotted},
		{"thick", UnderlineThick},
		{"dash", UnderlineDash},
		{"dotDash", UnderlineDotDash},
		{"dotDotDash", UnderlineDotDotDash},
		{"wavy", UnderlineWavy},
		{"dottedHeavy", UnderlineDottedHeavy},
		{"dashHeavy", UnderlineDashHeavy},
		{"dotDashHeavy", UnderlineDotDashHeavy},
		{"dotDotDashHeavy", UnderlineDotDotDashHeavy},
		{"wavyHeavy", UnderlineWavyHeavy},
		{"dashLong", UnderlineDashLong},
		{"wavyDouble", UnderlineWavyDouble},
		{"dashLongHeavy", UnderlineDashLongHeavy},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := UnderlineFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestUnderlineFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := UnderlineFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "invalid Underline value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestUnderline_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected Underline
	}{
		{`<element underline="none"></element>`, UnderlineNone},
		{`<element underline="single"></element>`, UnderlineSingle},
		{`<element underline="words"></element>`, UnderlineWords},
		{`<element underline="double"></element>`, UnderlineDouble},
		{`<element underline="dotted"></element>`, UnderlineDotted},
		{`<element underline="thick"></element>`, UnderlineThick},
		{`<element underline="dash"></element>`, UnderlineDash},
		{`<element underline="dotDash"></element>`, UnderlineDotDash},
		{`<element underline="dotDotDash"></element>`, UnderlineDotDotDash},
		{`<element underline="wavy"></element>`, UnderlineWavy},
		{`<element underline="dottedHeavy"></element>`, UnderlineDottedHeavy},
		{`<element underline="dashHeavy"></element>`, UnderlineDashHeavy},
		{`<element underline="dotDashHeavy"></element>`, UnderlineDotDashHeavy},
		{`<element underline="dotDotDashHeavy"></element>`, UnderlineDotDotDashHeavy},
		{`<element underline="wavyHeavy"></element>`, UnderlineWavyHeavy},
		{`<element underline="dashLong"></element>`, UnderlineDashLong},
		{`<element underline="wavyDouble"></element>`, UnderlineWavyDouble},
		{`<element underline="dashLongHeavy"></element>`, UnderlineDashLongHeavy},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName   xml.Name  `xml:"element"`
				Underline Underline `xml:"underline,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.Underline != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.Underline)
			}
		})
	}
}

func TestUnderline_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element underline="invalidValue"></element>`

	type Element struct {
		XMLName   xml.Name  `xml:"element"`
		Underline Underline `xml:"underline,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "invalid Underline value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
