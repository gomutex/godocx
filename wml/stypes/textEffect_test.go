package stypes

import (
	"encoding/xml"
	"testing"
)

func TestTextEffectFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected TextEffect
	}{
		{"blinkBackground", TextEffectBlinkBackground},
		{"lights", TextEffectLights},
		{"antsBlack", TextEffectAntsBlack},
		{"antsRed", TextEffectAntsRed},
		{"shimmer", TextEffectShimmer},
		{"sparkle", TextEffectSparkle},
		{"none", TextEffectNone},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := TextEffectFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestTextEffectFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	_, err := TextEffectFromStr(input)
	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none", input)
	}

	expectedError := "invalid TextEffect value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestTextEffect_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected TextEffect
	}{
		{`<element effect="blinkBackground"></element>`, TextEffectBlinkBackground},
		{`<element effect="lights"></element>`, TextEffectLights},
		{`<element effect="antsBlack"></element>`, TextEffectAntsBlack},
		{`<element effect="antsRed"></element>`, TextEffectAntsRed},
		{`<element effect="shimmer"></element>`, TextEffectShimmer},
		{`<element effect="sparkle"></element>`, TextEffectSparkle},
		{`<element effect="none"></element>`, TextEffectNone},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name   `xml:"element"`
				Effect  TextEffect `xml:"effect,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.Effect != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.Effect)
			}
		})
	}
}

func TestTextEffect_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element effect="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name   `xml:"element"`
		Effect  TextEffect `xml:"effect,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)
	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "invalid TextEffect value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
