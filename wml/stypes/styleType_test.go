package stypes

import (
	"encoding/xml"
	"testing"
)

func TestStyleTypeFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected StyleType
	}{
		{"paragraph", StyleTypeParagraph},
		{"character", StyleTypeCharacter},
		{"table", StyleTypeTable},
		{"numbering", StyleTypeNumbering},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := StyleTypeFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestStyleTypeFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := StyleTypeFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "invalid StyleType value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestStyleType_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected StyleType
	}{
		{`<element type="paragraph"></element>`, StyleTypeParagraph},
		{`<element type="character"></element>`, StyleTypeCharacter},
		{`<element type="table"></element>`, StyleTypeTable},
		{`<element type="numbering"></element>`, StyleTypeNumbering},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name  `xml:"element"`
				Type    StyleType `xml:"type,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.Type != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.Type)
			}
		})
	}
}

func TestStyleType_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element type="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name  `xml:"element"`
		Type    StyleType `xml:"type,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "invalid StyleType value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
