package stypes

import (
	"encoding/xml"
	"testing"
)

func TestHeightRuleFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected HeightRule
	}{
		{"auto", HeightRuleAuto},
		{"exact", HeightRuleExact},
		{"atLeast", HeightRuleAtLeast},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := HeightRuleFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestHeightRuleFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := HeightRuleFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid HeightRule value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestHeightRule_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected HeightRule
	}{
		{`<element val="auto"></element>`, HeightRuleAuto},
		{`<element val="exact"></element>`, HeightRuleExact},
		{`<element val="atLeast"></element>`, HeightRuleAtLeast},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name   `xml:"element"`
				Val     HeightRule `xml:"val,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.Val != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.Val)
			}
		})
	}
}

func TestHeightRule_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element val="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name   `xml:"element"`
		Val     HeightRule `xml:"val,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid HeightRule value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
