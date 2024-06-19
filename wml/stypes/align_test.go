package stypes

import (
	"encoding/xml"
	"testing"
)

func TestAlignFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected Align
	}{
		{"left", AlignLeft},
		{"center", AlignCenter},
		{"right", AlignRight},
		{"inside", AlignInside},
		{"outside", AlignOutside},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := AlignFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestAlignFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := AlignFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid Align value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestAlign_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected Align
	}{
		{`<element val="left"></element>`, AlignLeft},
		{`<element val="center"></element>`, AlignCenter},
		{`<element val="right"></element>`, AlignRight},
		{`<element val="inside"></element>`, AlignInside},
		{`<element val="outside"></element>`, AlignOutside},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name `xml:"element"`
				Val     Align    `xml:"val,attr"`
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

func TestAlign_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element val="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name `xml:"element"`
		Val     Align    `xml:"val,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid Align value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
