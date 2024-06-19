package stypes

import (
	"encoding/xml"
	"testing"
)

func TestAnchorFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected Anchor
	}{
		{"text", AnchorText},
		{"margin", AnchorMargin},
		{"page", AnchorPage},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := AnchorFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestAnchorFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := AnchorFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid Anchor value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestAnchor_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected Anchor
	}{
		{`<element val="text"></element>`, AnchorText},
		{`<element val="margin"></element>`, AnchorMargin},
		{`<element val="page"></element>`, AnchorPage},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name `xml:"element"`
				Val     Anchor   `xml:"val,attr"`
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

func TestAnchor_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element val="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name `xml:"element"`
		Val     Anchor   `xml:"val,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid Anchor value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
