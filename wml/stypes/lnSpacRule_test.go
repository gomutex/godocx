package stypes

import (
	"encoding/xml"
	"testing"
)

func TestLineSpacingRuleFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected LineSpacingRule
	}{
		{"auto", LineSpacingRuleAuto},
		{"exact", LineSpacingRuleExact},
		{"atLeast", LineSpacingRuleAtLeast},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := LineSpacingRuleFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestLineSpacingRuleFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := LineSpacingRuleFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "invalid LineSpacingRule value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestLineSpacingRule_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected LineSpacingRule
	}{
		{`<element lineSpacing="auto"></element>`, LineSpacingRuleAuto},
		{`<element lineSpacing="exact"></element>`, LineSpacingRuleExact},
		{`<element lineSpacing="atLeast"></element>`, LineSpacingRuleAtLeast},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName     xml.Name        `xml:"element"`
				LineSpacing LineSpacingRule `xml:"lineSpacing,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.LineSpacing != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.LineSpacing)
			}
		})
	}
}

func TestLineSpacingRule_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element lineSpacing="invalidValue"></element>`

	type Element struct {
		XMLName     xml.Name        `xml:"element"`
		LineSpacing LineSpacingRule `xml:"lineSpacing,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "invalid LineSpacingRule value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
