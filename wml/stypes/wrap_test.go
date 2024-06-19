package stypes

import (
	"encoding/xml"
	"testing"
)

func TestWrapFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected Wrap
	}{
		{"auto", WrapAuto},
		{"notBeside", WrapNotBeside},
		{"around", WrapAround},
		{"tight", WrapTight},
		{"through", WrapThrough},
		{"none", WrapNone},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := WrapFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestWrapFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := WrapFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid Wrap value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestWrap_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected Wrap
	}{
		{`<element val="auto"></element>`, WrapAuto},
		{`<element val="notBeside"></element>`, WrapNotBeside},
		{`<element val="around"></element>`, WrapAround},
		{`<element val="tight"></element>`, WrapTight},
		{`<element val="through"></element>`, WrapThrough},
		{`<element val="none"></element>`, WrapNone},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name `xml:"element"`
				Val     Wrap     `xml:"val,attr"`
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

func TestWrap_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element val="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name `xml:"element"`
		Val     Wrap     `xml:"val,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid Wrap value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
