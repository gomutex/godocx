package stypes

import (
	"encoding/xml"
	"testing"
)

func TestBinFlagFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected BinFlag
	}{
		{"0", BinFlagZero},
		{"1", BinFlagOne},
		{"false", BinFlagFalse},
		{"true", BinFlagTrue},
		{"off", BinFlagOff},
		{"on", BinFlagOn},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := BinFlagFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestBinFlagFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := BinFlagFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "invalid BinFlag string"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestBinFlag_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected BinFlag
	}{
		{`<element onOff="0"></element>`, BinFlagZero},
		{`<element onOff="1"></element>`, BinFlagOne},
		{`<element onOff="false"></element>`, BinFlagFalse},
		{`<element onOff="true"></element>`, BinFlagTrue},
		{`<element onOff="off"></element>`, BinFlagOff},
		{`<element onOff="on"></element>`, BinFlagOn},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name `xml:"element"`
				BinFlag BinFlag  `xml:"onOff,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.BinFlag != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.BinFlag)
			}
		})
	}
}

func TestBinFlag_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element onOff="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name `xml:"element"`
		BinFlag BinFlag  `xml:"onOff,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "invalid BinFlag string"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
