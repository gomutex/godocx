package dmlst

import (
	"encoding/xml"
	"testing"
)

func TestWrapTextFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected WrapText
	}{
		{"bothSides", WrapTextBothSides},
		{"left", WrapTextLeft},
		{"right", WrapTextRight},
		{"largest", WrapTextLargest},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := WrapTextFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestWrapTextFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := WrapTextFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid WrapText value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestWrapText_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected WrapText
	}{
		{`<element val="bothSides"></element>`, WrapTextBothSides},
		{`<element val="left"></element>`, WrapTextLeft},
		{`<element val="right"></element>`, WrapTextRight},
		{`<element val="largest"></element>`, WrapTextLargest},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name `xml:"element"`
				Val     WrapText `xml:"val,attr"`
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

func TestWrapText_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element val="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name `xml:"element"`
		Val     WrapText `xml:"val,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid WrapText value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
