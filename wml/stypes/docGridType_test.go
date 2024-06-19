package stypes

import (
	"encoding/xml"
	"testing"
)

func TestDocGridTypeFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected DocGridType
	}{
		{"default", DocGridDefault},
		{"lines", DocGridLines},
		{"linesAndChars", DocGridLinesAndChars},
		{"snapToChars", DocGridSnapToChars},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := DocGridTypeFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestDocGridTypeFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := DocGridTypeFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid Docgrid Type"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestDocGridType_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected DocGridType
	}{
		{`<element gridType="default"></element>`, DocGridDefault},
		{`<element gridType="lines"></element>`, DocGridLines},
		{`<element gridType="linesAndChars"></element>`, DocGridLinesAndChars},
		{`<element gridType="snapToChars"></element>`, DocGridSnapToChars},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName  xml.Name    `xml:"element"`
				GridType DocGridType `xml:"gridType,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.GridType != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.GridType)
			}
		})
	}
}

func TestDocGridType_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element gridType="invalidValue"></element>`

	type Element struct {
		XMLName  xml.Name    `xml:"element"`
		GridType DocGridType `xml:"gridType,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid Docgrid Type"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
