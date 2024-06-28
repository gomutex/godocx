package stypes

import (
	"encoding/xml"
	"testing"
)

func TestBreakTypeFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected BreakType
	}{
		{"page", BreakTypePage},
		{"column", BreakTypeColumn},
		{"textWrapping", BreakTypeTextWrapping},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := BreakTypeFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestBreakTypeFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := BreakTypeFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "invalid BreakType value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestBreakType_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected BreakType
	}{
		{`<element breakType="page"></element>`, BreakTypePage},
		{`<element breakType="column"></element>`, BreakTypeColumn},
		{`<element breakType="textWrapping"></element>`, BreakTypeTextWrapping},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName   xml.Name  `xml:"element"`
				BreakType BreakType `xml:"breakType,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.BreakType != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.BreakType)
			}
		})
	}
}

func TestBreakType_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element breakType="invalidValue"></element>`

	type Element struct {
		XMLName   xml.Name  `xml:"element"`
		BreakType BreakType `xml:"breakType,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "invalid BreakType value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestBreakClearFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected BreakClear
	}{
		{"none", BreakClearNone},
		{"left", BreakClearLeft},
		{"right", BreakClearRight},
		{"all", BreakClearAll},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := BreakClearFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestBreakClearFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := BreakClearFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "invalid BreakClear value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestBreakClear_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected BreakClear
	}{
		{`<element breakClear="none"></element>`, BreakClearNone},
		{`<element breakClear="left"></element>`, BreakClearLeft},
		{`<element breakClear="right"></element>`, BreakClearRight},
		{`<element breakClear="all"></element>`, BreakClearAll},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName    xml.Name   `xml:"element"`
				BreakClear BreakClear `xml:"breakClear,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.BreakClear != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.BreakClear)
			}
		})
	}
}

func TestBreakClear_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element breakClear="invalidValue"></element>`

	type Element struct {
		XMLName    xml.Name   `xml:"element"`
		BreakClear BreakClear `xml:"breakClear,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "invalid BreakClear value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
