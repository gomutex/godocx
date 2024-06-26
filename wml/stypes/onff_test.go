package stypes

import (
	"encoding/xml"
	"testing"
)

func TestOnOffFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected OnOff
	}{
		{"0", OnOffZero},
		{"1", OnOffOne},
		{"false", OnOffFalse},
		{"true", OnOffTrue},
		{"off", OnOffOff},
		{"on", OnOffOn},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := OnOffFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestOnOffFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := OnOffFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "invalid OnOff string"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestOnOff_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected OnOff
	}{
		{`<element onOff="0"></element>`, OnOffZero},
		{`<element onOff="1"></element>`, OnOffOne},
		{`<element onOff="false"></element>`, OnOffFalse},
		{`<element onOff="true"></element>`, OnOffTrue},
		{`<element onOff="off"></element>`, OnOffOff},
		{`<element onOff="on"></element>`, OnOffOn},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name `xml:"element"`
				OnOff   OnOff    `xml:"onOff,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.OnOff != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.OnOff)
			}
		})
	}
}

func TestOnOff_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element onOff="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name `xml:"element"`
		OnOff   OnOff    `xml:"onOff,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "invalid OnOff string"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
