package stypes

import (
	"encoding/xml"
	"testing"
)

func TestTextAlignFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected TextAlign
	}{
		{"top", TextAlignTop},
		{"center", TextAlignCenter},
		{"baseline", TextAlignBaseline},
		{"bottom", TextAlignBottom},
		{"auto", TextAlignAuto},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := TextAlignFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestTextAlignFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := TextAlignFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid Text Alignment"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestTextAlign_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected TextAlign
	}{
		{`<element align="top"></element>`, TextAlignTop},
		{`<element align="center"></element>`, TextAlignCenter},
		{`<element align="baseline"></element>`, TextAlignBaseline},
		{`<element align="bottom"></element>`, TextAlignBottom},
		{`<element align="auto"></element>`, TextAlignAuto},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name  `xml:"element"`
				Align   TextAlign `xml:"align,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.Align != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.Align)
			}
		})
	}
}

func TestTextAlign_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element align="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name  `xml:"element"`
		Align   TextAlign `xml:"align,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid Text Alignment"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
