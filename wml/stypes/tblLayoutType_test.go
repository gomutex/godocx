package stypes

import (
	"encoding/xml"
	"testing"
)

func TestTableLayoutFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected TableLayout
	}{
		{"fixed", TableLayoutFixed},
		{"autofit", TableLayoutAutoFit},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := TableLayoutFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestTableLayoutFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := TableLayoutFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid Table Layout Type"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestTableLayout_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected TableLayout
	}{
		{`<element layout="fixed"></element>`, TableLayoutFixed},
		{`<element layout="autofit"></element>`, TableLayoutAutoFit},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name    `xml:"element"`
				Layout  TableLayout `xml:"layout,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.Layout != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.Layout)
			}
		})
	}
}

func TestTableLayout_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element layout="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name    `xml:"element"`
		Layout  TableLayout `xml:"layout,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid Table Layout Type"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
