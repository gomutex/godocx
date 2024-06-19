package stypes

import (
	"encoding/xml"
	"testing"
)

func TestDropCapFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected DropCap
	}{
		{"none", DropCapNone},
		{"drop", DropCapInside},
		{"margin", DropCapMargin},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := DropCapFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestDropCapFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := DropCapFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid DropCap value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestDropCap_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected DropCap
	}{
		{`<element val="none"></element>`, DropCapNone},
		{`<element val="drop"></element>`, DropCapInside},
		{`<element val="margin"></element>`, DropCapMargin},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name `xml:"element"`
				Val     DropCap  `xml:"val,attr"`
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

func TestDropCap_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element val="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name `xml:"element"`
		Val     DropCap  `xml:"val,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid DropCap value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
