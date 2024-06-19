package stypes

import (
	"encoding/xml"
	"testing"
)

func TestEmFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected Em
	}{
		{"none", EmNone},
		{"dot", EmDot},
		{"comma", EmComma},
		{"circle", EmCircle},
		{"underDot", EmUnderDot},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := EmFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestEmFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := EmFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid Em value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestEm_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected Em
	}{
		{`<element val="none"></element>`, EmNone},
		{`<element val="dot"></element>`, EmDot},
		{`<element val="comma"></element>`, EmComma},
		{`<element val="circle"></element>`, EmCircle},
		{`<element val="underDot"></element>`, EmUnderDot},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name `xml:"element"`
				Val     Em       `xml:"val,attr"`
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

func TestEm_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element val="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name `xml:"element"`
		Val     Em       `xml:"val,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid Em value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
