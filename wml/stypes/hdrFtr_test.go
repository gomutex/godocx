package stypes

import (
	"encoding/xml"
	"testing"
)

func TestHdrFtrFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected HdrFtrType
	}{
		{"default", HdrFtrDefault},
		{"even", HdrFtrEven},
		{"first", HdrFtrFirst},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := HdrFtrFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestHdrFtrFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := HdrFtrFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid Header or Footer Type"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestHdrFtr_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected HdrFtrType
	}{
		{`<element hdrFtrType="default"></element>`, HdrFtrDefault},
		{`<element hdrFtrType="even"></element>`, HdrFtrEven},
		{`<element hdrFtrType="first"></element>`, HdrFtrFirst},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName    xml.Name   `xml:"element"`
				HdrFtrType HdrFtrType `xml:"hdrFtrType,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.HdrFtrType != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.HdrFtrType)
			}
		})
	}
}

func TestHdrFtr_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element hdrFtrType="invalidValue"></element>`

	type Element struct {
		XMLName    xml.Name   `xml:"element"`
		HdrFtrType HdrFtrType `xml:"hdrFtrType,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid Header or Footer Type"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
