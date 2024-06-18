package stypes

import (
	"encoding/xml"
	"testing"
)

func TestPageOrientFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected PageOrient
	}{
		{"portrait", PageOrientPortrait},
		{"landscape", PageOrientLandscape},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := PageOrientFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestPageOrientFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := PageOrientFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid Orient Input"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestPageOrient_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected PageOrient
	}{
		{`<element orient="portrait"></element>`, PageOrientPortrait},
		{`<element orient="landscape"></element>`, PageOrientLandscape},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name   `xml:"element"`
				Orient  PageOrient `xml:"orient,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.Orient != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.Orient)
			}
		})
	}
}

func TestPageOrient_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element orient="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name   `xml:"element"`
		Orient  PageOrient `xml:"orient,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid Orient Input"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
