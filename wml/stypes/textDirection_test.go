package stypes

import (
	"encoding/xml"
	"testing"
)

func TestTextDirectionFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected TextDirection
	}{
		{"lrTb", TextDirectionLrTb},
		{"tbRl", TextDirectionTbRl},
		{"btLr", TextDirectionBtLr},
		{"lrTbV", TextDirectionLrTbV},
		{"tbRlV", TextDirectionTbRlV},
		{"tbLrV", TextDirectionTbLrV},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := TextDirectionFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestTextDirectionFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := TextDirectionFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid Text Direction"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestTextDirection_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected TextDirection
	}{
		{`<element direction="lrTb"></element>`, TextDirectionLrTb},
		{`<element direction="tbRl"></element>`, TextDirectionTbRl},
		{`<element direction="btLr"></element>`, TextDirectionBtLr},
		{`<element direction="lrTbV"></element>`, TextDirectionLrTbV},
		{`<element direction="tbRlV"></element>`, TextDirectionTbRlV},
		{`<element direction="tbLrV"></element>`, TextDirectionTbLrV},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName   xml.Name      `xml:"element"`
				Direction TextDirection `xml:"direction,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.Direction != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.Direction)
			}
		})
	}
}

func TestTextDirection_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element direction="invalidValue"></element>`

	type Element struct {
		XMLName   xml.Name      `xml:"element"`
		Direction TextDirection `xml:"direction,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid Text Direction"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
