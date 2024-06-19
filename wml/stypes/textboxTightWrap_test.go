package stypes

import (
	"encoding/xml"
	"testing"
)

func TestTextboxTightWrapFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected TextboxTightWrap
	}{
		{"none", TextboxTightWrapNone},
		{"allLines", TextboxTightWrapAllLines},
		{"firstAndLastLine", TextboxTightWrapFirstAndLastLine},
		{"firstLineOnly", TextboxTightWrapFirstLineOnly},
		{"lastLineOnly", TextboxTightWrapLastLineOnly},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := TextboxTightWrapFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestTextboxTightWrapFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := TextboxTightWrapFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid Textbox Tight Wrap value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestTextboxTightWrap_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected TextboxTightWrap
	}{
		{`<element val="none"></element>`, TextboxTightWrapNone},
		{`<element val="allLines"></element>`, TextboxTightWrapAllLines},
		{`<element val="firstAndLastLine"></element>`, TextboxTightWrapFirstAndLastLine},
		{`<element val="firstLineOnly"></element>`, TextboxTightWrapFirstLineOnly},
		{`<element val="lastLineOnly"></element>`, TextboxTightWrapLastLineOnly},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name         `xml:"element"`
				Val     TextboxTightWrap `xml:"val,attr"`
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

func TestTextboxTightWrap_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element val="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name         `xml:"element"`
		Val     TextboxTightWrap `xml:"val,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid Textbox Tight Wrap value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
